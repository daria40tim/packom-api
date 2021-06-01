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

func (s *TechService) GetAll(O_Id int /*filter packom.TechFilter*/) ([]packom.TechAll, []packom.CP_srv, error) {
	return s.repo.GetAll(O_Id /*filter*/)
}

func (s *TechService) GetById(O_Id, tz_id int) (packom.Tech, []packom.Cost, []packom.Calendar, error) {
	return s.repo.GetById(O_Id, tz_id)
}
