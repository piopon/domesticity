import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { UpdateTextEventPageRoutingModule } from './update-text-event-routing.module';

import { UpdateTextEventPage } from './update-text-event.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    UpdateTextEventPageRoutingModule
  ],
  declarations: [UpdateTextEventPage]
})
export class UpdateTextEventPageModule {}
