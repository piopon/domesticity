import { Injectable } from '@angular/core';
import { Event } from '../model/event.model';

@Injectable({
  providedIn: 'root'
})
export class EventsService {

  constructor() { }

  getTestEvents():Event[] {
    return [
      {
        id: '0',
        icon: 'american-football-outline',
        title: 'Test event 1 title',
        category: 'Red',
        content: 'This is a content for event 1',
        owner: 'Admin',
        occurence: {
          start: new Date(),
          stop: new Date(),
        }
      },
      {
        id: '1',
        icon: 'bicycle-outline',
        title: 'Test event 2 title',
        category: 'Blue',
        content: 'This is a content for event 2',
        owner: 'Admin',
        occurence: {
          start: new Date(),
          stop: new Date(),
        }
      },
      {
        id: '2',
        icon: 'earth-outline',
        title: 'Test event 3 title',
        category: 'Blue',
        content: 'This is a content for event 3',
        owner: 'Admin',
        occurence: {
          start: new Date(),
          stop: new Date(),
        }
      },
      {
        id: '3',
        icon: 'game-controller-outline',
        title: 'Test event 4 title',
        category: 'Blue',
        content: 'This is a content for event 4',
        owner: 'Admin',
        occurence: {
          start: new Date(),
          stop: new Date(),
        }
      },
      {
        id: '4',
        icon: 'school-outline',
        title: 'Test event 5 title',
        category: 'Blue',
        content: 'This is a content for event 5',
        owner: 'Admin',
        occurence: {
          start: new Date(),
          stop: new Date(),
        }
      }
    ];
  }
}
