import { FormControl, FormGroup, Validators } from "@angular/forms";
import { Component, OnInit } from "@angular/core";
import { AlertController, ModalController } from "@ionic/angular";
import { Category } from "src/app/model/category.model";
import { Event } from "src/app/model/event.model";
import { CategoriesService } from "src/app/services/categories.service";
import { TextEventsService } from "src/app/services/text-events.service";
import { UsersService } from "src/app/services/users.service";

@Component({
  selector: "app-add-text-event",
  templateUrl: "./add-text-event.page.html",
  styleUrls: ["./add-text-event.page.scss"],
})
export class AddTextEventPage implements OnInit {
  protected availableUsers: String[];
  protected availableCategories: Category[];
  private event: Event;
  private eventForm: FormGroup;
  private tempDateStart: string;
  private tempDateStop: string;

  constructor(
    public modalController: ModalController,
    private alertController: AlertController,
    private categoriesService: CategoriesService,
    private usersService: UsersService,
    private eventService: TextEventsService
  ) {}

  ngOnInit() {
    this.event = Event.empty();
    this.tempDateStart = this.event.date.start.toISOString();
    this.tempDateStop = this.event.date.stop.toISOString();
    this.updateUsers();
    this.updateCategories();
    this.eventForm = new FormGroup({
      title: new FormControl("", Validators.required),
      category: new FormControl("", Validators.required),
      owner: new FormControl("", Validators.required),
      content: new FormControl("", Validators.required),
      start: new FormControl("", Validators.required),
      stop: new FormControl("", Validators.required),
    });
    this.eventForm.enable();
  }

  closeDialog(): void {
    this.modalController.dismiss();
  }

  async addEvent(): Promise<void> {
    this.event.date.start = new Date(this.tempDateStart);
    this.event.date.stop = new Date(this.tempDateStop);
    if (!this.event.validate()) {
      const alert = await this.alertController.create({
        header: "error",
        subHeader: "cannot create event",
        message: "event has errors",
        buttons: ["OK"],
      });
      alert.present();
      return;
    }
    console.log(this.event);
    this.modalController.dismiss();
  }

  private updateUsers(): void {
    this.availableUsers = this.usersService.getTestUsers();
  }

  private updateCategories(): void {
    this.availableCategories = this.categoriesService.getTestCategories();
  }
}
