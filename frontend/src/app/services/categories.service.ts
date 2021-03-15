import { Injectable } from '@angular/core';
import { Category } from '../model/category.model';

@Injectable({
  providedIn: 'root'
})
export class CategoriesService {

  constructor() { }

  getTestCategories():Category[] {
    return [
      {
        id: '0',
        name: 'Red',
        color: '#ff0000',
      },
      {
        id: '1',
        name: 'Green',
        color: '#00ff00',
      },
      {
        id: '2',
        name: 'Blue',
        color: '#0000ff',
      },
      {
        id: '3',
        name: 'Yellow',
        color: '#ffff00',
      },
      {
        id: '4',
        name: 'Orange',
        color: '#ffa200',
      },
      {
        id: '5',
        name: 'Pink',
        color: '#ff00ff',
      },
    ];
  }
}
