import { Component, OnInit, Output, EventEmitter, Input } from "@angular/core";
import { ActionSheetController } from "@ionic/angular";

@Component({
  selector: "icon-selector",
  templateUrl: "./icon-selector.component.html",
  styleUrls: ["./icon-selector.component.scss"],
})
export class IconSelectorComponent implements OnInit {
  @Input() currentIcon: string = "";
  @Input() availableIcons: string[] = [];
  @Output() selectedIcon = new EventEmitter<string>();

  private iconStyle: string = "outline";
  constructor(public actionSheetController: ActionSheetController) {}

  ngOnInit() {}

  selectIcon(): void {
    this.presentIcons();
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
            this.selectedIcon.emit(this.currentIcon);
          }
        };
      }),
    });
    await actionSheet.present();
  }
}
