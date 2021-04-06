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

  getEventsByOwner(owner: string): Observable<Event[]> {
    return this.http.get<Event[]>(`${this.url}events?owner=${encodeURI(owner)}`);
  }

  getEventsByDateStart(dateStart: string): Observable<Event[]> {
    return this.http.get<Event[]>(`${this.url}events?dayStart=${encodeURI(dateStart)}`);
  }

}
