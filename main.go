package main

import (
	database "main/dbconnection"
)

func main() {
	db := database.Connection()
	defer db.Close()
}
