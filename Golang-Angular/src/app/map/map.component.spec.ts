import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Map } from './todo.component';

describe('Map', () => {
  let component: Map;
  let fixture: ComponentFixture<Map>;

  beforeEach(async(() => {
   TestBed.configureTestingModule({
      declarations: [
         Map
      ],
   }).compileComponents();
}));

  it('component should be created', async(() => {
    const fixture = TestBed.createComponent(Map);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  });
}));

  it(`component should have title 'app-map'`, async(() => {
     const fixture = TestBed.createComponent(Map);
     const app = fixture.debugElement.componentInstance;
     expect(app.title).toEqual('app-map');
}));
