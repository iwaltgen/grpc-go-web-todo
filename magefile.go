// +build mage

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"golang.org/x/oauth2"
)

const (
	packageName = "github.com/iwaltgen/grpc-go-web-todo"
	version     = "0.1.0"
	ldflags     = "-ldflags=-s -w" +
		" -X main.version=$VERSION" +
		" -X main.commitHash=$COMMIT_HASH" +
		" -X main.buildDate=$BUILD_DATE"
)

var (
	goexe = "go"
	git   = sh.RunCmd("git")

	workspace = packageName
	started   = time.Now().Unix()
)

func init() {
	goexe = mg.GoCmd()
	workspace, _ = os.Getwd()
}

func buildEnv() map[string]string {
	hash, _ := sh.Output("git", "rev-parse", "--verify", "HEAD")
	return map[string]string{
		"PACKAGE":     packageName,
		"WORKSPACE":   workspace,
		"VERSION":     version,
		"COMMIT_HASH": hash,
		"BUILD_DATE":  fmt.Sprintf("%d", started),
	}
}

// Build build frontend & backend app
func Build() error {
	mg.Deps(Lint)

	if err := sh.RunV("npm", "run", "build"); err != nil {
		return err
	}

	mg.Deps(GenStatic)

	args := []string{"build", "-trimpath", ldflags, "-o", "./build/server", "./cmd/server"}
	return sh.RunWith(buildEnv(), goexe, args...)
}

// Dev serve frontend & backend development
func Dev() error {
	mg.Deps(Lint)

	go func() {
		_ = sh.RunV("npm", "run", "dev")
	}()
	return sh.RunV("air")
}

// BuildDev build backend for development
func BuildDev() error {
	args := []string{"build", "-trimpath", ldflags, "-o", "./tmp/server", "./cmd/server"}
	return sh.RunWith(buildEnv(), goexe, args...)
}

// Clean clean build artifacts
func Clean() {
	_ = sh.Rm("public/build")
	_ = sh.Rm("build")
}

// Test test frontend & backend app
func Test() error {
	return sh.RunV("sh", "-c", "go test ./pkg/... -cover -json | tparse -all")
}

// Lint lint frontend & backend app
func Lint() error {
	// TODO(iwaltgen): svelte typescript support not yet
	if err := sh.RunV("npm", "run", "lint"); err != nil {
		return err
	}

	return sh.RunV("golangci-lint", "run", "--timeout", "3m", "-E", "misspell")
}

// GenAPI generate API
func GenAPI() error {
	gopath, err := sh.Output(goexe, "env", "GOPATH")
	if err != nil {
		return err
	}

	env := map[string]string{
		"PROTOTOOL_PROTOC_BIN_PATH": "bin/protoc",
		"PROTOTOOL_PROTOC_WKT_PATH": gopath + "/src/github.com/gogo/protobuf/protobuf",
	}
	if err := sh.RunWith(env, "prototool", "lint"); err != nil {
		return err
	}
	return sh.RunWith(env, "prototool", "generate")
}

// GenWire generate wire code
func GenWire() error {
	return sh.RunV("wire", "./pkg/...")
}

// GenStatic generate frontend static resource for backend embed
func GenStatic() error {
	return sh.RunV(goexe, "generate", "./pkg/server/...")
}

// Gen generate API & embed resource
func Gen() error {
	mg.Deps(GenAPI, GenWire)
	return sh.RunV(goexe, "generate", "./pkg/...")
}

// All generate, build app
func All() {
	mg.Deps(Lint, Test, Build)
}

// Install install package & tool
func Install() error {
	color.Green("install tools...")
	gg := gg{}
	modules := []string{
		"github.com/golang/protobuf/protoc-gen-go",
		"github.com/gogo/protobuf/protoc-gen-gogo",
		"github.com/gogo/protobuf/protoc-gen-gofast",
		"github.com/gogo/protobuf/protoc-gen-gogofast",
		"github.com/gogo/protobuf/protoc-gen-gogofaster",
		"github.com/gogo/protobuf/protoc-gen-gogoslick",
		// TODO(iwaltgen): protoc-gen-validate
		"golang.org/x/tools/cmd/stringer",
		"github.com/google/wire/cmd/wire",
		"github.com/rakyll/statik",
		"github.com/mfridman/tparse",
		"github.com/cosmtrek/air",
	}
	for _, v := range modules {
		if err := gg.installModule(v); err != nil {
			return err
		}
	}

	gh := newGithub()
	repos := []ghRepo{
		{owner: "grpc", name: "grpc-web", target: "protoc-gen-grpc-web"},
		{owner: "uber", name: "prototool", target: "prototool"},
		// TODO(iwaltgen): protoc
		// TODO(iwaltgen): golangci-lint
		// TODO(iwaltgen): grpcurl
	}
	for _, v := range repos {
		if err := gh.downloadLatestReleaseFile(v); err != nil {
			return err
		}
	}

	if err := installLint(); err != nil {
		return err
	}

	color.Green("install packages...")
	return sh.RunV("npm", "install")
}

// get get
type gg struct{}

func (g gg) installModule(uri string) error {
	env := map[string]string{"GO111MODULE": "off"}
	return sh.RunWith(env, goexe, "get", uri)
}

// github
type gh struct {
	*github.Client
}

func newGithub() gh {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	return gh{github.NewClient(tc)}
}

type ghRepo struct {
	owner  string
	name   string
	target string
}

func (g gh) downloadLatestReleaseFile(repo ghRepo) error {
	target := "bin/" + repo.target
	if existsFile(target) {
		return nil
	}

	ctx := context.Background()
	release, _, err := g.Repositories.GetLatestRelease(ctx, repo.owner, repo.name)
	if err != nil {
		return err
	}

	opt := &github.ListOptions{
		Page:    1,
		PerPage: 64,
	}
	assets, _, err := g.Repositories.ListReleaseAssets(ctx, repo.owner, repo.name, release.GetID(), opt)
	if err != nil {
		return err
	}

	os, _ := g.machineInfo()
	arch := "x86_64"
	for _, asset := range assets {
		name := strings.ToLower(asset.GetName())
		if strings.Contains(name, repo.target) && strings.Contains(name, os) {
			archIndex := strings.LastIndex(name, arch)
			if archIndex != -1 && (len(name)-archIndex == len(arch)) {
				if err := g.downloadFile(asset.GetBrowserDownloadURL(), target); err != nil {
					return err
				}
				if err := sh.RunV("chmod", "+x", target); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (g gh) downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func installLint() error {
	target := "bin/golangci-lint"
	if existsFile(target) {
		return nil
	}

	curl := exec.Command("curl", "-sSfL", "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh")
	install := exec.Command("sh", "-s", "latest")
	install.Stdin, _ = curl.StdoutPipe()
	install.Stdout = os.Stdout

	_ = install.Start()
	_ = curl.Run()
	return install.Wait()
}

func (g gh) machineInfo() (string, string) {
	os, _ := sh.Output(goexe, "env", "GOOS")
	arch, _ := sh.Output(goexe, "env", "GOARCH")
	return os, arch
}

func existsFile(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}
