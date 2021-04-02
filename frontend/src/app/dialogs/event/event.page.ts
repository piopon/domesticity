import { Component, Input, OnInit } from '@angular/core';
import { ModalController } from '@ionic/angular';
import { Category } from 'src/app/model/category.model';
import { Event } from 'src/app/model/event.model';
import { CategoriesService } from 'src/app/services/categories.service';
import { EventsService } from 'src/app/services/events.service';
import { UsersService } from 'src/app/services/users.service';

@Component({
  selector: 'app-event',
  templateUrl: './event.page.html',
  styleUrls: ['./event.page.scss'],
})
export class EventPage implements OnInit {

  @Input() dayTime: Date;
  @Input() dayEvents: Event[];

  protected availableUsers: String[];
  protected availableCategories: Category[];
  private visibleDetails: number[];

  constructor(public modalController: ModalController,
    private eventsService: EventsService,
    private categoriesService: CategoriesService,
    private usersService: UsersService) { }

  ngOnInit() {
    this.eventsService.getEventsByOwner("ponio").subscribe(
      e => this.dayEvents = e
    );
    this.availableUsers = this.usersService.getTestUsers();
    this.availableCategories = this.categoriesService.getTestCategories();
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
