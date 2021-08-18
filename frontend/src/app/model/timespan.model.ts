export class TimeSpan {
  start: Date;
  stop: Date;

  public constructor(start: Date, stop: Date) {
    this.start = start;
    this.stop = stop;
  }

  public verify(): boolean {
      return this.start.getFullYear() > 2020 && this.stop.getFullYear() > 2020;
  }

  public static empty(): TimeSpan {
    return new TimeSpan(new Date(1900, 1, 1, 0, 0, 0, 0), new Date(1900, 1, 1, 0, 0, 0, 0));
  }

  public static now(): TimeSpan {
    let startDate: Date = new Date();
    let stopDate: Date = new Date(startDate);
    stopDate.setHours(stopDate.getHours() + 1);
    return new TimeSpan(startDate, stopDate);
  }
}
