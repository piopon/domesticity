import { Component, OnInit, ViewChild } from "@angular/core";
import { ModalController } from "@ionic/angular";
import { CalendarComponent } from "ionic2-calendar";
import { IEvent } from "ionic2-calendar/calendar";
import { DayEventsPage } from "src/app/dialogs/day-events/day-events.page";
import { TextEventsService } from "src/app/services/text-events.service";

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

  constructor(public modalController: ModalController, private textEventsService: TextEventsService) {}

  ngOnInit() {
    this.calendarData = {
      titleMonth: "",
      currentDate: new Date(),
      mode: this.availableModes[0],
    };
    let currentMonth = this.calendarData.currentDate.getMonth();
    this.eventSource = this.getEventSourceForMonth(currentMonth);

    let start = new Date();
    start.setDate(11);
    let stop = new Date();
    stop.setDate(11);
    this.eventSource.push({
      allDay: false,
      endTime: stop,
      startTime: start,
      title: "test",
    });
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
    let currEventSource: IEvent[] = [];
    this.textEventsService.getEventsInMonth(monthNo).subscribe((events) => {
      events?.forEach((event) => {
        currEventSource.push({
          allDay: false,
          endTime: new Date(event.date.stop),
          startTime: new Date(event.date.start),
          title: event.title,
        });
      });
    });
    return currEventSource;
  }
}
