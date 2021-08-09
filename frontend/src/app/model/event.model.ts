import { TimeSpan } from "./timespan.model";

export class Event {
  id: string;
  title: string;
  icon: string;
  owner: string;
  date: TimeSpan;
  category: string;
  content: string;

  public constructor(id: string, title: string, icon: string,
                     owner: string, date: TimeSpan, category: string, content: string) {
    this.id = id;
    this.title = title;
    this.icon = icon;
    this.owner = owner;
    this.date = date;
    this.category = category;
    this.content = content;
  }

  public static empty(): Event {
    let emptyDate: TimeSpan = {
      start: new Date(1900, 1, 1, 1, 1, 1, 1),
      stop: new Date(1900, 1, 1, 1, 1, 1, 1),
    }
    return new Event('', '', '', '', emptyDate, '', '');
  }
}
