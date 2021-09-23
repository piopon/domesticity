import { formatDate } from "@angular/common";
import { Component, Input, OnInit } from "@angular/core";
import { ModalController } from "@ionic/angular";
import { Subscription } from "rxjs";
import { Event } from "src/app/model/event.model";
import { TimeSpan } from "src/app/model/timespan.model";
import { TextEventsService } from "src/app/services/text-events.service";
import { UpdateTextEventPage } from "../update-text-event/update-text-event.page";

@Component({
  selector: "app-day-events",
  templateUrl: "./day-events.page.html",
  styleUrls: ["./day-events.page.scss"],
})
export class DayEventsPage implements OnInit {
  @Input() dayTime: Date;

  private dayEvents: Event[] = [];
  private todayString: string;
  private selectedEventIndex: number;
  private subscription: Subscription;
  private static readonly NO_EVENT_SELECTED: number = -1;

  constructor(public modalController: ModalController, private textEventsService: TextEventsService) {
    this.subscription = this.textEventsService.watch().subscribe(_ => this.getAllTextEvents(this.todayString));
  }

  ngOnInit() {
    this.selectedEventIndex = DayEventsPage.NO_EVENT_SELECTED;
    this.todayString = formatDate(this.dayTime, "yyyy-MM-dd", "en");
    this.getAllTextEvents(this.todayString);
  }

  ngOnDestroy() {
    this.subscription.unsubscribe();
  }

  closeDialog() {
    this.modalController.dismiss();
  }

  deleteAllEvents() {
    this.dayEvents.forEach((event) => {
      this.textEventsService.deleteEvent(event._id).subscribe(() => {
        console.log("Removed events!");
      });
    });
    this.dayEvents = [];
  }

  deleteEvent(eventIndex: number) {
    let selectedEvent = this.dayEvents[eventIndex];
    this.textEventsService.deleteEvent(selectedEvent._id).subscribe(() => {
      console.log("Removed event!");
    });
    this.dayEvents.splice(eventIndex, 1);
    this.selectedEventIndex = DayEventsPage.NO_EVENT_SELECTED;
  }

  async updateEvent(eventIndex: number): Promise<void> {
    if (this.textEventsService.isOnline()) {
      const modal = await this.modalController.create({
        component: UpdateTextEventPage,
        componentProps: {
          event: this.dayEvents[eventIndex],
        },
      });
      return await modal.present();
    }
  }

  isEventSelected(eventIndex: number): boolean {
    return this.selectedEventIndex === eventIndex;
  }

  toggleEventSelection(eventIndex: number): void {
    this.selectedEventIndex = this.isEventSelected(eventIndex) ? DayEventsPage.NO_EVENT_SELECTED : eventIndex;
  }

  private getAllTextEvents(dayString: string): void {
    this.dayEvents = [];
    this.textEventsService.getEventsByDateStart(dayString).subscribe((events) => {
      events?.forEach((event) => {
        let eventDate = new TimeSpan(new Date(event.date.start), new Date(event.date.stop));
        this.dayEvents.push(
          new Event(event._id, event.title, event.icon, event.owner, eventDate, event.category, event.content)
        );
      });
    });
  }
}
