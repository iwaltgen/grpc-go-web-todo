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
	version     = "0.0.1"
	ldflags     = "-ldflags=-s -w" +
		" -X $PACKAGE/cmd/server/main.version=$VERSION" +
		" -X $PACKAGE/cmd/server/main.commitHash=$COMMIT_HASH" +
		" -X $PACKAGE/cmd/server/main.buildDate=$BUILD_DATE"
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

func buildFlags() map[string]string {
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
	mg.Deps(GenAPI, Lint)

	if err := sh.RunV("npm", "run", "build"); err != nil {
		return err
	}

	args := []string{"build", "-trimpath", ldflags, "-o", "./build/server", "./cmd/server"}
	return sh.RunWith(buildFlags(), goexe, args...)
}

// Clean clean build artifacts
func Clean() {
	_ = sh.Rm("public/build")
	_ = sh.Rm("build")
}

// Test test frontend & backend app
func Test() error {
	test := exec.Command(goexe, "test", "./pkg/...", "-cover", "-json")
	parse := exec.Command("tparse")
	parse.Stdin, _ = test.StdoutPipe()
	parse.Stdout = os.Stdout

	_ = parse.Start()
	_ = test.Run()
	return parse.Wait()
}

// Lint lint frontend & backend app
func Lint() error {
	if err := sh.RunV("npm", "run", "validate"); err != nil {
		return err
	}

	return sh.RunV("bin/golangci-lint", "run", "--timeout", "3m", "-E", "misspell")
}

// GenAPI generate API
func GenAPI() error {
	if err := sh.RunV("bin/prototool", "lint"); err != nil {
		return err
	}
	return sh.RunV("bin/prototool", "generate")
}

// All generate, build app
func All() {
	mg.Deps(GenAPI, Lint, Test, Build)
}

// Install install package & tool
func Install() error {
	color.Green("install tools...")
	gg := gg{}
	modules := []string{
		"golang.org/x/lint/golint",
		"github.com/golang/protobuf/protoc-gen-go",
		"github.com/gogo/protobuf/protoc-gen-gogo",
		"github.com/gogo/protobuf/protoc-gen-gofast",
		"github.com/gogo/protobuf/protoc-gen-gogofast",
		"github.com/gogo/protobuf/protoc-gen-gogofaster",
		"github.com/gogo/protobuf/protoc-gen-gogoslick",
		"github.com/rakyll/statik",
		"github.com/mfridman/tparse",
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
