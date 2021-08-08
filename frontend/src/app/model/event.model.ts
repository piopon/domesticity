import { TimeSpan } from "./timespan.model";

export interface Event {
  id: string;
  title: string;
  icon: string;
  owner: string;
  date: TimeSpan;
  category: string;
  content: string;
}

export function emptyEvent(): Event {
  return {
    id: "",
    title: "",
    icon: "",
    owner: "",
    date: {
      start: new Date(1900, 1, 1, 1, 1, 1, 1),
      stop: new Date(1900, 1, 1, 1, 1, 1, 1),
    },
    category: "",
    content: "",
  };
}
