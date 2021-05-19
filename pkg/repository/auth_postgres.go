package repository

import (
	"github.com/daria40tim/packom"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateOrg(org packom.Org) (int, error) {
	var o_id int
	query := `insert into public."Orgs" (o_id, login, hashed_pwd, name, group_id, status) values (default, $1, $2, '', 0, false) returning o_id`

	row := r.db.QueryRow(query, org.Login, org.Pwd)
	if err := row.Scan(&o_id); err != nil {
		return 0, err
	}

	return o_id, nil
}

func (r *AuthPostgres) GetOrg(login, pwd string) (packom.Org, error) {
	var org packom.Org
	query := `select o_id from public."Orgs" where login=$1 and hashed_pwd=$2`
	err := r.db.Get(&org, query, login, pwd)

	return org, err
}
