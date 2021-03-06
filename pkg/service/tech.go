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

func (s *TechService) GetAll(O_Id int) ([]packom.TechAll, error) {
	return s.repo.GetAll(O_Id)
}

func (s *TechService) GetById(O_Id, tz_id int) (packom.Tech, []packom.Cost, []packom.Calendar, error) {
	return s.repo.GetById(O_Id, tz_id)
}

func (s *TechService) SelectAll() (packom.Select, error) {
	return s.repo.SelectAll()
}

func (s *TechService) DeleteCost(tz_id int, task, history string) (int, error) {
	return s.repo.DeleteCost(tz_id, task, history)
}

func (s *TechService) DeleteCal(tz_id int, task, history string) (int, error) {
	return s.repo.DeleteCal(tz_id, task, history)
}

func (s *TechService) UpdateById(id int, input packom.Tech) (int, error) {
	return s.repo.UpdateById(id, input)
}

func (s *TechService) AddTechDoc(name string, o_id, tz_id int) error {
	return s.repo.AddTechDoc(name, o_id, tz_id)
}

func (s *TechService) GetFilterData() (packom.TechFilterData, error) {
	return s.repo.GetFilterData()
}

func (s *TechService) GetAllTechsFiltered(O_Id int, EDate, SDate string, Clients, Projs, TZ_STS, CP_STS, Tender_STS []int) ([]packom.TechAll, error) {
	return s.repo.GetAllTechsFiltered(O_Id, EDate, SDate, Clients, Projs, TZ_STS, CP_STS, Tender_STS)
}
