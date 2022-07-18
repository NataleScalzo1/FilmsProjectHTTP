package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Film struct {
	ID      int       `json:"id"`
	Title   string    `json:"titlename"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Expires string    `json:"expires"`
}
