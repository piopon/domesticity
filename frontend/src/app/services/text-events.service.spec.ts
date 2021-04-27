import { TestBed } from '@angular/core/testing';

import { TextEventsService } from './text-events.service';

describe('EventsService', () => {
  let service: TextEventsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TextEventsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
