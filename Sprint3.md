Work Completed in Spirit 3:

    We've completed the schedule for the backend and frontend which was uncompleted in the previous sprint with some little issues to be resolved with the integreation between the backend and front end and being able to finish the front and backend of the weather widget while still working on the news widget. While also doing the unit tests for these new api's as they need to be finished as well, once the actual api's are finished. 

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

    import { async, ComponentFixture, TestBed } from '@angular/core/testing';

    import { WeatherComponent } from './weather.component';

    describe('WeatherComponent', () => {
      let component: WeatherComponent;
      let fixture: ComponentFixture<WeatherComponent>;

      beforeEach(async(() => {
        TestBed.configureTestingModule({
          declarations: [ WeatherComponent ]
        })
        .compileComponents();
      }));

      beforeEach(() => {
        fixture = TestBed.createComponent(WeatherComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
      });

      it('should create', () => {
        expect(component).toBeTruthy();
      });
    });
    
---

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

    });
    
---

    import { ComponentFixture, TestBed } from '@angular/core/testing';

    import { EventComponent } from './event.component';
    import { MapComponent } from '../map/map.component';
    import {FullCalendarModule} from '@fullcalendar/angular';
    import {HttpClientModule} from '@angular/common/http';
    import { BrowserModule } from '@angular/platform-browser';
    import { FormsModule } from '@angular/forms';
    import { LeafletModule } from '@asymmetrik/ngx-leaflet';

    describe('EventComponent', () => {
      let component: EventComponent;
      let fixture: ComponentFixture<EventComponent>;

      beforeEach(async () => {
        await TestBed.configureTestingModule({
          declarations: [ EventComponent, MapComponent ],
          imports: [HttpClientModule, FormsModule, BrowserModule, FullCalendarModule, LeafletModule],
        })
        .compileComponents();

        fixture = TestBed.createComponent(EventComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
      });

      it('should create', () => {
        expect(component).toBeTruthy();
      });
    });

---

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
            WeatherComponent
          ],
          imports: [
            BrowserModule,
            FormsModule,
            HttpClientModule,
            LeafletModule,
            FullCalendarModule,
          ],
          providers: [TodoService, NoteService, EventService],
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

---

  Cypress:
  
    We tested the addition of the Todos button widget and if the textbox along with it worked
    
    Also testing the notes widgit as well


# Unit Test in Backend:

## Updated Documentation For Backend API:

### Package **notes**

*import "github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/notes"*

Overview:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The notes package implements and maintains a list of notes entered by the user

Index:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Add(title string, message string) string`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Delete(id string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Edit(id string, new_msg string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`type Note`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Get() []Note`  

Package files:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;notes.go  

func Add  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Add(title string, message string) string`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Add creates and adds a note to the notes list  

func Delete  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Delete(id string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Delete removes and deletes a note from the notes list  

func Edit  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Edit(id string, new_msg string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Edit finds a note in the list and edits its message  

type Note  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Note is a struct that holds info needed for notes list  

type Note struct {  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;gorm.Model  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ID      string `json:"id"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;User_ID string `json:"uid"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Title   string `json:"title"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Message string `json:"message"`  
}  
func Get  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Get() []Note`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Get returns the notes list  
  
### Package **todo**

*import "github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/todo"*

Overview:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;This package implements and maintains a todo list for the user to work with.  

Index:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Add(message string) (string, error)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Complete(id string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Delete(id string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`type Todo`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Get() []Todo`  
Package files:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;todo.go  

func Add  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Add(message string) (string, error)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Add a new item to the Todo list based off message input  

func Complete  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Complete(id string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Complete marks a Todo list item as complete  

func Delete  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Delete(id string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Delete a Todo list item  

type Todo  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Todo is a struct model containing the Todo list item info  

type Todo struct {  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;gorm.Model  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ID       string `json:"id"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;User_ID  string `json:"uid"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Message  string `json:"message" gorm:"size:256"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Complete bool   `json:"complete"`  
}  
func Get  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Get() []Todo`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Get returns the Todo list  

### Package **handlers**

*import "github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/handlers"*

Overview:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;This package implements and maintains handler functions for our application  

Index:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func AddNotesHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func AddTodoHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func CompleteTodoHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func DeleteNotesHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func DeleteTodoHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func EditNotesHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func GetNotesHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func GetTodoListHandler(c *gin.Context)`  
Package files:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;handlers.go  

func AddNotesHandler  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func AddNotesHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;AddNotesHandler is used to add a note to the notes list from the Notes package  

func AddTodoHandler  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func AddTodoHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;AddTodoHandler is used to add to the todo list from the Todo package  

func CompleteTodoHandler  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func CompleteTodoHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;CompleteTodoHandler is used to mark a todo list item as complete  

func DeleteNotesHandler  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func DeleteNotesHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;DeleteNotesHandler is used to delete a note from the notes list  

func DeleteTodoHandler  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func DeleteTodoHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;DeleteTodoHandler is used to delete an item from the todo list  

func EditNotesHandler  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func EditNotesHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;EditNotesHandler is used to edit a note from the notes list  

func GetNotesHandler  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func GetNotesHandler(c *gin.Context)`  
Notes Handlers GetNotesHandler is used to request the notes list from the Notes package  

func GetTodoListHandler  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func GetTodoListHandler(c *gin.Context)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;To Do Handlers GetTodoListHandler is a handler to request the todo list from the Todo package  

### Package **utils**

import "github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"

Overview:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;This package contains a utility function for the sqlite databases  

Index:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func GetDB(path string) *gorm.DB`  
Package files:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;utils.go  

func GetDB  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func GetDB(path string) *gorm.DB`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;GetDB returns the specified database for use in other packages  

### Package **weather**

import "github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/weather"

Overview:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;This is a package using the openweathermap API to get local weather conditions  

Index:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Constants  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func GetForecast(where string, unit string, lang string) string`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func GetWeather(where string, unit string, lang string) string`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;type Data  
Package files:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;weather.go  

Constants  
URL is a constant that contains where to find the user IP info  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`const URL = "http://ip-api.com/json"`  
func GetForecast  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func GetForecast(where string, unit string, lang string) string`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;GetForecast gets the five day weather forecast at the location provided with a specified language and specified units.  

func GetWeather  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func GetWeather(where string, unit string, lang string) string`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;GetWeather gets the current weather at the location provided with a specified language and specified units.  

type Data  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Data will hold the result of the query to get the IP address of the caller.  

type Data struct {
    Status      string  `json:"status"`
    Country     string  `json:"country"`
    CountryCode string  `json:"countryCode"`
    Region      string  `json:"region"`
    RegionName  string  `json:"regionName"`
    City        string  `json:"city"`
    Zip         string  `json:"zip"`
    Lat         float64 `json:"lat"`
    Lon         float64 `json:"lon"`
    Timezone    string  `json:"timezone"`
    ISP         string  `json:"isp"`
    ORG         string  `json:"org"`
    AS          string  `json:"as"`
    Message     string  `json:"message"`
    Query       string  `json:"query"`
}


# Backend Unit Tests

    func TestDB(t *testing.T) {
        id, err := todo.Add("Test Todo Task")
        if err != nil {
            t.Errorf("Add error")
        }

        list := todo.Get()

        for _, task := range list {
            if id == task.ID {
                if task.Message != "Test Todo Task" {
                    t.Errorf("Did not add task")
                }
                break
            }
        }

        err = todo.Delete(id)
        if err != nil {
            t.Errorf("Failed to delete task")
        }
    }

---

    func TestEmptyMessageAdd(t *testing.T) {
        // Don't Allow Empty Task Message
        _, err := todo.Add("")
        if err == nil {
            t.Errorf("Added Task with empty message")
        }
    }

---

    func TestLargeMessageAdd(t *testing.T) {
        // Don't Allow Empty Task Message
        _, err := todo.Add("Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?")
        if err == nil {
            t.Errorf("Added Task with large message")
        }
    }

---

    func TestDeleteImaginaryTask(t *testing.T) {
        // Attempt To Delete Task That Does Not Exist
        err := todo.Delete("")
        if err == nil {
            t.Errorf("Failed to capture error")
        }
    }

---

    func TestAddTodoItem(t *testing.T) {

        for _, test := range addTestTodo {
            Add(test.arg)
            want := false

            testList = Get()

            for i, t := range testList {
                if t.Message == test.arg {
                    want = true
                    _ = i
                    continue
                }
            }

            if want == false {
                t.Errorf("note was not added")
            }
        }
    }

---

    func TestAddNote(t *testing.T) {

        for _, test := range addTestNotes {
            Add("Title", test.arg)
            want := false

            testList = Get()

            for i, t := range testList {
                if t.Message == test.arg {
                    want = true
                    _ = i
                    continue
                }
            }

            if want == false {
                t.Errorf("note was not added")
            }
        }
    }
