import { ComponentFixture, TestBed } from '@angular/core/testing';
import { DatetimeComponent } from './todo.component';

describe('DatetimeComponent', () => {
  let component: DatetimeComponent;
  let fixture: ComponentFixture<DateTime>;

  beforeEach(async(() => {
   TestBed.configureTestingModule({
      declarations: [
         DatetimeComponent
      ],
   }).compileComponents();
}));

  it('component should be created', async(() => {
    const fixture = TestBed.createComponent(DatetimeComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  });
}));

  it(`component should have title 'app-datetime'`, async(() => {
     const fixture = TestBed.createComponent(DatetimeComponent);
     const app = fixture.debugElement.componentInstance;
     expect(app.title).toEqual('app-datetime');
}));
