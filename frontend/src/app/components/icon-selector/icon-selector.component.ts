import { Component, OnInit } from "@angular/core";
import { ActionSheetController } from "@ionic/angular";

@Component({
  selector: "icon-selector",
  templateUrl: "./icon-selector.component.html",
  styleUrls: ["./icon-selector.component.scss"],
})
export class IconSelectorComponent implements OnInit {
  private currentIcon: string;
  private iconStyle: string = "outline";
  private availableIcons: string[] = [
    "airplane", "basketball", "beer", "bicycle", "book", "bonfire", "bus", "cafe",
    "car", "cart", "construct", "dice", "fast-food", "football", "game-controller", "school",
  ];

  constructor(public actionSheetController: ActionSheetController) {
    this.currentIcon = this.randomIcon();
  }

  ngOnInit() {}

  selectIcon(): void {
    this.presentIcons();
  }

  private randomIcon(): string {
    let randomIndex = Math.floor(Math.random() * this.availableIcons.length);
    return this.availableIcons[randomIndex] + "-" + this.iconStyle;
  }

  private async presentIcons() {
    const actionSheet = await this.actionSheetController.create({
      header: "Icons",
      buttons: this.availableIcons.map((iconName) => {
        return {
          text: iconName,
          icon: iconName,
          handler: () => {
            this.currentIcon = iconName + "-" + this.iconStyle;
          }
        };
      }),
    });
    await actionSheet.present();
  }
}
