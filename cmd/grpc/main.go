package main

import (
	"log"
	"user/internal"
)

func main() {
	app, err := internal.NewGrpcApp()
	if err != nil {
		log.Fatalln(err)
	}

	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
