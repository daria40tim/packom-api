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
}

type CP interface {
	Create(O_Id int, cp packom.CP) (int, error)
}

type Tender interface {
}

type Tech interface {
	Create(O_Id int, tech packom.Tech) (int, error)
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
	}
}
