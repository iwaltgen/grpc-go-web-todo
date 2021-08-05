package main

import (
	"context"
	"time"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/server"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/version"
)

func main() {
	logger := log.L("cmd.server")
	logger.Info("metadata", log.String("version", version.Version()))
	logger.Info("metadata", log.String("commit_hash", version.CommitHash()))
	logger.Info("metadata", log.String("build_date", version.BuildDate().Format(time.RFC3339)))

	daemon := server.NewComposer()
	daemon.Serve(context.Background())
}
