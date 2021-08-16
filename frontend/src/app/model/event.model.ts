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

  public validate(): boolean {
    return this.title !== "" && this.owner !== "" && this.category !== "" && this.content !== "";
  }

  public static empty(): Event {
    return new Event("", "", "", "", TimeSpan.now(), "", "");
  }
}
