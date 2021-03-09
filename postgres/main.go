package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	// SQL driver
	// https://www.calhoun.io/why-we-import-sql-drivers-with-the-blank-identifier/
	// The sql package must be used in conjunction with a database driver. In this case PostgreSQL driver.
	// See https://golang.org/s/sqldrivers for a list of drivers.
	_ "github.com/lib/pq"
)

const (
	httpServicePort    = 8080
	postgresConnection = "dbname=postgres host=postgres password=postgres user=postgres sslmode=disable port=5432"
)

func main() {
	log.Println("Start bookmark server")

	// open Postgres connection
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		log.Fatal(err)
	}

	// construct application
	bookmarkRepository := NewBookmarkRepository(dbConn)
	bookmarkHandler := NewBookmarkHandler(bookmarkRepository)

	router := mux.NewRouter()
	// POST /bookmark
	router.HandleFunc("/bookmark", bookmarkHandler.CreateBookmark).Methods(http.MethodPost)
	// GET /bookmark
	router.HandleFunc("/bookmark", bookmarkHandler.GetBookmarks).Methods(http.MethodGet)

	// setup http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpServicePort),
		Handler: router,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server")
	}

	log.Println("Stop bookmark server")
}
