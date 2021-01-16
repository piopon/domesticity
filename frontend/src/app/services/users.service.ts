import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class UsersService {

  constructor() { }

  getTestUsers():String[] {
    return [
      'Admin',
      'User',
      'Moderator',
    ];
  }
}
