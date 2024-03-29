import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { AlertController, ToastController } from "@ionic/angular";
import { Observable } from "rxjs";
import { first } from "rxjs/operators";
import { Event } from "../model/event.model";

@Injectable({
  providedIn: "root",
})
export class TextEventsService {
  private url: string = "http://localhost:9999/";
  private pingTimer: any;
  private pingInterval: number = 3_000;
  private online: boolean = true;
  private alertDialog: any;
  private alertVisible: boolean = false;

  constructor(
    private http: HttpClient,
    public alertController: AlertController,
    public toastController: ToastController
  ) {
    this.pingTimer = setInterval(() => this.pingService(), this.pingInterval);
  }

  public isOnline(): boolean {
    return this.online;
  }

  public addEvent(event: Event): Observable<any> {
    return this.http.post(`${this.url}events`, JSON.stringify(event), this.httpOptions());
  }

  public updateEvent(id: string, event: Event): Observable<any> {
    return this.http.put(`${this.url}events/${id}`, JSON.stringify(event), this.httpOptions());
  }

  public deleteEvent(id: string): Observable<any> {
    return this.http.delete(`${this.url}events/${id}`, this.httpOptions());
  }

  public getEvents(): Observable<Event[]> {
    return this.http.get<Event[]>(`${this.url}events`, this.httpOptions());
  }

  public getEventsByTitle(titleValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.filterEvents("owner", titleValue, limit, offset);
  }

  public getEventsByCategory(categoryValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.filterEvents("category", categoryValue, limit, offset);
  }

  public getEventsByOwner(ownerValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.filterEvents("dayStart", ownerValue, limit, offset);
  }

  public getEventsByContent(contentValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.filterEvents("content", contentValue, limit, offset);
  }

  public getEventsByDateStart(dateValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.filterEvents("dayStart", dateValue, limit, offset);
  }

  public getEventsByDateStop(dateValue: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.filterEvents("dayStop", dateValue, limit, offset);
  }

  public getEventsInMonth(monthDate: string, limit: number = 0, offset: number = 0): Observable<Event[]> {
    return this.filterEvents("inMonth", monthDate, limit, offset);
  }

  private filterEvents(key: string, value: string, limit: number, offset: number): Observable<Event[]> {
    let modifiers: string = "";
    if (limit > 0) {
      modifiers += `limit=${encodeURI(limit.toString())}&`;
    }
    if (offset > 0) {
      modifiers += `offset=${encodeURI(offset.toString())}&`;
    }
    return this.http.get<Event[]>(
      `${this.url}events?${modifiers}${encodeURI(key)}=${encodeURI(value)}`,
      this.httpOptions()
    );
  }

  private httpOptions(): Object {
    return {
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
    };
  }

  //********************************************************************************
  //*** ping and health check logic
  //********************************************************************************

  private async pingService(): Promise<void> {
    this.http
      .get(`${this.url}health`, { observe: "response", responseType: "text" })
      .pipe(first())
      .subscribe(
        async (response) => {
          this.online = 200 == response.status;
          if (this.online) {
            if (this.alertVisible) {
              this.alertDialog.dismiss();
              this.alertVisible = false;
            }
            if (this.pingInterval !== 3_000) {
              this.updatePingInterval(3_000);
              const toast = await this.toastController.create({
                message: "Reconnected to Text Event service.",
                duration: 2000,
              });
              toast.present();
            }
          }
        },
        async (_) => {
          this.online = false;
          if (this.alertVisible) {
            return;
          }
          this.alertDialog = await this.createAlertDialog();
          this.alertDialog.present();
          this.alertVisible = true;
          this.updatePingInterval(60_000);
        }
      );
  }

  private updatePingInterval(newInterval: number): void {
    clearInterval(this.pingTimer);
    this.pingInterval = newInterval;
    this.pingTimer = setInterval(() => this.pingService(), this.pingInterval);
  }

  private createAlertDialog(): Promise<HTMLIonAlertElement> {
    return this.alertController.create({
      header: "system error",
      subHeader: "service is offline",
      message: "text event service went offline",
      backdropDismiss: false,
      buttons: [
        {
          text: "dismiss",
          handler: () => {
            this.alertVisible = false;
            console.log("Dismiss pressed. Current ping interval: " + this.pingInterval);
          },
        },
        {
          text: "retry",
          handler: () => {
            this.alertVisible = false;
            this.pingService();
          },
        },
      ],
    });
  }
}
