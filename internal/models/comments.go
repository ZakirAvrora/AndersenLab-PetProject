package models

type Comment struct {
	CommentID int    `json:"id" db:"id"`
	PostId    int    `json:"postId" db:"post_id"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	Body      string `json:"body" db:"body"`
}
