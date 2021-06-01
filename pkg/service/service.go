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
	Create(O_Id int, cp packom.CP) (int, error)
	GetAll(O_Id int /*filter packom.TechFilter*/) ([]packom.CPAll, error)
	GetById(O_Id, cp_id int) (packom.CP, error)
}

type Tender interface {
	Create(O_Id int, tender packom.Tender) (int, error)
	GetAll(O_Id int /*filter packom.TechFilter*/) ([]packom.TenderAll, error)
}

type Tech interface {
	Create(O_Id int, tech packom.Tech) (int, error)
	GetAll(O_Id int) ([]packom.TechAll, []packom.CP_srv, error)
	GetById(O_Id, tz_id int) (packom.Tech, []packom.Cost, []packom.Calendar, error)
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
