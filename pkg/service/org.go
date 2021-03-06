package service

import (
	"github.com/daria40tim/packom"
	"github.com/daria40tim/packom/pkg/repository"
)

type OrgService struct {
	repo repository.Org
}

func NewOrgService(repo repository.Org) *OrgService {
	return &OrgService{repo: repo}
}

func (s *OrgService) GetAll(O_Id int) ([]packom.OrgAll, error) {
	return s.repo.GetAll(O_Id)
}

func (s *OrgService) GetById(O_Id, o_id int) (packom.OrgId, error) {
	return s.repo.GetById(O_Id, o_id)
}

func (s *OrgService) UpdateById(O_Id int, input packom.OrgI) (int, error) {
	if input.Pwd != "" {
		input.Pwd = generatePasswordHash(input.Pwd)
	}
	return s.repo.UpdateById(O_Id, input)
}

func (s *OrgService) AddById(O_Id, input int) (int, error) {
	return s.repo.AddById(O_Id, input)
}

func (s *OrgService) SelectAllSpecs() (packom.Specs, error) {
	return s.repo.SelectAllSpecs()
}

func (s *OrgService) AddDoc(name string, o_id int) error {
	return s.repo.AddDoc(name, o_id)
}

func (s *OrgService) DeleteTrustedOrg(O_Id, id int) error {
	return s.repo.DeleteTrustedOrg(O_Id, id)
}

func (s *OrgService) GetFilterData() (packom.OrgFilterData, error) {
	return s.repo.GetFilterData()
}

func (s *OrgService) GetAllFiltered(O_Id int, names, groups, specs, countries []int) ([]packom.OrgAll, error) {
	return s.repo.GetAllFiltered(O_Id, names, groups, specs, countries)
}
