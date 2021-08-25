package models

import (
	"testing"
)

// Receives a CreateNoteCMD and returns a Note
// It add automatically an id and a CreatedAt parameters
func newNote(owner string, title string, details string) *CreateNoteCMD {

	resp := &CreateNoteCMD{
		Owner:   owner,
		Title:   title,
		Details: details,
	}

	return resp
}

func Test_CreateNote_CorrectParams(t *testing.T) {

	n := newNote("Ramiro", "Note title for test", "Note details. Need to check if it is working")

	err := n.Validate()

	if err != nil {
		t.Error("There was an error creating the note inspite of sending correct parameters")
		t.Fail()
	} else {
		t.Log("Success")
	}

}

func Test_CreateNote_WrongParams(t *testing.T) {

	n := newNote("", "", "")

	err := n.Validate()

	if err == nil {
		t.Error("Inspite of sending wrong parameters the note was created.")
		t.Fail()
	} else {
		t.Log("Success")
	}

}
