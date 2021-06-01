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

func (s *CPService) Create(O_Id int, cp packom.CP) (int, error) {
	return s.repo.Create(O_Id, cp)
}

func (s *CPService) GetAll(O_Id int) ([]packom.CPAll, error) {
	return s.repo.GetAll(O_Id)
}

func (s *CPService) GetById(O_Id, cp_id int) (packom.CP, error) {
	return s.repo.GetById(O_Id, cp_id)
}
