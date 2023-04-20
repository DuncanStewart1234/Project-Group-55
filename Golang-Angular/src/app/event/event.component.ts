import { Component, OnInit, ViewChild, ChangeDetectorRef } from '@angular/core';
import { CalendarOptions, EventInput } from '@fullcalendar/core';
import { MapComponent } from '../map/map.component';
import interactionPlugin from '@fullcalendar/interaction';
import dayGridPlugin from '@fullcalendar/daygrid';
import listPlugin from '@fullcalendar/list';
import timegridplugin from '@fullcalendar/timegrid';
import iCalendarPlugin from '@fullcalendar/icalendar'

@Component({
  selector: 'app-event',
  templateUrl: './event.component.html',
  styleUrls: ['./event.component.css'],
})


export class EventComponent implements OnInit{
  events: any[] = [];
  // calendarOptions!: CalendarOptions;

  title: string;
  year: string;
  month: string;
  day: string;
  hour: string;
  minute: string;
  total: string;
  trash: boolean;

  date: string;
  start_time: string;
  end_time: string;
  all_day: boolean;
  daily: boolean;
  weekly: boolean;
  color: string;
  days: [];

  url: string;
  showMap: boolean = true;

  calendarOptions: CalendarOptions = {
    height: 500,
    headerToolbar: {
      left: 'today',
      center: 'prev,next',
      right: 'listWeek,dayGridMonth'
    },
    // headerToolbar: false,
    plugins: [dayGridPlugin, interactionPlugin, listPlugin, timegridplugin, iCalendarPlugin],
    initialView: 'listWeek',
    // listWeek, dayGridWeek, timeGridDay
    weekends: true,
    dateClick: this.onDateClick.bind(this),
    eventClick:  this.onEventClick.bind(this),
    // events: this.events,
    editable: true,
    selectable: true
  };
  // eventsPromise: Promise<EventInput>;

  @ViewChild("mapComponent") mapComponent: MapComponent;
  
  constructor(private cd: ChangeDetectorRef) { 
    this.events = [
      {
        title: 'Meeting',
        start: '2023-04-19T14:30:00',
        end: '2023-04-19T15:30:00',
        backgroundColor: 'green',
        extendedProps: {
          lat: 29.643946,
          lng: -82.355659
        }
      },
      {
        title: 'Birthday Party',
        start: '2023-04-20T07:00:00',
        extendedProps: {
          lat: 29.643946,
          lng: -82.355659
        }
      },
      // {
      //   url: "https://ufl.instructure.com/feeds/calendars/user_7bCfJjkzu8jEcbw1Pe0i2YyEbnqXfFPmH9Hq9PK3.ics",
      //   format: "ics"
      // }
    ]
    this.all_day = false;
    this.daily = false;
    this.weekly = false;
    this.color = "red";
    // this.showMap = false;
  }


  ngOnInit() {

    setTimeout(() => {
      this.calendarOptions.events = this.events;
      this.showMap = false;
      // this.cd.detectChanges();
    }, 500);

    setInterval(() => {         //replaced function() by ()=>
      // this.cd.detectChanges();
      this.calendarOptions.events = this.events;
      // this.cd.detectChanges();
    }, 1000);
  }

  onEventClick(arg){
    if (this.trash){
      var title_ = arg.event.title;

      this.events = this.events.filter(item => item.title !== title_);
      // console.log(this.events);
      this.calendarOptions.events = [];
      this.cd.detectChanges();
      this.calendarOptions.events = this.events;
    }
    else{
      this.showMap = !this.showMap;
      this.mapComponent.createMarker(arg.event.title, arg.event.extendedProps.lat, arg.event.extendedProps.lng);
    }
  
  }


  onDateClick(res: { dateStr: string; }){
    alert('You clicked on : ' + res.dateStr);
    // this.calendarOptions.height = 400;
  }

  addEvent(){
    this.events.push({
      title: this.title,
      start: this.year + '-' + this.month + '-' + this.day + 'T' + this.hour + ':' + this.minute + ':00',
      // backgroundColor: 'green',
      // borderColor: 'green',
      extendedProps: {
        lat: 29.643946,
        lng: -82.355659
    }});

  
  
    // this.cd.detectChanges();
    this.calendarOptions.events = [];
    this.cd.detectChanges();
    this.calendarOptions.events = this.events;
  }

  CalView(){
    this.calendarOptions.initialView = "";
    this.cd.detectChanges();
    this.calendarOptions.initialView = "dayGridMonth";
    // $('#fullcalendar'). 
    
  }

  createEvent(){

    if (this.all_day || this.daily || this.weekly)
    {
      if (this.all_day){
        if (this.daily || this.weekly){
          if (this.daily){
            console.log("Daily, All-Day")
            this.events.push({
              title: this.title,
              allDAy: true,
              start: this.date,
              end: this.date,
              daysOfWeek: [1,2,3,4,5,6,7],
              backgroundColor: this.color,
              // borderColor: 'green',
              extendedProps: {
                lat: this.mapComponent.lat,
                lng: this.mapComponent.lng
            }});
          }
          else{

          }
        }
        else {
          console.log("All day, not daily or weekly")
          this.events.push({
            title: this.title,
            allDAy: true,
            start: this.date,
            end: this.date,
            editable: true,
            startEditable: true,
            durationEditable: true,
            resourceEditable: true,
            backgroundColor: this.color,
            // borderColor: 'green',
            extendedProps: {
              lat: this.mapComponent.lat,
              lng: this.mapComponent.lng
          }});
        }
      }
      // Not ALL DAY
      else{
        if (this.daily || this.weekly){
          if (this.daily){
            console.log("Daily, not all-day")
            this.events.push({
              title: this.title,
              allDay: false,
              start: this.date + 'T' + this.start_time,
              end: this.date + 'T' + this.end_time,
              daysOfWeek: [1,2,3,4,5,6,7],
              backgroundColor: this.color,
              // borderColor: 'green',
              extendedProps: {
                lat: this.mapComponent.lat,
                lng: this.mapComponent.lng
            }});
          }
          else{

          }
        }
        else {
          console.log("not daily, not weekly, not all-day")
          this.events.push({
            title: this.title,
            start: this.date + 'T' + this.start_time,
            end: this.date + 'T' + this.end_time,
            backgroundColor: this.color,
            // borderColor: 'green',
            extendedProps: {
              lat: this.mapComponent.lat,
              lng: this.mapComponent.lng
          }});
        }

      }
    }
    else{
      this.events.push({
        title: this.title,
        start: this.date + 'T' + this.start_time,
        end: this.date + 'T' + this.end_time,
        backgroundColor: this.color,
        // borderColor: 'green',
        extendedProps: {
          lat: this.mapComponent.lat,
          lng: this.mapComponent.lng
      }});
    }
  
    // this.cd.detectChanges();
    this.calendarOptions.events = [];
    this.cd.detectChanges();
    this.calendarOptions.events = this.events;
  }

  addiCal(){
    this.events.push({
      url: "https://ufl.instructure.com/feeds/calendars/user_7bCfJjkzu8jEcbw1Pe0i2YyEbnqXfFPmH9Hq9PK3.ics",
      format: 'ics'
    });
    // this.cd.detectChanges();
    this.calendarOptions.events = [];
    this.cd.detectChanges();
    this.calendarOptions.events = this.events;
  }

  toggleMap()
  {
    this.showMap = !this.showMap;
    // this.mapComponent.markers = [];
    // this.cd.detectChanges();
  }
  
}
