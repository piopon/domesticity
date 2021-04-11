import { HttpClient, HttpHeaders, HttpParams, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Event } from '../model/event.model';

@Injectable({
  providedIn: 'root'
})
export class EventsService {
  url = "http://localhost:9999/"

  constructor(private http: HttpClient) { }

  addEvent(event: Event) : Observable<any> {
    return this.http.post(`${this.url}events`, JSON.stringify(event), this.httpOptions());
  }

  updateEvent(id: string, event: Event) : Observable<any> {
    return this.http.put(`${this.url}events/${id}`, JSON.stringify(event), this.httpOptions());
  }

  deleteEvent(id: string) : Observable<any> {
    return this.http.delete(`${this.url}events/${id}`);
  }

  getEvents() : Observable<Event[]> {
    return this.http.get<Event[]>(`${this.url}events`);
  }

  getEventsByTitle(titleValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.filterEvents("owner", titleValue, limit, offset);
  }

  getEventsByCategory(categoryValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.filterEvents("category", categoryValue, limit, offset);
  }

  getEventsByOwner(ownerValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.filterEvents("dayStart", ownerValue, limit, offset);
  }

  getEventsByContent(contentValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.filterEvents("content", contentValue, limit, offset);
  }

  getEventsByDateStart(dateValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.filterEvents("dayStart", dateValue, limit, offset);
  }

  getEventsByDateStop(dateValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.filterEvents("dayStop", dateValue, limit, offset);
  }

  private filterEvents(key: string, value: string, limit: number, offset: number) : Observable<Event[]> {
    let modifiers : string = "";
    if (limit > 0) {
      modifiers += `limit=${encodeURI(limit.toString())}&`;
    }
    if (offset > 0) {
      modifiers += `offset=${encodeURI(offset.toString())}&`;
    }
    return this.http.get<Event[]>(`${this.url}events?${modifiers}${encodeURI(key)}=${encodeURI(value)}`);
  }

  private httpOptions(): Object {
    return {
      headers: {
        'Accept' : 'application/json',
        'Content-Type': 'application/json'
      }
    };
  }

}
