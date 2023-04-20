import { TestBed } from '@angular/core/testing';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { AppComponent } from './app.component';
import { DatetimeComponent } from './datetime/datetime.component';
import { TodoComponent } from './todo/todo.component';
import { NoteComponent } from './note/note.component';

import { TodoService } from './todo.service';
import { NoteService } from './note.service';
import { FormsModule } from '@angular/forms';
import { MapComponent } from './map/map.component';

import { LoginService } from './login.service'
import { RegisterService } from './register.service'

import { LoginComponent } from './login/login.component'
import { RegisterComponent } from './register/register.component'

import { LeafletModule } from '@asymmetrik/ngx-leaflet';
import { EventComponent } from './event/event.component';
import { EventService } from './event.service';
// import "leaflet/dist/leaflet.css";

import { FullCalendarModule } from '@fullcalendar/angular';

import { WeatherComponent } from './weather/weather.component';

describe('AppComponent', () => {
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        AppComponent,
        DatetimeComponent,
        TodoComponent,
        MapComponent,
        NoteComponent,
        EventComponent,
        WeatherComponent,
        LoginComponent,
        RegisterComponent
      ],
      imports: [
        BrowserModule,
        FormsModule,
        HttpClientModule,
        LeafletModule,
        FullCalendarModule,
      ],
      providers: [TodoService, NoteService, EventService, LoginService, RegisterService],
    }).compileComponents();
  });

  // it('should create the app', () => {
  //   const fixture = TestBed.createComponent(AppComponent);
  //   const app = fixture.componentInstance;
  //   expect(app).toBeTruthy();
  // });

  // it(`should have as title 'ui'`, () => {
  //   const fixture = TestBed.createComponent(AppComponent);
  //   const app = fixture.componentInstance;
  //   expect(app.title).toEqual('ui');
  // });

  it('should render title', () => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('.content h2')?.textContent).toContain('Todos');
  });
});
