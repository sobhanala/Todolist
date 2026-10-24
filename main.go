package main

import (
	"log"

	"github.com/sobhan/tod/router"
)

func main() {
	if err := router.SetupRouter(); err != nil {
		log.Fatal(err)
	}
}
