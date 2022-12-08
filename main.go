package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Database credentials
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "demodb"
)

// To check if there is any error
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func databseConnection() {
	dbConnection := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbConnection)
	CheckError(err)
	defer db.Close() // Db connection will be closed after the main function completes its execution

	addData(db) // Inserting data into
	fmt.Println("Connection Established Sucessfully!")
}

func addData(db *sql.DB) {
	insertData := `INSERT INTO "vehicle".motor ("vehicle_name","vehicle_company") VALUES ('Pridor','Honda')`
	_, e := db.Exec(insertData)
	CheckError(e)
}

func main() {
	databseConnection()
}
