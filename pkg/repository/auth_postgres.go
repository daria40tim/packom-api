package repository

import (
	"os"
	"strconv"

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
	query := `insert into public."Orgs" (o_id, login, hashed_pwd, name, group_id, status, site, phone, email, adress, info, history) 
	values (default, $1, $2, $3, $4, false, '', '', '', '', '', '') returning o_id`

	row := r.db.QueryRow(query, org.Login, org.Pwd, org.Name, org.Group)
	if err := row.Scan(&o_id); err != nil {
		return 0, err
	}

	err := os.MkdirAll("assets/"+strconv.Itoa(o_id), 0777)
	if err != nil {
		return 0, err
	}

	var country_id int
	query = `SELECT country_id
	FROM public."Countries" where name = $1`
	err = r.db.Get(&country_id, query, org.Countries)

	if err != nil {
		return 0, err
	}

	query = `insert into public."Org_countries" (o_id, country_id) values ($1, $2) returning o_id`

	row = r.db.QueryRow(query, o_id, country_id)
	if err := row.Scan(&o_id); err != nil {
		return 0, err
	}

	query = `INSERT INTO public."Orgs_specs"(
		o_id, spec_id)
		VALUES ($1, 0) returning o_id`
	row = r.db.QueryRow(query, o_id)
	if err := row.Scan(&country_id); err != nil {
		return 0, err
	}

	return o_id, nil
}

func (r *AuthPostgres) GetOrg(login, pwd string) (packom.Org, error) {
	var org packom.Org
	query := `select o_id, group_id, name from public."Orgs" where login=$1 and hashed_pwd=$2`
	err := r.db.Get(&org, query, login, pwd)

	return org, err
}

func (r *AuthPostgres) SelectAllCountries() (packom.Countries, error) {
	var countries []string
	var res packom.Countries

	query := `SELECT name FROM public."Countries"`

	err := r.db.Select(&countries, query)
	if err != nil {
		return res, err
	}

	res.Countries = countries

	return res, nil
}

func (r *AuthPostgres) SelectLogin(input string) (packom.Countries, error) {
	var countries []string
	var res packom.Countries

	query := `SELECT distinct login FROM public."Orgs"`

	err := r.db.Select(&countries, query)
	if err != nil {
		return res, err
	}

	res.Countries = countries

	return res, nil
}
