package media

import (
	"io"

	"github.com/knadh/listmonk/models"
	"gopkg.in/volatiletech/null.v6"
)

// Media represents an uploaded object.
type Media struct {
	Meta        models.JSON `db:"meta" json:"meta"`
	CreatedAt   null.Time   `db:"created_at" json:"created_at"`
	UUID        string      `db:"uuid" json:"uuid"`
	Filename    string      `db:"filename" json:"filename"`
	ContentType string      `db:"content_type" json:"content_type"`
	Thumb       string      `db:"thumb" json:"-"`
	Provider    string      `json:"provider"`
	URL         string      `json:"url"`

	ThumbURL null.String `json:"thumb_url"`
	ID       int         `db:"id" json:"id"`

	Total int `db:"total" json:"-"`
}

// Store represents functions to store and retrieve media (files).
type Store interface {
	Put(string, string, io.ReadSeeker) (string, error)
	Delete(string) error
	GetURL(string) string
	GetBlob(string) ([]byte, error)
}
