package main

import (
	"database/sql"
	"fmt"
	"github.com/alexgtn/esi2021-lab4/pkg/repository"
	"github.com/alexgtn/esi2021-lab4/pkg/service"
	"github.com/alexgtn/esi2021-lab4/pkg/transport"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"

	// SQL driver
	// https://www.calhoun.io/why-we-import-sql-drivers-with-the-blank-identifier/
	// The sql package must be used in conjunction with a database driver. In this case PostgreSQL driver.
	// See https://golang.org/s/sqldrivers for a list of drivers.
	_ "github.com/lib/pq"
)

var (
	httpServicePort    = os.Getenv("SERVICE_PORT")
	postgresConnection = os.Getenv("POSTGRES_CONNECTION")
)

func main() {
	log.Println("Start bookmark server")
	log.Printf("envs %s %s", httpServicePort, postgresConnection)

	// open Postgres connection
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		log.Fatal(err)
	}

	// construct application
	bookmarkRepository := repository.NewBookmarkRepository(dbConn)
	bookmarkService := service.NewBookmarkService(bookmarkRepository)
	bookmarkHandler := transport.NewBookmarkHandler(bookmarkService)

	router := mux.NewRouter()
	// POST /bookmark
	router.HandleFunc("/bookmark", bookmarkHandler.CreateBookmark).Methods(http.MethodPost)
	// GET /bookmark
	router.HandleFunc("/bookmark", bookmarkHandler.GetBookmarks).Methods(http.MethodGet)

	portInt, err := strconv.Atoi(httpServicePort)
	if err != nil {
		fmt.Errorf("error parsing port: %s", err.Error())
	}
	// setup http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", portInt),
		Handler: router,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server %s", err.Error())
	}

	log.Println("Stop bookmark server")
}
