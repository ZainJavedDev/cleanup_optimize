package main

import "github.com/ZainJavedDev/cleanup_optimize/database"

func main() {
	database.RemoveOldMatches()
	database.OptimizeTables()
}
