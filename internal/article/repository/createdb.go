package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres dbname=postgres password=yura11011 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	schema := `CREATE TABLE IF NOT EXISTS place (
		country text,
		city text NULL,
		telcode integer);`

	// execute a query on the server
	_, err = db.Exec(schema)
	schema = `DROP TABLE place`
	_, err = db.Exec(schema)

}
