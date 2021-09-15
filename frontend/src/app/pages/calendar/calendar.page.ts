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

  private calendarModes: CalendarMode[] = ["month", "week", "day"];
  private pageData = {
    title: "" as string,
    today: new Date() as Date,
    events: [] as IEvent[],
    viewMode: this.calendarModes[0] as CalendarMode,
  };

  constructor(public modalController: ModalController, private textEventsService: TextEventsService) {}

  ngOnInit() {
    this.updateEventSource(this.pageData.today.getMonth());
  }

  protected nextMonth() {
    this.myCalendar.slideNext();
  }

  protected previousMonth() {
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
      .then((_) => (this.pageData.events = currentEvents));
  }
}
