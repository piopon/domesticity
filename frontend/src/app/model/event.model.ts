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