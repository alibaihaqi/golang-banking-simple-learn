package main

import (
	"github.com/alibaihaqi/banking/app"
	"github.com/alibaihaqi/banking/logger"
)

func main() {
	logger.Info("Starting our application")
	app.Start()
}
