package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/iwaltgen/grpc-go-web-todo/pkg/log"
	"github.com/iwaltgen/grpc-go-web-todo/pkg/server"
)

var (
	version    = "dev"
	commitHash = "dev"
	buildDate  = ""
	startTime  time.Time
)

func init() {
	startTime = time.Now()
}

// BuildDate build date
func BuildDate() time.Time {
	buildTime, err := unixStringToTime(buildDate)
	if err != nil {
		return startTime
	}
	return buildTime
}

func unixStringToTime(unixStr string) (t time.Time, err error) {
	i, err := strconv.ParseInt(unixStr, 10, 64)
	if err != nil {
		return t, fmt.Errorf("parse unix timestamp string: %w", err)
	}
	return time.Unix(i, 0).UTC(), nil
}

func main() {
	log.L("cmd.server").Info("metadata",
		log.String("version", version),
		log.String("commit_hash", commitHash),
		log.String("build_date", BuildDate().UTC().Format(time.RFC3339)),
	)
	daemon := server.NewComposer()
	daemon.Serve(context.Background())
}
