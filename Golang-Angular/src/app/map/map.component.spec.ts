import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MapComponent } from './map.component';

import {HttpClientModule} from '@angular/common/http';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { LeafletModule } from '@asymmetrik/ngx-leaflet';

describe('MapComponent', () => {
  let component: MapComponent;
  let fixture: ComponentFixture<MapComponent>;

  beforeEach(async () => {
   TestBed.configureTestingModule({
      declarations: [
         MapComponent
      ],
      imports: [HttpClientModule, FormsModule, BrowserModule, LeafletModule],
   }).compileComponents();
   
    fixture = TestBed.createComponent(MapComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
});

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  // it(`component should have title 'app-map'`, () => {
  //   const fixture = TestBed.createComponent(MapComponent);
  //    const app = fixture.debugElement.componentInstance;
  //    expect(app.title).toEqual('app-map');
  // });

});

