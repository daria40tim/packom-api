package service

import (
	"github.com/daria40tim/packom"
	"github.com/daria40tim/packom/pkg/repository"
)

type TenderService struct {
	repo repository.Tender
}

func NewTenderService(repo repository.Tender) *TenderService {
	return &TenderService{repo: repo}
}

func (s *TenderService) Create(O_Id int, tender packom.Tender) (int, error) {
	return s.repo.Create(O_Id, tender)
}

func (s *TenderService) GetAll(O_Id int) ([]packom.TenderAll, error) {
	return s.repo.GetAll(O_Id)
}

func (s *TenderService) GetById(id int) (packom.TenderById, error) {
	return s.repo.GetById(id)
}

func (s *TenderService) GetFullCosts(id int) ([]packom.FullCost, error) {
	return s.repo.GetFullCosts(id)
}

func (s *TenderService) UpdateById(input packom.Tender) (int, error) {
	return s.repo.UpdateById(input)
}

func (s *TenderService) GetTenderFilterData() (packom.TenderFilterData, error) {
	return s.repo.GetTenderFilterData()
}

func (s *TenderService) GetAllTendersFiltered(O_Id int, EDate, SDate string, Projs, TZ_Ids, Tender_STS []int) ([]packom.TenderAll, error) {
	return s.repo.GetAllTendersFiltered(O_Id, EDate, SDate, Projs, TZ_Ids, Tender_STS)
}
