import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { AddTextEventPage } from './add-text-event.page';

const routes: Routes = [
  {
    path: '',
    component: AddTextEventPage
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class AddTextEventPageRoutingModule {}
