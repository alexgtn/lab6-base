package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// BookmarkHandler is the API handler instances for bookmarks
type BookmarkHandler struct {
	bookmarkRepository *BookmarkRepository
}

// NewBookmarkHandler constructor returns new BookmarkHandler instance
func NewBookmarkHandler(bookmarkRepository *BookmarkRepository) *BookmarkHandler {
	return &BookmarkHandler{
		bookmarkRepository: bookmarkRepository,
	}
}

// CreateBookmark handles creation of bookmark
func (h *BookmarkHandler) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request %v", r)
	bookmark := &Bookmark{}
	err := json.NewDecoder(r.Body).Decode(bookmark)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// close body to avoid memory leak
	err = r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createdBookmark, err := h.bookmarkRepository.Create(bookmark)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&createdBookmark)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// getBookmarks retrieves all bookmarks
func (h *BookmarkHandler) GetBookmarks(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request %v", r)
	bookmarks, err := h.bookmarkRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&bookmarks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
