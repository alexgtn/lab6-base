package main

import (
	"context"
	"database/sql"
	"fmt"
)

// BookmarkRepository stores bookmarks
type BookmarkRepository struct {
	db *sql.DB
}

// NewBookmarkRepository builds new bookmark repository using a db connection
func NewBookmarkRepository(db *sql.DB) *BookmarkRepository {
	return &BookmarkRepository{
		db: db,
	}
}

// Create creates bookmark
func (r *BookmarkRepository) Create(bookmark *Bookmark) (*Bookmark, error) {
	query := "INSERT into bookmark (category, name, uri) values($1, $2, $3) RETURNING id, created_at"
	// QueryRowContext executes the query
	// context usually holds an execution timer, metadata etc.
	// context.Background() creates an default empty context entity
	// row will contain id, created_at attributes based on the 'RETURNING'
	row := r.db.QueryRowContext(context.Background(), query, bookmark.Category, bookmark.Name, bookmark.URI)
	if row == nil {
		return nil, fmt.Errorf("error inserting bookmark %v", bookmark)
	}
	// Scan copies the columns from the row into the struct params
	// Scan takes arbitrary number of parameters
	err := row.Scan(&bookmark.ID, &bookmark.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error inserting bookmark %v", bookmark)
	}

	return bookmark, nil
}

// GetAll gets all bookmarks
func (r *BookmarkRepository) GetAll() ([]*Bookmark, error) {
	query := "SELECT id, category, name, uri, created_at from bookmark"
	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error querying bookmarks, err: %v", err)
	}

	bookmarks := []*Bookmark{}
	for rows.Next() {
		b := &Bookmark{}
		err := rows.Scan(&b.ID, &b.Category, &b.Name, &b.URI, &b.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scaning query, err: %v", err)
		}
		bookmarks = append(bookmarks, b)
	}
	// close rows to avoid memory leak
	err = rows.Close()
	if err != nil {
		return nil, fmt.Errorf("could not close rows, err %v", err)
	}

	return bookmarks, nil
}
