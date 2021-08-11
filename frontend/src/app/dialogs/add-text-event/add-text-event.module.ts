import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { AddTextEventPageRoutingModule } from './add-text-event-routing.module';

import { AddTextEventPage } from './add-text-event.page';
import { IconSelectorModule } from 'src/app/components/icon-selector/icon-selector.wrapper';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    AddTextEventPageRoutingModule,
    IconSelectorModule,
    ReactiveFormsModule
  ],
  declarations: [AddTextEventPage]
})
export class AddTextEventPageModule {}
