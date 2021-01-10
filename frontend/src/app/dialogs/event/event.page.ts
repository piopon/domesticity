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

  private visibleInfoEvents: number[]

  constructor(public modalController: ModalController, private eventsService: EventsService) { }

  ngOnInit() {
    this.dayEvents = this.eventsService.getTestEvents();
    this.visibleInfoEvents = [];
  }

  closeDialog() {
    this.modalController.dismiss();
  }

  clearEvents() {
    this.dayEvents = [];
  }

  isEventOpened(eventIndex:number):boolean {
    return this.visibleInfoEvents.indexOf(eventIndex) !== -1;
  }

  toggleEventInfo(eventIndex:number):void {
    if (this.isEventOpened(eventIndex)) {
      const index:number = this.visibleInfoEvents.indexOf(eventIndex);
      this.visibleInfoEvents.splice(index, 1);
    } else {
      this.visibleInfoEvents.push(eventIndex);
    }
  }
}
