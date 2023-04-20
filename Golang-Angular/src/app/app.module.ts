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


import { LeafletModule } from '@asymmetrik/ngx-leaflet';
import { EventComponent } from './event/event.component';
import { EventService } from './event.service';
// import "leaflet/dist/leaflet.css";

import { FullCalendarModule } from '@fullcalendar/angular';

import { WeatherComponent } from './weather/weather.component';
import { AppRoutingModule } from './app-routing.module';
import { HomepageComponent } from './homepage/homepage.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';

import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';
import { LoginService } from './login.service';
import { RegisterService } from './register.service';

@NgModule({
  declarations: [
    AppComponent,
    DatetimeComponent,
    TodoComponent,
    MapComponent,
    NoteComponent,
    EventComponent,
    WeatherComponent,
    HomepageComponent,
    LoginComponent,
    RegisterComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpClientModule,
    LeafletModule,
    FullCalendarModule,
    AppRoutingModule,
    MatInputModule
    
  ],
  providers: [TodoService, NoteService, EventService, LoginService, RegisterService],
  bootstrap: [AppComponent]
})
export class AppModule { }