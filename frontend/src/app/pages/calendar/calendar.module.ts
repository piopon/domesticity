import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { IonicModule } from '@ionic/angular';
import { CalendarPageRoutingModule } from './calendar-routing.module';
import { CalendarPage } from './calendar.page';
import { NgCalendarModule } from 'ionic2-calendar';
import { DayEventsPage } from 'src/app/dialogs/day-events/day-events.page';
import { AddEventModule } from 'src/app/components/add-event-btn/add-event.wrapper';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    CalendarPageRoutingModule,
    NgCalendarModule,
    AddEventModule
  ],
  declarations: [
    CalendarPage,
    DayEventsPage,
  ]
})
export class CalendarPageModule {}
