import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';
import { DayEventsPage } from './day-events.page';

describe('DayEventsPage', () => {
  let component: DayEventsPage;
  let fixture: ComponentFixture<DayEventsPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DayEventsPage ],
      imports: [IonicModule.forRoot()]
    }).compileComponents();

    fixture = TestBed.createComponent(DayEventsPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
