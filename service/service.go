package service

import "project/repository"

type Service struct {
}

func NewService(repo repository.Repository) Service {
	return Service{}
}
