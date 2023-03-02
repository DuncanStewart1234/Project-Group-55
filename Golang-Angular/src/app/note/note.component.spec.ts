import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NoteComponent } from './note.component';
import { NoteService } from '../note.service';

import {HttpClientModule} from '@angular/common/http';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

describe('NoteComponent', () => {
  let component: NoteComponent;
  let fixture: ComponentFixture<NoteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NoteComponent ],
      imports: [HttpClientModule, FormsModule, BrowserModule],
      providers: [NoteService]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NoteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
