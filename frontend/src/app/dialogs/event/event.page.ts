import { Component, Input, OnInit } from '@angular/core';
import { ModalController } from '@ionic/angular';

@Component({
  selector: 'app-event',
  templateUrl: './event.page.html',
  styleUrls: ['./event.page.scss'],
})
export class EventPage implements OnInit {

  @Input() dayTime: Date
  @Input() dayEvents: any[]

  constructor(public modalController: ModalController) { }

  ngOnInit() {
    this.dayEvents = [
      {icon: 'american-football-outline', title: 'Test event 1 title'},
      {icon: 'bicycle-outline', title: 'Test event 2 title'},
      {icon: 'earth-outline', title: 'Test event 3 title'},
      {icon: 'game-controller-outline', title: 'Test event 4 title'},
      {icon: 'school-outline', title: 'Test event 5 title'},
    ];
  }

  close() {
    this.modalController.dismiss();
  }

}
