import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DatetimeComponent } from './datetime.component';

import {HttpClientModule} from '@angular/common/http';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';


describe('DatetimeComponent', () => {
  let component: DatetimeComponent;
  let fixture: ComponentFixture<DatetimeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DatetimeComponent ],
      imports: [HttpClientModule, FormsModule, BrowserModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DatetimeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
