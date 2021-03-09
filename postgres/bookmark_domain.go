package main

import (
	"time"
)

// Bookmark domain entity
type Bookmark struct {
	ID        int
	Name      string
	URI       string
	Category  string
	CreatedAt time.Time
}
