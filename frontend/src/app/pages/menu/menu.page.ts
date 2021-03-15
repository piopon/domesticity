import { Component, OnInit } from '@angular/core';
import { Router, RouterEvent } from '@angular/router';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.page.html',
  styleUrls: ['./menu.page.scss'],
})
export class MenuPage implements OnInit {
  pages = [
    {
      title: 'calendar',
      icon: 'calendar-outline',
      url: '/menu/calendar'
    },
    {
      title: 'user',
      icon: 'person-circle-outline',
      url: '/menu/user'
    }
  ];

  selectedPath = '';

  constructor(private router:Router) {
    this.router.events.subscribe((event: RouterEvent) => {
      this.selectedPath = event.url;
    });
  }

  ngOnInit() {
  }

}
