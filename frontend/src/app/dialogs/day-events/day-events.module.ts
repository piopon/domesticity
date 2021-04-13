import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { IonicModule } from '@ionic/angular';
import { DayEventsPage } from './day-events.page';
import { DayEventsPageRoutingModule } from './day-events-routing.module';
import { AddEventModule } from 'src/app/components/add-event/add-event.wrapper';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    DayEventsPageRoutingModule,
    AddEventModule
  ],
  declarations: [
    DayEventsPage,
  ]
})
export class DayEventsPageModule {}
