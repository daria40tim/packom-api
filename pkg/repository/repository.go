package repository

import (
	"github.com/daria40tim/packom"
	"github.com/jmoiron/sqlx"
)

type Autherization interface {
	CreateOrg(org packom.Org) (int, error)
	GetOrg(login, pwd string) (packom.Org, error)
}

type Org interface {
	GetAll(O_Id int /*, filter packom.TechFilter*/) ([]packom.OrgAll, error)
	GetById(O_Id, o_id int) (packom.OrgId, error)
	UpdateById(O_Id int, input packom.OrgI) (int, error)
	AddById(O_Id int, input int) (int, error)
	SelectAllSpecs() (packom.Specs, error)
	AddDoc(name string, o_id int) error
}

type CP interface {
	Create(O_Id int, cp packom.CPIns) (int, error)
	GetAll(O_Id int /*, filter packom.TechFilter*/) ([]packom.CPAll, error)
	GetById(O_Id, cp_id int) (packom.CPId, error)
	UpdateById(cp_id int, input packom.CPIns) (int, error)
	DeleteCst(id int) (int, error)
	DeleteCal(id int) (int, error)
	AddCPDoc(name string, o_id, cp_id int) error
}

type Tender interface {
	Create(O_Id int, tender packom.Tender) (int, error)
	GetAll(O_Id int /*, filter packom.TechFilter*/) ([]packom.TenderAll, error)
	GetById(id int) (packom.TenderById, error)
	GetFullCosts(id int) ([]packom.FullCost, error)
	UpdateById(input packom.Tender) (int, error)
}

type Tech interface {
	Create(O_Id int, tech packom.Tech) (int, error)
	GetAll(O_Id int) (packom.TechAllCP, error)
	GetById(O_Id, tz_id int) (packom.Tech, []packom.Cost, []packom.Calendar, error)
	SelectAll() (packom.Select, error)
	DeleteCost(tz_id int, task, history string) (int, error)
	DeleteCal(tz_id int, task, history string) (int, error)
	UpdateById(id int, input packom.Tech) (int, error)
	AddTechDoc(name string, o_id, tz_id int) error
}

type Repository struct {
	Autherization
	Tech
	Tender
	CP
	Org
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autherization: NewAuthPostgres(db),
		Tech:          NewTechPostgres(db),
		CP:            NewCPPostgres(db),
		Tender:        NewTenderPostgres(db),
		Org:           NewOrgPostgres(db),
	}
}
