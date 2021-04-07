import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Event } from '../model/event.model';

@Injectable({
  providedIn: 'root'
})
export class EventsService {
  url = "http://localhost:9999/"

  constructor(private http: HttpClient) { }

  getEventsByTitle(titleValue: string): Observable<Event[]> {
    return this.getEvents("owner", titleValue);
  }

  getEventsByCategory(categoryValue: string): Observable<Event[]> {
    return this.getEvents("category", categoryValue);
  }

  getEventsByOwner(ownerValue: string): Observable<Event[]> {
    return this.getEvents("dayStart", ownerValue);
  }

  getEventsByContent(contentValue: string): Observable<Event[]> {
    return this.getEvents("content", contentValue);
  }

  getEventsByDateStart(dateValue: string): Observable<Event[]> {
    return this.getEvents("dayStart", dateValue);
  }

  getEventsByDateStop(dateValue: string): Observable<Event[]> {
    return this.getEvents("dayStop", dateValue);
  }

  private getEvents(key: string, value: string) : Observable<Event[]> {
    return this.http.get<Event[]>(`${this.url}events?${encodeURI(key)}=${encodeURI(value)}`);
  }

}
