import { Component, OnInit } from '@angular/core';
import { ModalController } from '@ionic/angular';
import { AddTextEventPage } from 'src/app/dialogs/add-text-event/add-text-event.page';
import { Event } from 'src/app/model/event.model';
import { TextEventsService } from 'src/app/services/text-events.service';

@Component({
  selector: 'add-event-btn',
  templateUrl: './add-event-btn.component.html',
  styleUrls: ['./add-event-btn.component.scss'],
})
export class AddEventComponent implements OnInit {

  constructor(public modalController: ModalController,
    private textEventsService: TextEventsService) {}

  ngOnInit() {}

  getTextEventState(): boolean {
    return this.textEventsService.isOnline();
  }

  async openNewTextEventDialog() {
    if (this.getTextEventState()) {
      const modal = await this.modalController.create({
        component: AddTextEventPage,
      });
      return await modal.present();
    }
  }

  getTodoEventState(): boolean {
    return false;
  }

  async openNewTodoEventDialog() {
    if (this.getTodoEventState()) {
      const modal = await this.modalController.create({
        component: AddTextEventPage,
      });
      return await modal.present();
    }
  }

  getShopEventState(): boolean {
    return false;
  }

  async openNewShopEventDialog() {
    if (this.getShopEventState()) {
      const modal = await this.modalController.create({
        component: AddTextEventPage,
      });
      return await modal.present();
    }
  }
}
