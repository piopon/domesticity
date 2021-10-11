import { formatDate } from "@angular/common";
import { Component, OnInit, ViewChild } from "@angular/core";
import { ModalController } from "@ionic/angular";
import { CalendarComponent } from "ionic2-calendar";
import { CalendarMode, IEvent } from "ionic2-calendar/calendar";
import { Subscription } from "rxjs";
import { DayEventsPage } from "src/app/dialogs/day-events/day-events.page";
import { IpcType } from "src/app/model/ipc-message";
import { IpcMessagesService } from "src/app/services/ipc-messages.service";
import { TextEventsService } from "src/app/services/text-events.service";

@Component({
  selector: "app-calendar",
  templateUrl: "./calendar.page.html",
  styleUrls: ["./calendar.page.scss"],
})
export class CalendarPage implements OnInit {
  @ViewChild(CalendarComponent) myCalendar: CalendarComponent;

  private subscription: Subscription;
  private calendarModes: CalendarMode[] = ["month", "week", "day"];
  private pageData = {
    title: "" as string,
    today: new Date() as Date,
    events: [] as IEvent[],
    slider: { allowTouchMove: false, speed: 10 } as any,
    viewDate: new Date() as Date,
    viewMode: this.calendarModes[0] as CalendarMode,
  };

  constructor(
    public modalController: ModalController,
    private textEventsService: TextEventsService,
    private ipcMessagesService: IpcMessagesService
  ) {
    this.subscription = this.ipcMessagesService.watch().subscribe((ipcMessage) => {
      var ipcMessageDate = new Date(ipcMessage.message).toISOString().split('T')[0];

      var countEventsMatchingDate = this.pageData.events
        .map((event) => event.startTime.toISOString().split("T")[0])
        .filter((date) => date === ipcMessageDate).length;

      if (IpcType.AddEvent === ipcMessage.type && 0 === countEventsMatchingDate) {
        this.textEventsService
          .getEventsByDateStart(ipcMessageDate)
          .toPromise()
          .then((events) => {
            events?.forEach((event) => {
              this.pageData.events.push({
                title: event.title,
                startTime: new Date(event.date.start),
                endTime: new Date(event.date.stop),
                allDay: false,
              });
            });
          })
          .then((_) => this.myCalendar.loadEvents());
      } else if (IpcType.DeleteEvent === ipcMessage.type && 1 === countEventsMatchingDate) {
        var index = this.pageData.events.findIndex(
          (event) => event.startTime.toISOString().split("T")[0] === ipcMessageDate
        );
        this.pageData.events.splice(index, 1);
        this.myCalendar.loadEvents();
      }
    });
  }

  ngOnInit() {
    this.updateEventSource(this.getDateString(this.pageData.viewDate));
  }

  ngOnDestroy() {
    this.subscription.unsubscribe();
  }

  protected nextMonth() {
    this.pageData.viewDate.setMonth(this.pageData.viewDate.getMonth() + 1);
    this.updateEventSource(this.getDateString(this.pageData.viewDate));
    this.myCalendar.slideNext();
  }

  protected previousMonth() {
    this.pageData.viewDate.setMonth(this.pageData.viewDate.getMonth() - 1);
    this.updateEventSource(this.getDateString(this.pageData.viewDate));
    this.myCalendar.slidePrev();
  }

  protected onMonthChanged(newTitle: string) {
    this.pageData.title = newTitle;
  }

  protected async onTimeSelected(event: { selectedTime: Date; events: any[] }) {
    if ("month" === this.pageData.viewMode) {
      const modal = await this.modalController.create({
        component: DayEventsPage,
        componentProps: {
          dayTime: event.selectedTime,
        },
      });
      return await modal.present();
    }
  }

  protected changeView() {
    let modeIndex: number = this.calendarModes.indexOf(this.pageData.viewMode);
    modeIndex = (modeIndex + 1) % this.calendarModes.length;
    this.pageData.viewMode = this.calendarModes[modeIndex];
  }

  private updateEventSource(monthDate: string): void {
    var currentEvents: IEvent[] = [];
    this.textEventsService
      .getEventsInMonth(monthDate)
      .toPromise()
      .then((events) => {
        events?.forEach((event) => {
          currentEvents.push({
            title: event.title,
            startTime: new Date(event.date.start),
            endTime: new Date(event.date.stop),
            allDay: false,
          });
        });
      })
      .then((_) => this.pageData.events = currentEvents);
  }

  private getDateString(date: Date): string {
    return formatDate(date, "yyyy-MM-dd", "en");
  }
}
