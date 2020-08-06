package message

import (
	"time"

	"github.com/gogo/protobuf/types"
)

// TimestampFromProto convert *types.Timestamp to time.Time
func TimestampFromProto(value *types.Timestamp) (ret time.Time) {
	ret, _ = types.TimestampFromProto(value)
	return ret
}

// TimestampProto convert time.Time to *types.Timestamp
func TimestampProto(value time.Time) (ret *types.Timestamp) {
	ret, _ = types.TimestampProto(value)
	return ret
}
