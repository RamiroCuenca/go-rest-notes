package models

import (
	"errors"
	"time"
)

// Note structure wich wi save at the db
type Note struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Title     string    `json:"title"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"created_at"`
}

// The structure we should receive from the POST method as
// ID and Created at are added automatically.
type CreateNoteCMD struct {
	Owner   string `json:"owner"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

func (n CreateNoteCMD) Validate() error {

	// Check if owner, title or details are empty
	if n.Owner == "" {
		return errors.New("You cant create a note with an empty owner")
	}

	if n.Title == "" {
		return errors.New("You cant create a note with an empty title")
	}

	if n.Details == "" {
		return errors.New("You cant create a note with an empty detail")
	}

	return nil
}
