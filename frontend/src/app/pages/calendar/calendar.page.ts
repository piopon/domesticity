import { Component, OnInit, ViewChild } from "@angular/core";
import { ModalController } from "@ionic/angular";
import { CalendarComponent } from "ionic2-calendar";
import { IEvent } from "ionic2-calendar/calendar";
import { DayEventsPage } from "src/app/dialogs/day-events/day-events.page";
import { TimeSpan } from "src/app/model/timespan.model";

@Component({
  selector: "app-calendar",
  templateUrl: "./calendar.page.html",
  styleUrls: ["./calendar.page.scss"],
})
export class CalendarPage implements OnInit {
  private calendarData: any;
  private eventSource: IEvent[] = [];
  private availableModes: string[] = ["month", "week", "day"];

  @ViewChild(CalendarComponent) myCalendar: CalendarComponent;

  constructor(public modalController: ModalController) {}

  ngOnInit() {
    this.calendarData = {
      titleMonth: "",
      currentDate: new Date(),
      mode: this.availableModes[0],
    };
    let currentMonth = this.calendarData.currentDate.getMonth();
    this.eventSource = this.getEventSourceForMonth(currentMonth);
  }

  nextMonth() {
    this.myCalendar.slideNext();
  }

  previousMonth() {
    this.myCalendar.slidePrev();
  }

  onMonthChanged(newTitle: string) {
    this.calendarData.titleMonth = newTitle;
  }

  async onTimeSelected(event: { selectedTime: Date; events: any[] }) {
    const modal = await this.modalController.create({
      component: DayEventsPage,
      componentProps: {
        dayTime: event.selectedTime,
        dayEvents: event.events,
      },
    });
    return await modal.present();
  }

  changeView() {
    let modeIndex: number = this.availableModes.indexOf(this.calendarData.mode);
    modeIndex = (modeIndex + 1) % this.availableModes.length;
    this.calendarData.mode = this.availableModes[modeIndex];
  }

  private getEventSourceForMonth(monthNo: number): IEvent[] {
    this.eventSource.push({
      allDay: false,
      endTime: TimeSpan.now().stop,
      startTime: TimeSpan.now().start,
      title: "test",
    });
    return [];
  }
}
