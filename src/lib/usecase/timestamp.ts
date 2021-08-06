import type { Timestamp } from '../entity';

export function now(): Timestamp {
  const date = new Date();
  return {
    seconds: Math.floor(date.getTime() / 1_000),
    nanos: date.getMilliseconds() * 1_000_000,
  };
}
