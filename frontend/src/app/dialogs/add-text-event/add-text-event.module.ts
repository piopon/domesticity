import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { AddTextEventPageRoutingModule } from './add-text-event-routing.module';

import { AddTextEventPage } from './add-text-event.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    AddTextEventPageRoutingModule
  ],
  declarations: [AddTextEventPage]
})
export class AddTextEventPageModule {}
