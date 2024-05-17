package main

import (
	"log"

	"github.com/ZainJavedDev/cleanup_optimize/database"
)

func main() {
	err := database.RemoveOldMatches()
	if err != nil {
		log.Fatal(err)
	}
	err = database.VacuumFullTables()
	if err != nil {
		log.Fatal(err)
	}
}
