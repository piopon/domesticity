import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class CategoriesService {

  constructor() { }

  getTestCategories():String[] {
    return [
      'Red',
      'Green',
      'Blue',
      'Yellow',
      'Orange',
      'Pink',
    ];
  }
}
