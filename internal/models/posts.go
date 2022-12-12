package models

type Post struct {
	PostId int    `json:"id" db:"id"`
	UserId int    `json:"userId" db:"user_id"`
	Title  string `json:"title" db:"title"`
	Body   string `json:"body" db:"body"`
}
