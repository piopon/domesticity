import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { CalendarPageRoutingModule } from './calendar-routing.module';

import { CalendarPage } from './calendar.page';
import { NgCalendarModule } from 'ionic2-calendar';
import { EventPage } from 'src/app/dialogs/event/event.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    CalendarPageRoutingModule,
    NgCalendarModule,
    EventPage
  ],
  declarations: [CalendarPage]
})
export class CalendarPageModule {}
