package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	PostService PostService
	UserService UserService
}

type PostService interface {
	Create(context.Context, *Post) error
}

type UserService interface {
	Create(context.Context, *User) error
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		PostService: &PostServiceImpl{db},
		UserService: &UserServiceImpl{db},
	}
}
