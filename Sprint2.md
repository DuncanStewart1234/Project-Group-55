**Work We've Completed in Sprint 2:**

  Frontend:
  
    We've completed notes widget and are continuing to work on the schedule widget and will integrate it to the backend in the upcoming sprint.
    
    While also having addition unexpected additions with a dark mode to the overall design.
    
  Backend:

**Unit Tests and Cypress Tests in Frontend:**
  
  Unit Tests:
  
    If they're able to see the Todo widget
    
    Able to see the date and time
    
    Able to see the map
    
    Able to see the Note Widget
    
  Cypress:
  
    We tested the addition of the Todos button widget and if the textbox along with it worked

**Unit Test in Backend:**


**Documentation For Backend API:**

Package notes
import "github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/notes"

Overview:
The notes package implements and maintains a list of notes entered by the user

Index:
  func Add(title string, message string) string
  func Delete(id string) error
  func Edit(id string, new_msg string) error
  type Note
  func Get() []Note

Package files:
  notes.go

func Add
  func Add(title string, message string) string
    Add creates and adds a note to the notes list

func Delete
  func Delete(id string) error
    Delete removes and deletes a note from the notes list

func Edit
  func Edit(id string, new_msg string) error
    Edit finds a note in the list and edits its message

type Note
  Note is a struct that holds info needed for notes list
type Note struct {
    gorm.Model
    ID      string `json:"id"`
    User_ID string `json:"uid"`
    Title   string `json:"title"`
    Message string `json:"message"`
}

func Get
  func Get() []Note
    Get returns the notes list
