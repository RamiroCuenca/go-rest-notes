package main

import "github.com/RamiroCuenca/go-rest-notes/common/logger"

func main() {
	// Init ZapLog in order to be able to call it all over the app
	logger.InitZapLog()
}
