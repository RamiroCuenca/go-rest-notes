package connection

import (
	"database/sql"
	"fmt"

	"github.com/RamiroCuenca/go-rest-notes/common/logger"
)

type PostgreClient struct {
	*sql.DB // database
}

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "secret"
	dbname   = "notes_db"
)

// Creates a connection with a postgre database
// Should receive the source from it as a parameter
func NewPostgreClient() *PostgreClient {

	// Log db credentials
	logger.ZapLog().Infof("psql info: \nhost=%s\nport=%d\nuser=%s\npassword=%s\ndbname=%s\nsslmode=disable", host, port, user, password, dbname)

	// db source
	source := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open method does not create the connection, it simply check if the arguments work
	// That's why we MUST check with Ping() method if it's working!
	db, err := sql.Open("postgres", source)
	// Close DB after program exits
	defer db.Close()

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
