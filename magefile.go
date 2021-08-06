// +build mage

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/bitfield/script"
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"
	"github.com/mattn/go-zglob"
)

const (
	packageName = "github.com/iwaltgen/grpc-go-web-todo"
	version     = "0.1.1"
	targetDir   = "build"
	binaryDir   = "bin"
)

type (
	API   mg.Namespace
	GEN   mg.Namespace
	BUILD mg.Namespace
	DEV   mg.Namespace
)

var (
	started       int64
	gocmd, npmcmd func(args ...string) error
	workspace     string
)

func init() {
	workspace, _ = os.Getwd()
	started = time.Now().Unix()
	gocmd = sh.RunCmd(mg.GoCmd())
	npmcmd = func(args ...string) error {
		return sh.RunV("npm", args...)
	}
}

// Run lint frontend & backend app
func Lint() error {
	if err := npmcmd("run", "lint"); err != nil {
		return err
	}

	return sh.RunV("golangci-lint", "run")
}

// Run test frontend & backend app
func Test() error {
	mg.Deps(Lint)
	mg.Deps(GEN.API)

	output, err := sh.Output(mg.GoCmd(), "list", "./...")
	if err != nil {
		return fmt.Errorf("failed to go list: %w", err)
	}

	var pkgs []string
	for _, line := range strings.Split(output, "\n") {
		if !strings.Contains(line, "infra") {
			pkgs = append(pkgs, line)
		}
	}

	return execPipeCmdsStdout(
		fmt.Sprintf("go test %s -timeout 10s -race -cover -json", strings.Join(pkgs, "\n")),
		"tparse -all",
	)
}

// Run dev serve frontend & backend development
func Dev() error {
	mg.Deps(Build)

	go func() {
		_ = npmcmd("run", "dev", "--", "--open")
	}()
	return sh.RunV("server")
}

// Build frontend & backend app
func Build() error {
	mg.Deps(Lint)
	mg.Deps(GEN.API)

	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		return err
	}

	b := BUILD{}
	if err := b.Front(); err != nil {
		return err
	}

	return sh.RunWith(b.buildEnv(), mg.GoCmd(), b.buildParameters()...)
}

func (BUILD) Front() error {
	changes, err := changesTarget("public/build/bundle.js", "src/**/*")
	if err != nil {
		return err
	}
	if !changes {
		return nil
	}

	return npmcmd("run", "build")
}

func (BUILD) buildParameters() []string {
	ldflags := "-ldflags=" +
		"-X $PACKAGE/pkg/version.version=$VERSION " +
		"-X $PACKAGE/pkg/version.commitHash=$COMMIT_HASH " +
		"-X $PACKAGE/pkg/version.buildDate=$BUILD_DATE"
	return []string{"build", "-trimpath", ldflags, "-o", targetDir, "./cmd/..."}
}

func (BUILD) buildEnv() map[string]string {
	var hash string
	if gitsha := os.Getenv("GITHUB_SHA"); gitsha != "" {
		hash = gitsha
	} else {
		hash, _ = sh.Output("git", "rev-parse", "--verify", "HEAD")
	}

	return map[string]string{
		"CGO_ENABLED": "0",
		"PACKAGE":     packageName,
		"WORKSPACE":   workspace,
		"VERSION":     version,
		"COMMIT_HASH": hash,
		"BUILD_DATE":  fmt.Sprintf("%d", started),
	}
}

// Clean build artifacts
func Clean() error {
	_ = sh.Rm("build")
	return npmcmd("cache", "clean", "--force")
}

// Generate API
func (GEN) API() {
	mg.Deps(API.Gen)
}

// Generate wire code
func (GEN) Wire() error {
	return sh.RunV("wire", "./pkg/...")
}

// Generate helper code
func (GEN) Code() error {
	return sh.RunV("go", "generate", "./pkg/...")
}

// Generate API & code & resource
func Gen() {
	mg.Deps(GEN.Code)
	mg.Deps(GEN.Wire)
	mg.Deps(GEN.API)
}

// Check lint API
func (API) Lint() error {
	return sh.RunV("buf", "lint")
}

// Check breaking API
func (API) Breaking() error {
	tag := "v" + version
	return sh.RunV("buf", "breaking", "--against", ".git#tag="+tag)
}

// Generate API code
func (API) Gen() error {
	mg.Deps(API.Lint)
	changes, err := changesTarget("api/todo/v1/event.pb.go", "api/todo/**/*.proto")
	if err != nil {
		return err
	}
	if !changes {
		return nil
	}

	files, err := zglob.Glob("api/todo/**/*.proto")
	if err != nil {
		return err
	}

	return sh.RunV("buf", "generate", "--path", strings.Join(files, ","))
}

// Install install package & tool
func Install() error {
	color.Green("install go tools...")
	err := gocmd("install",
		"github.com/bufbuild/buf/cmd/buf",
		"github.com/bufbuild/buf/cmd/protoc-gen-buf-check-breaking",
		"github.com/bufbuild/buf/cmd/protoc-gen-buf-check-lint",
		"github.com/envoyproxy/protoc-gen-validate",
		"github.com/fullstorydev/grpcurl/cmd/grpcurl",
		"github.com/golangci/golangci-lint/cmd/golangci-lint",
		"github.com/google/wire/cmd/wire",
		"github.com/magefile/mage",
		"github.com/mfridman/tparse",
		"golang.org/x/tools/cmd/stringer",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc",
		"google.golang.org/protobuf/cmd/protoc-gen-go",
	)
	if err != nil {
		return err
	}

	color.Green("install github release binary tools...")
	if err := os.MkdirAll(binaryDir, os.ModePerm); err != nil {
		return err
	}

	ghfiles := []struct {
		api, pattern, rename string
	}{
		{
			api:     "https://api.github.com/repos/grpc/grpc-web/releases/latest",
			pattern: ".+{{.os}}.+x86_64$",
			rename:  "protoc-gen-grpc-web",
		},
	}

	for _, v := range ghfiles {
		pattern := strings.ReplaceAll(v.pattern, "{{.os}}", runtime.GOOS)
		res := execPipeCmds(
			fmt.Sprintf("curl -fsSL %s", v.api),
			fmt.Sprintf(`jq -r '.assets[] | select(.name | test("%s")) | .browser_download_url'`, pattern),
		)

		output, err := res.String()
		if err != nil {
			return err
		}

		source := strings.Trim(output, "\n")
		filename := path.Base(source)
		target := filepath.Join(binaryDir, filename)
		if err := sh.RunV("curl", "-fsSL", "-o", target, source); err != nil {
			return err
		}

		dest := filepath.Join(binaryDir, v.rename)
		if err := os.Rename(target, dest); err != nil {
			return err
		}

		if err := sh.RunV("chmod", "+x", dest); err != nil {
			return err
		}
	}

	color.Green("install npm packages...")
	return npmcmd("install")
}

// Update 3rd party proto files
func Update3rdPartyProto() error {
	protos := []struct {
		remote, local string
	}{
		{
			remote: "https://raw.githubusercontent.com/envoyproxy/protoc-gen-validate/master/validate/validate.proto",
			local:  "api/envoyproxy/pgv/validate.proto",
		},
		{
			remote: "https://raw.githubusercontent.com/googleapis/googleapis/master/google/rpc/code.proto",
			local:  "api/google/rpc/code.proto",
		},
		{
			remote: "https://raw.githubusercontent.com/googleapis/googleapis/master/google/rpc/error_details.proto",
			local:  "api/google/rpc/error_details.proto",
		},
		{
			remote: "https://raw.githubusercontent.com/googleapis/googleapis/master/google/rpc/status.proto",
			local:  "api/google/rpc/status.proto",
		},
	}

	for _, v := range protos {
		if err := os.MkdirAll(filepath.Dir(v.local), os.ModePerm); err != nil {
			return err
		}

		if err := sh.RunV("curl", "-fsSL", "-o", v.local, v.remote); err != nil {
			return err
		}
	}
	return nil
}

func execPipeCmdsStdout(cmds ...string) error {
	pipe := execPipeCmds(cmds...)
	if pipe == nil {
		return nil
	}

	_, err := pipe.Stdout()
	return err
}

func execPipeCmds(cmds ...string) *script.Pipe {
	if len(cmds) < 1 {
		return nil
	}

	pipe := script.NewPipe()
	for _, cmd := range cmds {
		pipe = pipe.Exec(cmd)
		pipe.SetError(nil)
	}
	return pipe
}

func replaceTextInFile(path, old, new string) error {
	read, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	newContents := strings.Replace(string(read), old, new, -1)
	return os.WriteFile(path, []byte(newContents), 0)
}

func changesTarget(dst string, globs ...string) (bool, error) {
	for _, g := range globs {
		files, err := zglob.Glob(g)
		if err != nil {
			return false, err
		}
		if len(files) == 0 {
			return false, fmt.Errorf("failed to glob didn't match any files: %s" + g)
		}

		shouldDo, err := target.Path(dst, files...)
		if err != nil {
			return false, err
		}
		if shouldDo {
			return true, nil
		}
	}
	return false, nil
}
