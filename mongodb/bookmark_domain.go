package main

import (
	"time"
)

// Bookmark domain entity
type Bookmark struct {
	ID        interface{} `bson:"_id,omitempty" json:"ID"`
	Name      string
	URI       string
	Category  string
	CreatedAt time.Time
}
