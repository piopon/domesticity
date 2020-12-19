import { Component, OnInit, ViewChild } from '@angular/core';
import { CalendarComponent } from 'ionic2-calendar';

@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.page.html',
  styleUrls: ['./calendar.page.scss'],
})
export class CalendarPage implements OnInit {
  eventSource = [];
  titleMonth: string;

  calendar = {
    mode: 'month',
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

}
