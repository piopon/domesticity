import { FormControl, FormGroup, Validators } from "@angular/forms";
import { Component, Input, OnInit } from "@angular/core";
import { AlertController, ModalController, ToastController } from "@ionic/angular";
import { Category } from "src/app/model/category.model";
import { Event } from "src/app/model/event.model";
import { CategoriesService } from "src/app/services/categories.service";
import { TextEventsService } from "src/app/services/text-events.service";
import { UsersService } from "src/app/services/users.service";
import { IpcMessagesService } from "src/app/services/ipc-messages.service";
import { IpcMessage } from "src/app/model/ipc-message";

@Component({
  selector: "app-add-text-event",
  templateUrl: "./add-text-event.page.html",
  styleUrls: ["./add-text-event.page.scss"],
})
export class AddTextEventPage implements OnInit {
  availableIcons: string[];
  availableUsers: String[];
  availableCategories: Category[];

  private event: Event;
  private eventForm: FormGroup;
  private tempDateStart: string;
  private tempDateStop: string;

  constructor(
    public modalController: ModalController,
    private alertController: AlertController,
    private toastController: ToastController,
    private categoriesService: CategoriesService,
    private usersService: UsersService,
    private eventService: TextEventsService,
    private ipcMessagesService: IpcMessagesService
  ) {}

  ngOnInit() {
    this.event = Event.empty();
    this.tempDateStart = this.event.date.start.toISOString();
    this.tempDateStop = this.event.date.stop.toISOString();
    this.updateUsers();
    this.updateCategories();
    this.availableIcons = Event.getAvailableIcons();
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
    if (await this.checkEventErrors()) {
      return;
    }
    this.eventService.addEvent(this.event).subscribe(async (responseEvent) => {
      const toast = await this.toastController.create({
        color: responseEvent.id !== "" ? "success" : "danger",
        message: responseEvent.id !== "" ? "Event successfully added." : "Error while adding event.",
        duration: 2000,
      });
      toast.present();
      this.ipcMessagesService.sendMessage(IpcMessage.newEvent(this.tempDateStart));
    });
    this.modalController.dismiss();
  }

  hasError(widget: string): boolean {
    return this.eventForm.get(widget).hasError("required") && this.eventForm.get(widget).touched;
  }

  iconUpdated(newIcon: string) {
    this.event.icon = newIcon;
  }

  private async checkEventErrors(): Promise<boolean> {
    let errors: string[] = this.event.verify();
    if (errors.length > 0) {
      const alert = await this.alertController.create({
        header: "error",
        subHeader: "cannot create event",
        message: "event has errors:<br>- " + errors.join("<br>- "),
        buttons: ["OK"],
      });
      alert.present();
      return true;
    }
    return false;
  }

  private updateUsers(): void {
    this.availableUsers = this.usersService.getTestUsers();
  }

  private updateCategories(): void {
    this.availableCategories = this.categoriesService.getTestCategories();
  }
}
