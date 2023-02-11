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
// import "leaflet/dist/leaflet.css";

@NgModule({
  declarations: [
    AppComponent,
    DatetimeComponent,
    TodoComponent,
    MapComponent,
    NoteComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpClientModule,
    LeafletModule,
  ],
  providers: [TodoService, NoteService],
  bootstrap: [AppComponent]
})
export class AppModule { }