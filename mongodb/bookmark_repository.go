package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// BookmarkRepository stores bookmarks
type BookmarkRepository struct {
	db *mongo.Client
}

// NewBookmarkRepository builds new bookmark repository using a db connection
func NewBookmarkRepository(db *mongo.Client) *BookmarkRepository {
	return &BookmarkRepository{
		db: db,
	}
}

// Create creates bookmark
func (r *BookmarkRepository) Create(bookmark *Bookmark) (*Bookmark, error) {
	collection := r.db.Database("local").Collection("bookmark")

	bookmark.CreatedAt = time.Now()
	res, err := collection.InsertOne(context.Background(), bookmark)
	if err != nil {
		return nil, fmt.Errorf("error inserting bookmark %v, err: %v", bookmark, err)
	}
	bookmark.ID = res.InsertedID

	return bookmark, nil
}

// GetAll gets all bookmarkss
func (r *BookmarkRepository) GetAll() ([]*Bookmark, error) {
	collection := r.db.Database("local").Collection("bookmark")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting bookmarks, err: %v", err)
	}
	bookmarks := []*Bookmark{}
	for cursor.Next(context.Background()) {
		b := &Bookmark{}
		err := cursor.Decode(&b)
		if err != nil {
			return nil, fmt.Errorf("error decoding result, err: %v", err)
		}
		bookmarks = append(bookmarks, b)
	}
	err = cursor.Close(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting bookmarks, err: %v", err)
	}

	return bookmarks, nil
}
