import { ComponentFixture, TestBed } from '@angular/core/testing';
import { DateTimeComponent } from './todo.component';

describe('DateTimeComponent', () => {
  let component: DateTimeComponent;
  let fixture: ComponentFixture<DateTime>;

  beforeEach(async(() => {
   TestBed.configureTestingModule({
      declarations: [
         DateTimeComponent
      ],
   }).compileComponents();
}));

  it('component should be created', async(() => {
    const fixture = TestBed.createComponent(DateTimeComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  });
}));

  it(`component should have title 'app-datetime'`, async(() => {
     const fixture = TestBed.createComponent(DateTimeComponent);
     const app = fixture.debugElement.componentInstance;
     expect(app.title).toEqual('app-datetime');
}));
