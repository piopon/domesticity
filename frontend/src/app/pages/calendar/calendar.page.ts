import { Component, OnInit, ViewChild } from "@angular/core";
import { ModalController } from "@ionic/angular";
import { CalendarComponent } from "ionic2-calendar";
import { CalendarMode, IEvent } from "ionic2-calendar/calendar";
import { DayEventsPage } from "src/app/dialogs/day-events/day-events.page";
import { TextEventsService } from "src/app/services/text-events.service";

@Component({
  selector: "app-calendar",
  templateUrl: "./calendar.page.html",
  styleUrls: ["./calendar.page.scss"],
})
export class CalendarPage implements OnInit {
  @ViewChild(CalendarComponent) myCalendar: CalendarComponent;

  private calendarData: any;
  private eventSource: IEvent[] = [];
  private calendarModes: CalendarMode[] = ["month", "week", "day"];

  constructor(public modalController: ModalController, private textEventsService: TextEventsService) {}

  ngOnInit() {
    this.calendarData = {
      titleMonth: "" as string,
      currentDate: new Date() as Date,
      mode: this.calendarModes[0] as CalendarMode,
    };
    let currentMonth = this.calendarData.currentDate.getMonth();
    this.updateEventSource(currentMonth);
  }

  protected nextMonth() {
    this.myCalendar.slideNext();
  }

  protected previousMonth() {
    this.myCalendar.slidePrev();
  }

  protected onMonthChanged(newTitle: string) {
    this.calendarData.titleMonth = newTitle;
  }

  protected async onTimeSelected(event: { selectedTime: Date; events: any[] }) {
    if ("month" === this.calendarData.mode) {
      const modal = await this.modalController.create({
        component: DayEventsPage,
        componentProps: {
          dayTime: event.selectedTime,
          dayEvents: event.events,
        },
      });
      return await modal.present();
    }
  }

  protected changeView() {
    let modeIndex: number = this.calendarModes.indexOf(this.calendarData.mode);
    modeIndex = (modeIndex + 1) % this.calendarModes.length;
    this.calendarData.mode = this.calendarModes[modeIndex];
  }

  private updateEventSource(monthNo: number): void {
    var currentEvents: IEvent[] = [];
    this.textEventsService
      .getEventsInMonth(monthNo)
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
      .then((_) => (this.eventSource = currentEvents));
  }
}
