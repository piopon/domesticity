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

  private visibleDetails: number[]

  constructor(public modalController: ModalController, private eventsService: EventsService) { }

  ngOnInit() {
    this.dayEvents = this.eventsService.getTestEvents();
    this.visibleDetails = [];
  }

  closeDialog() {
    this.modalController.dismiss();
  }

  clearEvents() {
    this.dayEvents = [];
  }

  isDetailed(eventIndex:number):boolean {
    return this.visibleDetails.indexOf(eventIndex) !== -1;
  }

  toggleDetails(eventIndex:number):void {
    if (this.isDetailed(eventIndex)) {
      const index:number = this.visibleDetails.indexOf(eventIndex);
      this.visibleDetails.splice(index, 1);
    } else {
      this.visibleDetails.push(eventIndex);
    }
  }
}
