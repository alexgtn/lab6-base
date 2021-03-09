package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"time"
)

// BookmarkRepository stores bookmarks
type BookmarkRepository struct {
	db *redis.Client
}

// NewBookmarkRepository builds new bookmark repository using a db connection
func NewBookmarkRepository(db *redis.Client) *BookmarkRepository {
	return &BookmarkRepository{
		db: db,
	}
}

// Create creates bookmark
func (r *BookmarkRepository) Create(bookmark *Bookmark) (*Bookmark, error) {
	bookmarkID := uuid.NewString()
	bookmark.ID = bookmarkID
	bookmark.CreatedAt = time.Now()
	_, err := r.db.HSetNX(context.Background(), "app:bookmark", bookmarkID, bookmark).Result()
	if err != nil {
		return nil, fmt.Errorf("create: redis error: %w", err)
	}

	return bookmark, nil
}

// GetAll gets all bookmarkss
func (r *BookmarkRepository) GetAll() ([]*Bookmark, error) {

	res, err := r.db.HGetAll(context.Background(), "app:bookmark").Result()
	if err != nil {
		return nil, fmt.Errorf("error getting bookmarks, err: %v", err)
	}
	bookmarks := []*Bookmark{}
	for _, stringValue := range res {
		b := &Bookmark{}
		err := json.Unmarshal([]byte(stringValue), b)
		if err != nil {
			return nil, fmt.Errorf("error decoding result, err: %v", err)
		}
		bookmarks = append(bookmarks, b)
	}

	return bookmarks, nil
}
