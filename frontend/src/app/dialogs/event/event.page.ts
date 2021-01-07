import { Component, Input, OnInit } from '@angular/core';
import { ModalController } from '@ionic/angular';
import { Event } from 'src/app/model/event.model';
import { EventsService } from 'src/app/services/events.service';

@Component({
  selector: 'app-event',
  templateUrl: './event.page.html',
  styleUrls: ['./event.page.scss'],
})
export class EventPage implements OnInit {

  @Input() dayTime: Date
  @Input() dayEvents: Event[]

  constructor(public modalController: ModalController, private eventsService: EventsService) { }

  ngOnInit() {
    this.dayEvents = this.eventsService.getTestEvents();
  }

  close() {
    this.modalController.dismiss();
  }

  clearEvents() {
    this.dayEvents = [];
  }
}
