import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AlertController } from '@ionic/angular';
import { Observable } from 'rxjs';
import { first } from "rxjs/operators";
import { Event } from '../model/event.model';

@Injectable({
  providedIn: 'root'
})
export class TextEventsService {
  private url:string = 'http://localhost:9999/';
  private pingTimer:any;
  private pingInterval:number = 3_000;
  private online:boolean = true;
  private alertVisible:boolean = false;

  constructor(private http: HttpClient, public alertController: AlertController) {
    this.pingTimer = setInterval(() => this.pingService(), this.pingInterval);
  }

  isOnline(): boolean {
    return this.online;
  }

  addEvent(event: Event): Observable<any> {
    return this.http.post(`${this.url}events`, JSON.stringify(event), this.httpOptions());
  }

  updateEvent(id: string, event: Event): Observable<any> {
    return this.http.put(`${this.url}events/${id}`, JSON.stringify(event), this.httpOptions());
  }

  deleteEvent(id: string): Observable<any> {
    return this.http.delete(`${this.url}events/${id}`, this.httpOptions());
  }

  getEvents(): Observable<Event[]> {
    return this.http.get<Event[]>(`${this.url}events`, this.httpOptions());
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

  private pingService():void {
    this.http.get(`${this.url}health`, {observe: 'response', responseType:'text'})
      .pipe(first())
      .subscribe(
        response => {
          this.online = (200 == response.status);
          if (this.online && this.pingInterval !== 3_000) {
            this.updatePingInterval(3_000);
          }
        },
        async _ => {
          this.online = false;
          if (this.alertVisible) {
            return;
          }
          const alert = await this.alertController.create({
            header: 'System error.',
            message: 'Offline service: Text Event',
            buttons: [
              {
                text: 'Dismiss',
                handler: () => {
                  this.alertVisible = false;
                  console.log('Dismiss pressed. Current ping interval: ' + this.pingInterval);
                }
              },
              {
                text: 'Retry',
                handler: () => {
                  this.alertVisible = false;
                  this.pingService();
                }
              }]
          });
          alert.present();
          this.alertVisible = true;
          this.updatePingInterval(60_000);
        }
      );
  }

  private updatePingInterval(newInterval: number):void {
    clearInterval(this.pingTimer);
    this.pingInterval = newInterval;
    this.pingTimer = setInterval(() => this.pingService(), this.pingInterval);
  }

  private filterEvents(key: string, value: string, limit: number, offset: number): Observable<Event[]> {
    let modifiers : string = "";
    if (limit > 0) {
      modifiers += `limit=${encodeURI(limit.toString())}&`;
    }
    if (offset > 0) {
      modifiers += `offset=${encodeURI(offset.toString())}&`;
    }
    return this.http.get<Event[]>(`${this.url}events?${modifiers}${encodeURI(key)}=${encodeURI(value)}`, this.httpOptions());
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
