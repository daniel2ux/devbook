package models

import "time"

type Post struct {
	ID         int64     `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      uint64    `json:"likes,omitempty"`
	CreateAt   time.Time `json:"create_at,omitempty"`
}
