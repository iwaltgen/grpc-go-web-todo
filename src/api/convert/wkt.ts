import * as wkt_timestamp from 'google-protobuf/google/protobuf/timestamp_pb';
import type { Timestamp } from '../../entity'

export function timestampProto(v: Timestamp): wkt_timestamp.Timestamp {
  const ret = new wkt_timestamp.Timestamp();
  ret.setSeconds(v.seconds)
  ret.setNanos(v.nanos)
  return ret
}

export function timestampFromProto(v: wkt_timestamp.Timestamp): Timestamp {
  return v.toObject() as Timestamp
}
