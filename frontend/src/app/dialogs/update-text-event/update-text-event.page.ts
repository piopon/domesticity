import { FormControl, FormGroup, Validators } from "@angular/forms";
import { Component, Input, OnInit } from '@angular/core';
import { AlertController, ModalController, ToastController } from "@ionic/angular";
import { Category } from "src/app/model/category.model";
import { Event } from "src/app/model/event.model";
import { CategoriesService } from "src/app/services/categories.service";
import { TextEventsService } from "src/app/services/text-events.service";
import { UsersService } from "src/app/services/users.service";

@Component({
  selector: 'app-update-text-event',
  templateUrl: './update-text-event.page.html',
  styleUrls: ['./update-text-event.page.scss'],
})
export class UpdateTextEventPage implements OnInit {
  @Input() event: Event;

  availableIcons: string[];
  availableUsers: String[];
  availableCategories: Category[];

  private eventForm: FormGroup;
  private tempDateStart: string;
  private tempDateStop: string;

  constructor(
    public modalController: ModalController,
    private alertController: AlertController,
    private toastController: ToastController,
    private categoriesService: CategoriesService,
    private usersService: UsersService,
    private eventService: TextEventsService
  ) {}

  ngOnInit() {
    this.availableIcons = Event.getAvailableIcons();
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

  async updateEvent(): Promise<void> {
    return;
  }

  hasError(widget: string): boolean {
    return this.eventForm.get(widget).hasError("required") && this.eventForm.get(widget).touched;
  }

  iconUpdated(newIcon: string) {
    this.event.icon = newIcon;
  }

  private updateUsers(): void {
    this.availableUsers = this.usersService.getTestUsers();
  }

  private updateCategories(): void {
    this.availableCategories = this.categoriesService.getTestCategories();
  }
}
