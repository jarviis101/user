package main

import (
	"log"
	"user/internal"
)

func main() {
	app, err := internal.NewApp()
	if err != nil {
		log.Fatalln(err)
	}

	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
