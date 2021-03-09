package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

const (
	httpServicePort = 8080
	mongoConnection = "mongodb://mongo:27017"
)

func main() {
	log.Println("Start bookmark server")

	// open Mongo connection
	dbConn, err := mongo.NewClient(options.Client().ApplyURI(mongoConnection))
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Connect(context.Background())
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
		err = dbConn.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("Could not disconnect from db")
		}
		log.Fatalf("Could not start server")
	}

	log.Println("Stop bookmark server")
}
