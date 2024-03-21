package main

import (
	"crud-apis-db-app/initiate"
	"log"
)

func main() {
	err := initiate.Initiate()
	if err != nil {
		log.Fatalf("app crashed!")
	}
}
