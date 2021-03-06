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
          start: new Date(2021, 1, 1, 10, 0, 0, 0),
          stop: new Date(2021, 1, 1, 10, 1, 0, 0),
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
          start: new Date(2021, 2, 2, 12, 0, 0, 0),
          stop: new Date(2021, 2, 2, 12, 1, 0, 0),
        }
      },
      {
        id: '2',
        icon: 'earth-outline',
        title: 'Test event 3 title',
        category: 'Green',
        content: 'This is a content for event 3',
        owner: 'User',
        occurence: {
          start: new Date(2021, 3, 3, 13, 0, 0, 0),
          stop: new Date(2021, 3, 3, 13, 1, 0, 0),
        }
      },
      {
        id: '3',
        icon: 'game-controller-outline',
        title: 'Test event 4 title',
        category: 'Yellow',
        content: 'This is a content for event 4',
        owner: 'Moderator',
        occurence: {
          start: new Date(2021, 4, 4, 14, 0, 0, 0),
          stop: new Date(2021, 4, 4, 14, 1, 0, 0),
        }
      },
      {
        id: '4',
        icon: 'school-outline',
        title: 'Test event 5 title',
        category: 'Green',
        content: 'This is a content for event 5',
        owner: 'User',
        occurence: {
          start: new Date(2021, 5, 5, 15, 0, 0, 0),
          stop: new Date(2021, 5, 5, 15, 1, 0, 0),
        }
      }
    ];
  }
}
