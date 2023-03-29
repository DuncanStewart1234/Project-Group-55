Work Completed in Spirit 3:


FrontEnd Unit Tests:

    import { ComponentFixture, TestBed } from '@angular/core/testing';
    import { TodoService } from '../todo.service';
    import { TodoComponent } from './todo.component';

    import {HttpClientModule} from '@angular/common/http';
    import { BrowserModule } from '@angular/platform-browser';
    import { FormsModule } from '@angular/forms';

    describe('TodoComponent', () => {
      let component: TodoComponent;
      let fixture: ComponentFixture<TodoComponent>;

      beforeEach(async () => {
        await TestBed.configureTestingModule({
          declarations: [ TodoComponent ],
          imports: [HttpClientModule, FormsModule, BrowserModule],
          providers: [TodoService]
        })
        .compileComponents();

        fixture = TestBed.createComponent(TodoComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
      });

      it('should create', () => {
        expect(component).toBeTruthy();
      });

      // it('should create', () => {
      //   expect(fixture.title).toEqual('app-todo');
      // });
    });
---
    
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
---
    
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
---
    
  Cypress:
  
    We tested the addition of the Todos button widget and if the textbox along with it worked
    
    Also testing the notes widgit as well


BackEnd Unit Tests


Updated Documentation for BackEnd API
