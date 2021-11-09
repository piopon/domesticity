import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';
import { IpcMessage } from '../model/ipc-message';

@Injectable({
  providedIn: 'root'
})
export class IpcMessagesService {
  private subject = new Subject<IpcMessage>();

  constructor() { }

  public sendMessage(message: IpcMessage) {
    this.subject.next(message);
  }

  public watch(): Observable<IpcMessage> {
    return this.subject.asObservable();
  }
}