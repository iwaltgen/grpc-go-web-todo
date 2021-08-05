package version

import (
	"strconv"
	"time"
)

var (
	version    = "dev"
	commitHash = "dev"
	buildDate  = "1609416000" // 2020-12-31 12:00:00 PM
)

// Version represents version
func Version() string {
	return version
}

// CommitHash represents latest git hash
func CommitHash() string {
	return commitHash
}

// BuildDate represents build time
func BuildDate() time.Time {
	ts, _ := strconv.ParseInt(buildDate, 10, 64)
	return time.Unix(ts, 0).UTC()
}
