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
