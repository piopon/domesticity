import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class IpcMessagesService {
  private subject = new Subject<string>();

  constructor() { }

  public sendMessage(message: string) {
    this.subject.next(message);
  }

  public watch(): Observable<string> {
    return this.subject.asObservable();
  }
}
