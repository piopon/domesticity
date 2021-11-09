import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { UpdateTextEventPageRoutingModule } from './update-text-event-routing.module';

import { UpdateTextEventPage } from './update-text-event.page';
import { IconSelectorModule } from 'src/app/components/icon-selector/icon-selector.wrapper';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    UpdateTextEventPageRoutingModule,
    IconSelectorModule,
    ReactiveFormsModule
  ],
  declarations: [UpdateTextEventPage]
})
export class UpdateTextEventPageModule {}
