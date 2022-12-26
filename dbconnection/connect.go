package dbconnection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connection() *sql.DB {

	connStr := "postgres://postgres:postgres@localhost/demodb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	CheckError(err)

	return db
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Databases Connected Successfully")
	}
}
