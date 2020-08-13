// +build mage

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
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

	color.Green("go get global...")
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

	color.Green("install github release assets...")
	gh := gh{}
	repos := []ghdl{
		{repo: "grpc/grpc-web", asset: "protoc-gen-grpc-web", target: "protoc-gen-grpc-web"},
		{repo: "uber/prototool", asset: "prototool", target: "prototool"},
		{repo: "protocolbuffers/protobuf", asset: "protoc", pick: "protoc"},
		{repo: "golangci/golangci-lint", asset: "golangci-lint", pick: "golangci-lint"},
		{repo: "fullstorydev/grpcurl", asset: "grpcurl", pick: "grpcurl"},
	}
	for _, v := range repos {
		if err := gh.downloadReleaseAsset(v); err != nil {
			return err
		}
	}

	color.Green("install npm packages...")
	return sh.RunV("npm", "install")
}

// get get
type gg struct{}

func (g gg) installModule(uri string) error {
	env := map[string]string{"GO111MODULE": "off"}
	return sh.RunWith(env, goexe, "get", uri)
}

// github
type gh struct{}

type ghdl struct {
	repo   string
	asset  string
	target string
	pick   string
}

func (g gh) downloadReleaseAsset(target ghdl) error {
	destination := "bin/" + target.target
	if target.target == "" {
		destination += target.pick
	}
	if existsFile(destination) {
		return nil
	}

	err := sh.RunV("github-dl",
		"--repo", target.repo,
		"--asset", target.asset,
		"--dest", "bin",
		"--target", target.target,
		"--pick", target.pick,
	)
	if err != nil {
		return err
	}

	if target.target != "" {
		if err := sh.RunV("chmod", "+x", destination); err != nil {
			return err
		}
	}
	return nil
}

func existsFile(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}
