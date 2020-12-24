import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-event',
  templateUrl: './event.page.html',
  styleUrls: ['./event.page.scss'],
})
export class EventPage implements OnInit {

  @Input() dayTime: Date
  @Input() dayEvents: any[]

  constructor() { }

  ngOnInit() {
  }

}
