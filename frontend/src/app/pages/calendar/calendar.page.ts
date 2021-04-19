import { Component, OnInit, ViewChild } from '@angular/core';
import { ModalController } from '@ionic/angular';
import { CalendarComponent } from 'ionic2-calendar';
import { DayEventsPage } from 'src/app/dialogs/day-events/day-events.page';

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

  constructor(public modalController: ModalController) { }

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

  async onTimeSelected(event: { selectedTime: Date, events: any[] }) {
    const modal = await this.modalController.create({
      component: DayEventsPage,
      componentProps: {
        'dayTime': event.selectedTime,
        'dayEvents': event.events,
      }
    });
    return await modal.present();
  }

  changeView() {
    let modeIndex:number = this.availableModes.indexOf(this.calendar.mode);
    modeIndex = (modeIndex + 1) % this.availableModes.length;
    this.calendar.mode = this.availableModes[modeIndex]
  }

}
