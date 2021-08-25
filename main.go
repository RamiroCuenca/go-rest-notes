package main

import (
	"fmt"

	"github.com/RamiroCuenca/go-rest-notes/common/logger"
	"github.com/RamiroCuenca/go-rest-notes/database/connection"
)

func main() {
	// Init ZapLog in order to be able to call it all over the app
	logger.InitZapLog()

	// Init the connection with the database
	client := connection.NewPostgreClient()

	if client != nil {
		fmt.Print("Client connected succesfully")
	}
}
