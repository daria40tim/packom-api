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
}

type CP interface {
	Create(O_Id int, cp packom.CP) (int, error)
	GetAll(O_Id int /*, filter packom.TechFilter*/) ([]packom.CPAll, error)
	GetById(O_Id, cp_id int) (packom.CP, error)
}

type Tender interface {
	Create(O_Id int, tender packom.Tender) (int, error)
	GetAll(O_Id int /*, filter packom.TechFilter*/) ([]packom.TenderAll, error)
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
