package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const DEFAULT_MYSQL_DRIVERNAME = "mysql"

func main() {
	mainCtx := context.TODO()

	dataSourceName := "admin:admin@/Test"
	db, err := sql.Open(DEFAULT_MYSQL_DRIVERNAME, dataSourceName)
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.PingContext(mainCtx); err != nil {
		log.Fatal(err)
	}

	rows, err := db.QueryContext(mainCtx, "SELECT * FROM Users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var lastname string
		var email string
		if err := rows.Scan(&id, &name, &lastname, &email); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %s %s %s\n", id, name, lastname, email)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
