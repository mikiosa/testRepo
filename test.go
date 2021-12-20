package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"log"
)

var DbConnection *sql.DB

func main() {
	db, _ :=sql.Open("sqlite3","/var/tmp/testDb")
	defer db.Close()

	rows, err := db.Query("SELECT col1, col2 FROM tb1")
	if err != nil {
		log.Fatal("aaaa:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var c1 int
		var c2 string
		if err := rows.Scan(&c1, &c2); err != nil {
			log.Fatal(err)
		}
		fmt.Println("%d %s", c1, c2)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
