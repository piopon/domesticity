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

  getEventsByTitle(titleValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.getEvents("owner", titleValue, limit, offset);
  }

  getEventsByCategory(categoryValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.getEvents("category", categoryValue, limit, offset);
  }

  getEventsByOwner(ownerValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.getEvents("dayStart", ownerValue, limit, offset);
  }

  getEventsByContent(contentValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.getEvents("content", contentValue, limit, offset);
  }

  getEventsByDateStart(dateValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.getEvents("dayStart", dateValue, limit, offset);
  }

  getEventsByDateStop(dateValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.getEvents("dayStop", dateValue, limit, offset);
  }

  private getEvents(key: string, value: string, limit: number, offset: number) : Observable<Event[]> {
    let modifiers : string = "";
    if (limit > 0) {
      modifiers += `limit=${encodeURI(limit.toString())}&`;
    }
    if (offset > 0) {
      modifiers += `offset=${encodeURI(offset.toString())}&`;
    }
    return this.http.get<Event[]>(`${this.url}events?${modifiers}${encodeURI(key)}=${encodeURI(value)}`);
  }

}
