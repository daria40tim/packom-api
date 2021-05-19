package service

import (
	"github.com/daria40tim/packom"
	"github.com/daria40tim/packom/pkg/repository"
)

type Authorization interface {
	CreateOrg(org packom.Org) (int, error)
	GenerateToken(login, pwd string) (string, error)
	ParseToken(token string) (int, error)
}

type CP interface {
	Create(O_Id int, cp packom.CP) (int, error)
}

type Tender interface {
}

type Tech interface {
	Create(O_Id int, tech packom.Tech) (int, error)
}

type Org interface {
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
	}
}
