package service

import (
	"github.com/daria40tim/packom"
	"github.com/daria40tim/packom/pkg/repository"
)

type TechService struct {
	repo repository.Tech
}

func NewTechService(repo repository.Tech) *TechService {
	return &TechService{repo: repo}
}

func (s *TechService) Create(O_Id int, tech packom.Tech) (int, error) {
	return s.repo.Create(O_Id, tech)
}
