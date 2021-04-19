import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'add-event-btn',
  templateUrl: './add-event-btn.component.html',
  styleUrls: ['./add-event-btn.component.scss'],
})
export class AddEventComponent implements OnInit {

  constructor() { }

  ngOnInit() {}

  openNewTextEventDialog() {
    console.log('open text-event dialog');
  }
}
