import { Component, OnInit } from '@angular/core';
import { ModalController } from '@ionic/angular';

@Component({
  selector: 'app-add-text-event',
  templateUrl: './add-text-event.page.html',
  styleUrls: ['./add-text-event.page.scss'],
})
export class AddTextEventPage implements OnInit {

  constructor(public modalController: ModalController) { }

  ngOnInit() {
  }

  closeDialog() {
    this.modalController.dismiss();
  }

}
