import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MapComponent } from './map.component';

describe('MapComponent', () => {
  let component: MapComponent;
  let fixture: ComponentFixture<MapComponent>;

  beforeEach(async(() => {
   TestBed.configureTestingModule({
      declarations: [
         MapComponent
      ],
   }).compileComponents();
}));

  it('component should be created', async(() => {
    const fixture = TestBed.createComponent(MapComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  });
}));

  it(`component should have title 'app-map'`, async(() => {
     const fixture = TestBed.createComponent(MapComponent);
     const app = fixture.debugElement.componentInstance;
     expect(app.title).toEqual('app-map');
}));
