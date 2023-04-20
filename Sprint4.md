Work Completed in Sprint 4:

    Our goal in this Sprint was mistly to work on the login aspects and implementing that throught the backend and frontend. While also refining other API's to either make them look nicer or just more user friendly. There is still a few kinks that we would like to work out that would be better to make better overtime outside the scope of this class and with less time constrainsts.

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

	import { ComponentFixture, TestBed } from '@angular/core/testing';

	import { LoginComponent } from './login.component';

	describe('LoginComponent', () => {
	  let component: LoginComponent;
	  let fixture: ComponentFixture<LoginComponent>;

	  beforeEach(async () => {
	    await TestBed.configureTestingModule({
	      declarations: [ LoginComponent ]
	    })
	    .compileComponents();

	    fixture = TestBed.createComponent(LoginComponent);
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


## Updated Documentation For Backend API:

### Package **course**

*import "github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/course"*

Overview:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;A package used to create a schedule of classes that a college student may use.

Index:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Add(name string, abbrv string, loc string, scheduleBlock string) (int, error)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func AddCal(title string, abbrv string, loc string, start string, end string) (int, error)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Close()`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Delete(cid string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Edit(id int, name string, abbrv string, loc string, start string, end string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Start()`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;type Class  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Get() []Class`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;type Location  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;type Period  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;type Schedule  

Package files  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;class.go  

func Add  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Add(name string, abbrv string, loc string, scheduleBlock string) (int, error)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Add creates and adds a class element to the list  

func AddCal  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func AddCal(title string, abbrv string, loc string, start string, end string) (int, error)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;AddCal creates and adds a class to the calender list. It is similar in functionality to Add, except that a start and end time are now inputted too.  

func Close  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Close()`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Close simply closes the SDQ database being used.  

func Delete  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Delete(cid string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Delete removes and deletes a class element from the list  

func Edit  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Edit(id int, name string, abbrv string, loc string, start string, end string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Edit finds a class item by its location in the list and then edits it. All inputs required to add a calender item are required for this function.  

func Start  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Start()`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Start is a constructor that calls initialiseList  

type Class  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Class is a struct used to contain info about a student's class  

type Class struct {  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Class_ID int      `json:"cid" gorm:"primaryKey"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Name     string   `json:"name"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Abbrv    string   `json:"abbrv"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Location Location `json:"loc" gorm:"serializer:json"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Schedule Schedule `json:"schedule" gorm:"serializer:json"`  
}  

func Get  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Get() []Class`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Get returns the list of classes in the schedule  

type Location  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;type Location struct {  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`Lat  float64`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`Long float64`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}  

type Period  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;type Period struct {  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`S_Time string`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`E_Time string`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}  

type Schedule  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;type Schedule struct {  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Mon  []Period `json:"Mon"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tues []Period `json:"Tues"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Wed  []Period `json:"Wed"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Thur []Period `json:"Thur"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Fri  []Period `json:"Fri"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Sat  []Period `json:"Sat"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Sun  []Period `json:"Sun"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}  

### Package **class**

*import "github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/class"*

Overview:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;A package used to create a schedule of classes that a college student may use.

Index:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Add(loc string) (string, error)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Delete(cid string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;type Class  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Get() []Class`  

Package files:  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;class.go  

func Add  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Add(loc string) (string, error)`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Add creates and adds a class element to the list  

func Delete  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Delete(cid string) error`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Delete removes and deletes a class element from the list  

type Class  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Class is a struct used to contain info about a student's class  

type Class struct {  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;gorm.Model  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// ID       string `json:"id"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Class_ID string `json:"cid"`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Location string `json:"loc"`  
}  

func Get  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`func Get() []Class`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Get returns the list of classes in the schedule  

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
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Status      string  `json:"status"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Country     string  `json:"country"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;CountryCode string  `json:"countryCode"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Region      string  `json:"region"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;RegionName  string  `json:"regionName"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;City        string  `json:"city"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Zip         string  `json:"zip"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Lat         float64 `json:"lat"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Lon         float64 `json:"lon"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Timezone    string  `json:"timezone"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ISP         string  `json:"isp"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ORG         string  `json:"org"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;AS          string  `json:"as"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Message     string  `json:"message"`  
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Query       string  `json:"query"`  
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

---

    func TestGetWeather(t *testing.T) {
    
	    got := weather.GetWeather("Gainesville", "f", "en")

	    if got == `Current weather for :
	        Conditions:
	        Now:         0 imperial
	        High:        0 imperial
	        Low:         0 imperial` {
		        t.Errorf("Current weather not returned, fatal error occured.")
	        }
    }
    
---

    func TestGetForecast(t *testing.T) {
	    got := weather.GetForecast("Gainesville", "f", "en")

	    if got == "Weather Forecast for :" {
		    t.Errorf("Forecast not returned, fatal error occured.")
	    }
    }

---

    func TestGetWeatherMultiWordCity(t *testing.T) {
	    got := weather.GetWeather("Los Angeles", "f", "en")

	    if got == `Current weather for :
	        Conditions:
	        Now:         0 imperial
	        High:        0 imperial
	        Low:         0 imperial` {
		        t.Errorf("Weather for city not returned, fatal error occured.")
	        }
    }
    
---

    func TestGetForecastCityWithSpecialCharacters(t *testing.T) {
	    got := weather.GetForecast("Cancun", "f", "en")

	    if got == "Weather Forecast for :" {
		    t.Errorf("Forecast for city not returned, fatal error occured.")
	    }
    }
    
---

    func TestGetWeatherBadInput(t *testing.T) {
	    got := weather.GetWeather("Ggssf", "f", "en")

	    if got == "" {
		    t.Errorf("False input not handled, fatal error occured.")
	    }
    }

---

    func TestGetForecastBadInput(t *testing.T) {
	    got := weather.GetForecast("xsafsdf", "f", "en")

	    if got == "" {
		    t.Errorf("False input not handled, fatal error occured.")
	    }
    }
