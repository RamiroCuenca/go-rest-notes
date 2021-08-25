package controllers

import (
	"time"

	"github.com/RamiroCuenca/go-rest-notes/common/logger"
	"github.com/RamiroCuenca/go-rest-notes/database/connection"
	"github.com/RamiroCuenca/go-rest-notes/notes/models"
)

type database struct {
	*connection.PostgreClient
}

// Here we are going to add al the CRUD operations as methods of the database client.

// Create is a method of the database client and it receives an CreateNoteCMD struc
// and it return the created note or an error
func (db *database) Create(cmd *models.CreateNoteCMD) (*models.Note, error) {

	// We should follow these steps:
	// 1- Start a transaction to the db.  ✅
	// 2- Execute the SQL query with the Note  ✅
	// 		2-B- If there is an error --> ROLLBACK  ✅
	// 3- Get Note id.  ✅
	// 4- Commit the transaction!  ✅
	// 5- Return the data to the client.  ✅

	// 1- Start transaction to the db
	tx, err := db.PostgreClient.Begin()

	if err != nil {
		logger.ZapLog().Errorf("Couldnt start transaction: %v", err)
	}

	// 2- Execute the SQL query
	result, err := tx.Exec(`INSERT INTO notes (owner, title, details) VALUES (?, ?, ?)`, cmd.Owner, cmd.Title, cmd.Details)

	// If there is an error, ROLLBACK!
	if err != nil {
		logger.ZapLog().Errorf("Couldnt execute SQL query: %v", err)
		_ = tx.Rollback()
		return nil, err
	}

	// 3- Get created Note id
	id, err := result.LastInsertId()

	if err != nil {
		logger.ZapLog().Errorf("Couldnt fetch Note id: %v", err)
		_ = tx.Rollback()
		return nil, err
	}

	// 4- Commit the transaction!
	_ = tx.Commit()

	// 5- Return the data to the client.
	return &models.Note{
		ID:        id,
		Owner:     cmd.Owner,
		Title:     cmd.Title,
		Details:   cmd.Details,
		CreatedAt: time.Now(),
	}, nil
}
