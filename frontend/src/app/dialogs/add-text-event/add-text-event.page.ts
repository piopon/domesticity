import { Component, OnInit } from "@angular/core";
import { ModalController } from "@ionic/angular";
import { Category } from "src/app/model/category.model";
import { CategoriesService } from "src/app/services/categories.service";
import { UsersService } from "src/app/services/users.service";

@Component({
  selector: "app-add-text-event",
  templateUrl: "./add-text-event.page.html",
  styleUrls: ["./add-text-event.page.scss"],
})
export class AddTextEventPage implements OnInit {
  protected availableUsers: String[];
  protected availableCategories: Category[];

  constructor(
    public modalController: ModalController,
    private categoriesService: CategoriesService,
    private usersService: UsersService
  ) {}

  ngOnInit() {
    this.updateUsers();
    this.updateCategories();
  }

  closeDialog():void {
    this.modalController.dismiss();
  }

  addEvent():void {
    console.log('todo: use eventService to add Event object');
    this.modalController.dismiss();
  }

  private updateUsers(): void {
    this.availableUsers = this.usersService.getTestUsers();
  }

  private updateCategories(): void {
    this.availableCategories = this.categoriesService.getTestCategories();
  }
}
