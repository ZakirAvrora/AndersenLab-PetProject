package repository

import "github.com/ZakirAvrora/AndersenLab-PetProject/internal/models"

type Saver interface {
	PostRepo
	UserRepo
	CommentRepo
}

type PostRepo interface {
	SavePost(post models.Post) error
}

type UserRepo interface {
	SaveUser(user models.User) error
}

type CommentRepo interface {
	SaveComment(comment models.Comment) error
}
