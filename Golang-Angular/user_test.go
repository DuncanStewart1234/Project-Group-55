package test

import (
	"testing"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/user"
)

// func TestUserGet(t *testing.T) {
// 	list := user.Get()
// 	if list == nil {
// 		t.Errorf("error with get function")
// 	}
// }

func TestUserAdd(t *testing.T) {
	fname := "Endrick"
	lname := "Lafosse"
	uname := "elafosse"
	email := "elafosse@gmail.com"
	pass := "password"
    _, err := user.Add(fname, lname, uname, email, pass)
    if err != nil {
		t.Errorf("error with add function")
	}
}

// func TestUserDelete(t *testing.T) {
// 	err := user.Delete("419705")
// 	if err != nil {
// 		t.Errorf("error with delete function")
// 	}
// }
