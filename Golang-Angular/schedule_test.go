package test

import (
	"testing"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/schedule"
)

func TestPackageGet(t *testing.T) {
	list := schedule.Get()
	if list == nil {
		t.Errorf("error with get function")
	}
}

func TestPackageAdd(t *testing.T) {
	cid := "215679"
    _, err := schedule.Add(cid)
    if err != nil {
		t.Errorf("error with add function")
	}
}

func TestPackageDelete(t *testing.T) {
	err := schedule.Delete("cgdk46siuh41bo3t6100")
	if err != nil {
		t.Errorf("error with delete function")
	}
}
