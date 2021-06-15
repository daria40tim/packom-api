package service

import (
	"github.com/daria40tim/packom"
	"github.com/daria40tim/packom/pkg/repository"
)

type CPService struct {
	repo repository.CP
}

func NewCPService(repo repository.CP) *CPService {
	return &CPService{repo: repo}
}

func (s *CPService) Create(O_Id int, cp packom.CPIns) (int, error) {
	return s.repo.Create(O_Id, cp)
}

func (s *CPService) GetAll(O_Id int) ([]packom.CPAll, error) {
	return s.repo.GetAll(O_Id)
}

func (s *CPService) GetById(O_Id, cp_id int) (packom.CPId, error) {
	return s.repo.GetById(O_Id, cp_id)
}

func (s *CPService) UpdateById(cp_id int, input packom.CPIns) (int, error) {
	return s.repo.UpdateById(cp_id, input)
}

func (s *CPService) DeleteCal(id int) (int, error) {
	return s.repo.DeleteCal(id)
}

func (s *CPService) DeleteCst(id int) (int, error) {
	return s.repo.DeleteCst(id)
}

func (s *CPService) AddCPDoc(name string, o_id, cp_id int) error {
	return s.repo.AddCPDoc(name, o_id, cp_id)
}
