import { Component, OnInit } from '@angular/core';
import { ModalController } from '@ionic/angular';
import { AddTextEventPage } from 'src/app/dialogs/add-text-event/add-text-event.page';
import { EventsService } from 'src/app/services/events.service';

@Component({
  selector: 'add-event-btn',
  templateUrl: './add-event-btn.component.html',
  styleUrls: ['./add-event-btn.component.scss'],
})
export class AddEventComponent implements OnInit {

  constructor(public modalController: ModalController,
    private eventsService: EventsService) { }

  ngOnInit() {}

  getTextEventState(): string {
    return this.eventsService.isOnline() ? "tertiary" : "medium";
  }

  async openNewTextEventDialog() {
    const modal = await this.modalController.create({
      component: AddTextEventPage,
    });
    return await modal.present();
  }
}
