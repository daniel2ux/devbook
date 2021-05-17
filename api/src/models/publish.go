package models

import "time"

type Publish struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   string    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      string    `json:"likes"`
	CreateAt   time.Time `json:"createAt,omitempty"`
}
