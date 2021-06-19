package models

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	CreateAt  time.Time `json:"createAt"`
	Followers []User    `json:"followers"`
	Follow    []User    `json:"follow"`
	Posts     []Post    `json:"posts"`
}
