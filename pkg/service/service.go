package service

import (
	"AccountApp"
	"AccountApp/pkg/repository"
)

type Authorization interface {
	CreateUser(user AccountApp.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(soken string) (int, error)
}
type User interface {
}
type UserList interface {
}

type Service struct {
	Authorization
	User
	UserList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
