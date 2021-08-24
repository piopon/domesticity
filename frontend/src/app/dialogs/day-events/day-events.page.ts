import { formatDate } from "@angular/common";
import { Component, Input, OnInit } from "@angular/core";
import { ModalController } from "@ionic/angular";
import { Category } from "src/app/model/category.model";
import { Event } from "src/app/model/event.model";
import { TimeSpan } from "src/app/model/timespan.model";
import { CategoriesService } from "src/app/services/categories.service";
import { TextEventsService } from "src/app/services/text-events.service";
import { UsersService } from "src/app/services/users.service";
import { AddTextEventPage } from "../add-text-event/add-text-event.page";

@Component({
  selector: "app-day-events",
  templateUrl: "./day-events.page.html",
  styleUrls: ["./day-events.page.scss"],
})
export class DayEventsPage implements OnInit {
  @Input() dayTime: Date;
  @Input() dayEvents: Event[];

  protected availableUsers: String[];
  protected availableCategories: Category[];
  private visibleDetails: number[];
  private todayString: string;

  constructor(
    public modalController: ModalController,
    private textEventsService: TextEventsService,
    private categoriesService: CategoriesService,
    private usersService: UsersService
  ) {}

  ngOnInit() {
    this.todayString = formatDate(this.dayTime, "yyyy-dd-MM", "en");
    this.getAllTextEvents(this.todayString);
  }

  closeDialog() {
    this.modalController.dismiss();
  }

  clearEvents() {
    this.dayEvents = [];
  }

  async eventDetails(eventIndex: number): Promise<void> {
    if (this.textEventsService.isOnline()) {
      const modal = await this.modalController.create({
        component: AddTextEventPage,
        componentProps: {
          'event': this.dayEvents[eventIndex],
        }
      });
      return await modal.present();
    }
  }

  private getAllTextEvents(dayString: string): void {
    this.clearEvents();
    this.textEventsService.getEventsByDateStart(dayString).subscribe((events) => {
      events?.forEach((event) => {
        let eventDate = new TimeSpan(new Date(event.date.start), new Date(event.date.stop));
        this.dayEvents.push(
          new Event(event.id, event.title, event.icon, event.owner, eventDate, event.category, event.content)
        );
      });
    });
  }
}
