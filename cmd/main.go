package main

import (
	"fmt"
	"os"

	"github.com/coderj001/phoneguardian/app"
	"github.com/coderj001/phoneguardian/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	port := os.Getenv("SERVER_PORT")
	app.Run(fmt.Sprintf(":%s", port))
}
