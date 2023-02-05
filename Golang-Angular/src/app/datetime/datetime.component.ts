import { Component, OnInit } from '@angular/core';
import {formatDate } from '@angular/common';

@Component({
  selector: 'app-datetime',
  templateUrl: './datetime.component.html',
  styleUrls: ['./datetime.component.css']
})
export class DatetimeComponent implements OnInit{
  today: Date;
  weekday: String;
  month: String;
  realtime = '';
  weekdayList: string[];
  monthList: string[];

  year: number;
  date: number;

  ngOnInit(): void {
    this.utcTime();
}

  constructor() {
    this.weekdayList = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]
    this.monthList = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]
    this.today= new Date();
    this.date = this.today.getDate();
    this.year = this.today.getFullYear();
    this.weekday = this.weekdayList[this.today.getDay()];
    this.month = this.monthList[this.today.getMonth()];

  }

  utcTime(): void {
    setInterval(() => {         //replaced function() by ()=>
      this.today= new Date();
      this.realtime = formatDate(this.today, 'hh:mm:ss a', 'en-US');
    }, 1000);
}
}
