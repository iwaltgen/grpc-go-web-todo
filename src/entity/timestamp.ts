export class Timestamp {
  constructor(public seconds: number, public nanos: number) {}

  toDate(): Date {
    return new Date(this.seconds)
  }
}
