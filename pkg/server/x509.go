package server

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// reference: https://golang.org/src/crypto/tls/generate_cert.go
var (
	// DevHosts addr list
	DevHosts = []string{"localhost", "127.0.0.1", "::1"}
)

const (
	rsaBits            = 2048
	certificateExpires = time.Hour * 24 * 256
	mkcertRootName     = "rootCA.pem"
	mkcertRootKeyName  = "rootCA-key.pem"
)

var (
	errNotFoundMkcertCAROOT = errors.New("not found mkcert CAROOT")
)

type x509Generator struct{}

func (x x509Generator) x509SerialNumber() *big.Int {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, serialNumberLimit)
	return serialNumber
}

func (x x509Generator) x509PemBlockForPrivateKey(privateKey interface{}) *pem.Block {
	bytes, _ := x509.MarshalPKCS8PrivateKey(privateKey)
	return &pem.Block{Type: "PRIVATE KEY", Bytes: bytes}
}

func (x x509Generator) x509PemBlockForCertificate(derBytes []byte) *pem.Block {
	return &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}
}

// create x509 key pair bytes
func (x x509Generator) newX509KeyPairBytes(hosts []string) (cert, key []byte, err error) {
	certPEM, keyPEM, err := x.newX509KeyPair(hosts)
	if err != nil {
		return nil, nil, err
	}

	certOut := &bytes.Buffer{}
	if err := pem.Encode(certOut, certPEM); err != nil {
		return nil, nil, fmt.Errorf("encode PEM certificate error: %w", err)
	}
	cert = certOut.Bytes()

	keyOut := &bytes.Buffer{}
	if err := pem.Encode(keyOut, keyPEM); err != nil {
		return nil, nil, fmt.Errorf("encode PEM private key error: %w", err)
	}
	key = keyOut.Bytes()

	return cert, key, nil
}

// create x509 key pair
func (x x509Generator) newX509KeyPair(hosts []string) (cert, key *pem.Block, err error) {
	caRootCert, caRootKey, err := x.loadMkcertCARoot()
	if err != nil && err != errNotFoundMkcertCAROOT {
		return nil, nil, err
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, rsaBits)
	if err != nil {
		return nil, nil, fmt.Errorf("create x509 private key error: %w", err)
	}

	template := x509.Certificate{
		SerialNumber: x.x509SerialNumber(),
		Subject: pkix.Name{
			Country:      []string{"KR"},
			Organization: []string{"KPS DTTTD"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(certificateExpires),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	parent := &template
	if caRootCert != nil {
		parent = caRootCert
	}

	pub := &privateKey.PublicKey
	priv := interface{}(privateKey)
	if caRootKey != nil {
		priv = caRootKey
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, parent, pub, priv)
	if err != nil {
		return nil, nil, fmt.Errorf("create x509 certificate error: %w", err)
	}

	return x.x509PemBlockForCertificate(derBytes), x.x509PemBlockForPrivateKey(privateKey), nil
}

// load mkcert(https://github.com/FiloSottile/mkcert#mkcert) CARoot
func (x x509Generator) loadMkcertCARoot() (*x509.Certificate, interface{}, error) {
	caRoot := x.pathMkcertCARoot()
	if _, err := os.Stat(filepath.Join(caRoot, mkcertRootName)); err != nil {
		return nil, nil, errNotFoundMkcertCAROOT
	}
	if _, err := os.Stat(filepath.Join(caRoot, mkcertRootKeyName)); err != nil {
		return nil, nil, errNotFoundMkcertCAROOT
	}

	certPEMBlock, err := ioutil.ReadFile(filepath.Join(caRoot, mkcertRootName))
	if err != nil {
		return nil, nil, fmt.Errorf("read the CA certificate error: %w", err)
	}

	certDERBlock, _ := pem.Decode(certPEMBlock)
	if certDERBlock == nil || certDERBlock.Type != "CERTIFICATE" {
		return nil, nil, fmt.Errorf("read the CA certificate unexpected content error: %w", err)
	}

	cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("parse the CA certificate error: %w", err)
	}

	keyPEMBlock, err := ioutil.ReadFile(filepath.Join(caRoot, mkcertRootKeyName))
	if err != nil {
		return nil, nil, fmt.Errorf("read the CA key error: %w", err)
	}

	keyDERBlock, _ := pem.Decode(keyPEMBlock)
	if keyDERBlock == nil || keyDERBlock.Type != "PRIVATE KEY" {
		return nil, nil, fmt.Errorf("read the CA key unexpected content error: %w", err)
	}

	key, err := x509.ParsePKCS8PrivateKey(keyDERBlock.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("parse the CA key error: %w", err)
	}

	return cert, key, nil
}

func (x x509Generator) pathMkcertCARoot() string {
	if env := os.Getenv("CAROOT"); env != "" {
		return env
	}

	var dir string
	switch {
	case runtime.GOOS == "windows":
		dir = os.Getenv("LocalAppData")

	case os.Getenv("XDG_DATA_HOME") != "":
		dir = os.Getenv("XDG_DATA_HOME")

	case runtime.GOOS == "darwin":
		dir = os.Getenv("HOME")
		if dir == "" {
			return ""
		}
		dir = filepath.Join(dir, "Library", "Application Support")

	default: // Unix
		dir = os.Getenv("HOME")
		if dir == "" {
			return ""
		}
		dir = filepath.Join(dir, ".local", "share")

	}

	return filepath.Join(dir, "mkcert")
}
