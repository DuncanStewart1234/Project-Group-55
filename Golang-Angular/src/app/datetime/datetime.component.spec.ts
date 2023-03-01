import { ComponentFixture, TestBed } from '@angular/core/testing';
import { DateTime } from './todo.component';

describe('DateTime', () => {
  let component: DateTime;
  let fixture: ComponentFixture<DateTime>;

  beforeEach(async(() => {
   TestBed.configureTestingModule({
      declarations: [
         DateTime
      ],
   }).compileComponents();
}));

  it('component should be created', async(() => {
    const fixture = TestBed.createComponent(DateTime);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  });
}));

  it(`component should have title 'app-datetime'`, async(() => {
     const fixture = TestBed.createComponent(DateTime);
     const app = fixture.debugElement.componentInstance;
     expect(app.title).toEqual('app-datetime');
}));
