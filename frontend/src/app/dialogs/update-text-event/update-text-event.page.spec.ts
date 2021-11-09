import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';

import { UpdateTextEventPage } from './update-text-event.page';

describe('UpdateTextEventPage', () => {
  let component: UpdateTextEventPage;
  let fixture: ComponentFixture<UpdateTextEventPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ UpdateTextEventPage ],
      imports: [IonicModule.forRoot()]
    }).compileComponents();

    fixture = TestBed.createComponent(UpdateTextEventPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
