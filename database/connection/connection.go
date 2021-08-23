package connection

import (
	"database/sql"

	"github.com/RamiroCuenca/go-rest-notes/common/logger"
)

type PostgreClient struct {
	*sql.DB // database
}

// Creates a connection with a postgre database
// Should receive the source from it as a parameter
func NewPostgreClient(source string) *PostgreClient {
	db, err := sql.Open("postgre", source)

	if err != nil {
		// If we can not connect to the database, log the error and close the app with panic
		logger.ZapLog().Errorf("Error opening the database. Reason: %v", err)
		panic(err)
	}

	// Check if the connection with the database is stable and alive
	err = db.Ping()

	if err != nil {
		logger.ZapLog().Errorf("Error with the connection with the database. Reason: %v", err)
	}

	// As there are no errors, return the database connection
	return &PostgreClient{db}
}
