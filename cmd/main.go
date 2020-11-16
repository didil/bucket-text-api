package main

import (
	"log"

	"github.com/didil/bucket-text-api/api"
)

func main() {
	err := api.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
