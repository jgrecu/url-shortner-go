package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/jgrecu/url-shortner-go/internal/controllers"
	"github.com/jgrecu/url-shortner-go/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqliteDB, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer func(sqliteDB *sql.DB) {
		err := sqliteDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(sqliteDB)

	if err := db.CreateTable(sqliteDB); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("GET /{$}", controllers.ShowIndex)
	http.HandleFunc("GET /", controllers.Proxy(sqliteDB))
	http.HandleFunc("POST /", controllers.Shorten(sqliteDB))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
