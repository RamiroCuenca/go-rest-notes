package models

import "time"

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

func ValidateCreateNoteCMD() {

}
