package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RamiroCuenca/go-rest-notes/common/handler"
	"github.com/RamiroCuenca/go-rest-notes/common/logger"
	"github.com/RamiroCuenca/go-rest-notes/notes/controllers"
)

// // This handler manages the creation of new notes
// func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {

// 	// 1- We need to decode the json object to a CreateNoteCMD model
// 	// in order to send it to

// 	var db controllers.database
// }

// type database struct {
// 	*controllers.PostgreClient
// }

type CrudMethods struct {
	*controllers.Database
}

func (crud *CrudMethods) ReadAllNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Call the controller. Returns an array of notes.
	notesArr, err := crud.Database.ReadAll()

	if err != nil {
		logger.ZapLog().Errorf("ReadAll handler not working: %v", err)
		handler.SendError(w, http.StatusBadRequest) // REMEMBER TO SEND AN ERROR
		return
	}

	// Encode to json the array
	data, err := json.Marshal(notesArr)

	if err != nil {
		logger.ZapLog().Errorf("Could not decode the data: %v", err)
		handler.SendError(w, http.StatusBadRequest) // REMEMBER TO SEND AN ERROR
		return
	}

	// Use the common handler to send a response
	handler.SendResponse(w, http.StatusOK, data)

}
