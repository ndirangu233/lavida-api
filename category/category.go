package category

import (
	"github.com/muathendirangu/lavida-api/domains"
)

type service struct {
	repo domains.Repository
}

//NewService creates the category service/usecase
func NewService(userRepo domains.Repository) Service {
	return &service{
		repo: userRepo,
	}
}

func (s service) Store(name string) domains.Response {
	return s.repo.Store(name)
}
