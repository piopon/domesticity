import { Component, OnInit } from "@angular/core";

@Component({
  selector: "icon-selector",
  templateUrl: "./icon-selector.component.html",
  styleUrls: ["./icon-selector.component.scss"],
})
export class IconSelectorComponent implements OnInit {
  currentIcon: String;

  constructor() {
    this.currentIcon = "accessibility-outline";
  }

  ngOnInit() {}
}
