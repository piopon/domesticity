import { TimeSpan } from "./timespan.model";

export class Event {
  id: string;
  title: string;
  icon: string;
  owner: string;
  date: TimeSpan;
  category: string;
  content: string;

  private static availableIcons: string[] = [
    "airplane", "basketball", "beer", "bicycle", "book", "bonfire", "bus", "cafe",
    "car", "cart", "construct", "dice", "fast-food", "football", "game-controller", "school",
  ];

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

  public verify(): string[] {
    let result: string[] = [];
    if (this.title === "") {
      result.push("title cannot be empty");
    }
    if (this.owner === "") {
      result.push("select owner");
    }
    if (this.category === "") {
      result.push("select category");
    }
    if (this.content === "") {
      result.push("content cannot be empty");
    }
    this.date.verify().forEach(error => result.push(error));
    return result;
  }

  public static getAvailableIcons(): string[] {
    return Event.availableIcons;
  }

  public static empty(): Event {
    return new Event("", "", "", "", TimeSpan.now(), "", "");
  }
}
