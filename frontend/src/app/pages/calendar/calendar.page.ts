import { Component, OnInit, ViewChild } from '@angular/core';
import { CalendarComponent } from 'ionic2-calendar';

@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.page.html',
  styleUrls: ['./calendar.page.scss'],
})
export class CalendarPage implements OnInit {
  availableModes:string[] = [ 'month', 'week', 'day' ];
  eventSource = [];
  titleMonth: string;

  calendar = {
    mode: this.availableModes[0],
    currentDate: new Date()
  };

  @ViewChild(CalendarComponent) myCalendar: CalendarComponent;

  constructor() { }

  ngOnInit() {
  }

  nextMonth() {
    this.myCalendar.slideNext();
  }

  previousMonth() {
    this.myCalendar.slidePrev();
  }

  onMonthChanged(newTitle:string) {
    this.titleMonth = newTitle;
  }

  onTimeSelected(event: { selectedTime: Date, events: any[] }) {
    let eventEmpty: boolean = (event.events !== undefined && event.events.length !== 0);
    console.log('Selected time: ' + event.selectedTime + ', hasEvents: ' + eventEmpty);
  }

  changeView() {
    let modeIndex:number = this.availableModes.indexOf(this.calendar.mode);
    modeIndex = (modeIndex + 1) % this.availableModes.length;
    this.calendar.mode = this.availableModes[modeIndex]
  }

}
