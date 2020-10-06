package main

import (
	"github.com/dmazzella--/GoBasha_users-api/app"
	"github.com/dmazzella--/GoBasha_users-api/logger"
)

func main() {
	logger.Info("About to start")
	app.StartApplication()
	logger.Info("exiting")
}
