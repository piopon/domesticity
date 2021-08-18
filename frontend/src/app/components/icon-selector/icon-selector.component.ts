import { Component, OnInit } from "@angular/core";

@Component({
  selector: "icon-selector",
  templateUrl: "./icon-selector.component.html",
  styleUrls: ["./icon-selector.component.scss"],
})
export class IconSelectorComponent implements OnInit {
  private currentIcon: String;
  private iconStyle: String = "outline";
  private availableIcons: String[] = [
    "airplane", "balloon", "basketball", "beer", "bicycle", "book", "bonfire", "bowling-ball", "bus",
    "cafe", "car", "cart", "construct", "dice", "fast-food", "football", "game-controller", "school",
  ];

  constructor() {
    this.currentIcon = this.randomIcon();
  }

  ngOnInit() {}

  private randomIcon(): String {
    let randomIndex = Math.floor(Math.random() * this.availableIcons.length);
    return this.availableIcons[randomIndex] + "-" + this.iconStyle;
  }
}
