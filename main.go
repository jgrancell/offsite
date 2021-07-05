package main

import (
	"fmt"
	"os"

	"github.com/jgrancell/offsite/app"
)

var (
	version string = "0.0.1"
)

func main() {
	os.Exit(Run(os.Args[1:]))
}

func Run(args []string) int {

	app := &app.App{}
	err := app.Load(args, version)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return 1
	}

	// TODO: add functionality to the application

	err = app.Configuration.Save()
	if err != nil {
		fmt.Println("Failed to save offsite configuration file at path", app.Configuration.ConfigPath)
		fmt.Println("Error:", err.Error())
		return 1

	}
	return 0
}
