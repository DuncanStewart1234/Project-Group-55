import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';


@Injectable()
export class EventService {
  constructor(private httpClient: HttpClient) {}

  getEventList() {
    return this.httpClient.get(environment.gateway + '/event');
  }

  addEvent(event: Event) {
    return this.httpClient.post(environment.gateway + '/event', event);
  }


  deleteEvent(event: Event) {
    return this.httpClient.delete(environment.gateway + '/event/' + event.title);
  }
}

export class Event {
  title: string;
  start: string;
  end: string;
  backgroundColor: string;
  lat: number;
  lng: number;
  all_day: boolean;
  reccur: any[];
}