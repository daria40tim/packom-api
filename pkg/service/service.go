package service

import (
	"github.com/daria40tim/packom"
	"github.com/daria40tim/packom/pkg/repository"
)

type Authorization interface {
	CreateOrg(org packom.Org) (int, error)
	GenerateToken(login, pwd string) (string, error, packom.Org)
	ParseToken(token string) (int, error)
}

type CP interface {
	Create(O_Id int, cp packom.CPIns) (int, error)
	GetAll(O_Id int /*filter packom.TechFilter*/) ([]packom.CPAll, error)
	GetById(O_Id, cp_id int) (packom.CPId, error)
	UpdateById(cp_id int, input packom.CPIns) (int, error)
	DeleteCal(id int) (int, error)
	DeleteCst(id int) (int, error)
}

type Tender interface {
	Create(O_Id int, tender packom.Tender) (int, error)
	GetAll(O_Id int /*filter packom.TechFilter*/) ([]packom.TenderAll, error)
	GetById(id int) (packom.TenderById, error)
	GetFullCosts(id int) ([]packom.FullCost, error)
}

type Tech interface {
	Create(O_Id int, tech packom.Tech) (int, error)
	GetAll(O_Id int) (packom.TechAllCP, error)
	GetById(O_Id, tz_id int) (packom.Tech, []packom.Cost, []packom.Calendar, error)
	SelectAll() (packom.Select, error)
	DeleteCost(tz_id int, task string) (int, error)
	DeleteCal(tz_id int, task string) (int, error)
	UpdateById(id int, input packom.Tech) (int, error)
}

type Org interface {
	GetAll(O_Id int /*filter packom.TechFilter*/) ([]packom.OrgAll, error)
	GetById(O_Id, o_id int) (packom.OrgId, error)
	UpdateById(O_Id int, input packom.OrgI) (int, error)
	AddById(O_Id, input int) (int, error)
}

type Service struct {
	Authorization
	CP
	Tender
	Tech
	Org
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Autherization),
		Tech:          NewTechService(repos.Tech),
		CP:            NewCPService(repos.CP),
		Tender:        NewTenderService(repos.Tender),
		Org:           NewOrgService(repos.Org),
	}
}
