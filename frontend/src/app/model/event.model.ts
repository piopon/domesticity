import { TimeSpan } from "./timespan.model";

export class Event {
  id: string;
  title: string;
  icon: string;
  owner: string;
  date: TimeSpan;
  category: string;
  content: string;

  public static empty(): Event {
    return {
      id: '',
      title: '',
      icon: '',
      owner: '',
      date: {
        start: new Date(1900, 1, 1, 1, 1, 1, 1),
        stop: new Date(1900, 1, 1, 1, 1, 1, 1),
      },
      category: '',
      content: '',
    };
  }
}
