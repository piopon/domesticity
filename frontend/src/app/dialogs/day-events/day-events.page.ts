import { formatDate } from "@angular/common";
import { Component, Input, OnInit } from "@angular/core";
import { ModalController } from "@ionic/angular";
import { Category } from "src/app/model/category.model";
import { Event } from "src/app/model/event.model";
import { TimeSpan } from "src/app/model/timespan.model";
import { CategoriesService } from "src/app/services/categories.service";
import { TextEventsService } from "src/app/services/text-events.service";
import { UsersService } from "src/app/services/users.service";

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
    this.dayEvents = [];
    this.visibleDetails = [];
    this.todayString = formatDate(this.dayTime, "yyyy-MM-dd", "en");
    this.updateTextEvents(this.todayString);
    this.updateUsers();
    this.updateCategories();
  }

  closeDialog() {
    this.modalController.dismiss();
  }

  clearEvents() {
    this.dayEvents = [];
  }

  isDetailed(eventIndex: number): boolean {
    return this.visibleDetails.indexOf(eventIndex) !== -1;
  }

  toggleDetails(eventIndex: number): void {
    if (this.isDetailed(eventIndex)) {
      const index: number = this.visibleDetails.indexOf(eventIndex);
      this.visibleDetails.splice(index, 1);
    } else {
      this.visibleDetails.push(eventIndex);
    }
  }

  private updateTextEvents(dayString: string): void {
    this.textEventsService.getEventsByDateStart(dayString).subscribe((events) => {
      events?.forEach((event) => {
        let eventDate = new TimeSpan(new Date(event.date.start), new Date(event.date.stop));
        this.dayEvents.push(
          new Event(event.id, event.title, event.icon, event.owner, eventDate, event.category, event.content)
        );
      });
    });
  }

  private updateUsers(): void {
    this.availableUsers = this.usersService.getTestUsers();
  }

  private updateCategories(): void {
    this.availableCategories = this.categoriesService.getTestCategories();
  }
}
