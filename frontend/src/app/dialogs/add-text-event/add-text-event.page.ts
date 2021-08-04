import { Component, OnInit } from "@angular/core";
import { ModalController } from "@ionic/angular";
import { Category } from "src/app/model/category.model";
import { CategoriesService } from "src/app/services/categories.service";

@Component({
  selector: 'app-add-text-event',
  templateUrl: './add-text-event.page.html',
  styleUrls: ['./add-text-event.page.scss'],
})
export class AddTextEventPage implements OnInit {
  protected availableCategories: Category[];

  constructor(
    public modalController: ModalController,
    private categoriesService: CategoriesService,
  ) {}

  ngOnInit() {
    this.updateCategories();
  }

  closeDialog() {
    this.modalController.dismiss();
  }

  private updateCategories(): void {
    this.availableCategories = this.categoriesService.getTestCategories();
    console.log('got categories: ' + this.availableCategories);
  }
}
