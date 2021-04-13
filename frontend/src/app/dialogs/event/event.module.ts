import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { EventPageRoutingModule } from './event-routing.module';

import { EventPage } from './event.page';
import { AddEventModule } from 'src/app/components/add-event/add-event.module';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    EventPageRoutingModule,
    AddEventModule
  ],
  declarations: [
    EventPage,
  ]
})
export class EventPageModule {}
