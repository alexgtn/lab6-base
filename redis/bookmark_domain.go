package main

import (
	"encoding/json"
	"time"
)

// Bookmark domain entity
type Bookmark struct {
	ID        interface{}
	Name      string
	URI       string
	Category  string
	CreatedAt time.Time
}

func (t *Bookmark) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Bookmark) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	return nil
}
