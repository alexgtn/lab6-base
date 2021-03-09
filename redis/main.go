package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	httpServicePort = 8080
	redisURI        = "redis:6379"
	redisPassword   = "" // no password set
	redisDB         = 0  // use default DB
)

func main() {
	log.Println("Start bookmark server")

	// open Redis connection
	dbConn := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       redisDB,
	})

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

	err := srv.ListenAndServe()
	if err != nil {
		err = dbConn.Close()
		if err != nil {
			log.Fatalf("Could not disconnect from db")
		}
		log.Fatalf("Could not start server")
	}

	log.Println("Stop bookmark server")
}
