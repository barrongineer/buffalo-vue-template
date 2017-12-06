package main

import (
	"log"

	"github.com/barrongineer/plugin_test/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
