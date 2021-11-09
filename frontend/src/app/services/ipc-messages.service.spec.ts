import { TestBed } from '@angular/core/testing';

import { IpcMessagesService } from './ipc-messages.service';

describe('IpcMessagesService', () => {
  let service: IpcMessagesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(IpcMessagesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
