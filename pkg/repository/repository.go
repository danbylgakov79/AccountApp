package repository

import (
	"AccountApp"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user AccountApp.User) (int, error)
	GetUser(username, password string) (AccountApp.User, error)
}
type User interface {
}
type UserList interface {
}

type Repository struct {
	Authorization
	User
	UserList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorization: NewAuthPostgres(db)}
}
