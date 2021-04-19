import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';

import { AddTextEventPage } from './add-text-event.page';

describe('AddTextEventPage', () => {
  let component: AddTextEventPage;
  let fixture: ComponentFixture<AddTextEventPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AddTextEventPage ],
      imports: [IonicModule.forRoot()]
    }).compileComponents();

    fixture = TestBed.createComponent(AddTextEventPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
