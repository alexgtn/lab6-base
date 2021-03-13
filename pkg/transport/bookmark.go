package transport

import (
	"encoding/json"
	"github.com/alexgtn/esi2021-lab4/pkg/domain"
	"log"
	"net/http"
)

type bookmarkService interface {
	Create(book *domain.Bookmark) (*domain.Bookmark, error)
	GetAll() ([]*domain.Bookmark, error)
}

// BookmarkHandler is the API handler instances for bookmarks
type BookmarkHandler struct {
	bookmarkService bookmarkService
}

// NewBookmarkHandler constructor returns new BookmarkHandler instance
func NewBookmarkHandler(bS bookmarkService) *BookmarkHandler {
	return &BookmarkHandler{
		bookmarkService: bS,
	}
}

// CreateBookmark handles creation of bookmark
func (h *BookmarkHandler) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request %v", r)
	bookmark := &domain.Bookmark{}
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

	createdBookmark, err := h.bookmarkService.Create(bookmark)
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
	bookmarks, err := h.bookmarkService.GetAll()
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
