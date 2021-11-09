import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { UpdateTextEventPage } from './update-text-event.page';

const routes: Routes = [
  {
    path: '',
    component: UpdateTextEventPage
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class UpdateTextEventPageRoutingModule {}
