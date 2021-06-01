package repository

import (
	"github.com/daria40tim/packom"
	"github.com/jmoiron/sqlx"
)

type OrgPostgres struct {
	db *sqlx.DB
}

func NewOrgPostgres(db *sqlx.DB) *OrgPostgres {
	return &OrgPostgres{db: db}
}

func (r *OrgPostgres) GetAll(O_Id int) ([]packom.OrgAll, error) {
	var techs []packom.OrgAll

	query := `SELECT public."Orgs".name, public."Orgs".o_id, case group_id when 1 then 'Клиент' when 2 then 'Поставщик' else 'Клиент, Поставщик' end as group, site, phone, email, public."Specs".name as specs, public."Countries".name as countries 
	FROM public."Orgs"
	left join public."Orgs_specs" on public."Orgs".o_id=public."Orgs_specs".o_id
	left join public."Specs" on public."Specs".spec_id=public."Orgs_specs".spec_id
	left join public."Org_countries" on public."Orgs".o_id=public."Org_countries".o_id
	left join public."Countries" on public."Countries".country_id=public."Org_countries".country_id`

	err := r.db.Select(&techs, query)

	return techs, err
}

func (r *OrgPostgres) GetById(O_Id, o_id int) (packom.OrgId, error) {
	var org packom.OrgId
	var orgs []packom.OrgAll
	var e_orgs []packom.OrgAll
	var docs []string
	var e_docs []string

	query := `SELECT public."Orgs".name, public."Orgs".o_id, case group_id when 1 then 'Поставщик' when 2 then 'Клиент' else 'Клиент, Поставщик' end as group, site, phone, email, public."Specs".name as specs, public."Countries".name as countries 
	FROM public."Orgs"
	left join public."Orgs_specs" on public."Orgs".o_id=public."Orgs_specs".o_id
	left join public."Specs" on public."Specs".spec_id=public."Orgs_specs".spec_id
	left join public."Org_countries" on public."Orgs".o_id=public."Org_countries".o_id
	left join public."Countries" on public."Countries".country_id=public."Org_countries".country_id
	where public."Orgs".o_id in (select f_o_id from public."Orgs_orgs" where public."Orgs_orgs".o_id = $1)`

	err := r.db.Select(&orgs, query, o_id)
	if err != nil {
		return org, err
	}

	if orgs == nil {
		org.Orgs = e_orgs
	} else {
		org.Orgs = orgs
	}

	query = `SELECT file_name
	FROM public."Orgs_docs"
	where o_id = $1;`

	err = r.db.Select(&docs, query, o_id)
	if err != nil {
		return org, err
	}

	if docs == nil {
		org.Docs = e_docs
	} else {
		org.Docs = docs
	}

	query = `SELECT public."Orgs".o_id, public."Orgs".login, public."Orgs".name, case group_id when 1 then 'Клиент' when 2 then 'Поставщик' else 'Клиент, Поставщик' end as group, site, phone, email, 
	adress, '' as info, status, '' as history, public."Specs".name as specs, public."Countries".name as countries
		FROM public."Orgs"
		join public."Orgs_specs" on  public."Orgs_specs".o_id=public."Orgs".o_id
		join public."Specs" on  public."Specs".spec_id=public."Orgs_specs".spec_id
		join public."Org_countries" on  public."Org_countries".o_id=public."Orgs".o_id
		join public."Countries" on  public."Countries".country_id=public."Org_countries".country_id
		where public."Orgs".o_id = $1`

	err = r.db.Get(&org, query, o_id)

	return org, err
}

func (r *OrgPostgres) UpdateById(O_Id int, input packom.OrgI) (int, error) {
	var id int

	query := `UPDATE public."Orgs"
	SET site=$2, phone=$3, email=$4, adress=$5, info=$6, hashed_pwd=$7
	WHERE o_id = $1 
	returning o_id;`

	row := r.db.QueryRow(query, O_Id, input.Site, input.Phone, input.Email, input.Adress, input.Info, input.Pwd)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	var spec_id int
	query = `select spec_id from public."Specs" where name=$1`
	err := r.db.Get(&spec_id, query, input.Specs)
	if err != nil {
		query = `INSERT INTO public."Specs" (spec_id, name) VALUES (default, $1) returning  spec_id`
		row := r.db.QueryRow(query, input.Specs)
		if err := row.Scan(&spec_id); err != nil {
			return 0, err
		}
		query = `UPDATE public."Orgs_specs" SET  spec_id=$1
		WHERE o_id = $2`
		row = r.db.QueryRow(query, spec_id, O_Id)
		if err := row.Scan(); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *OrgPostgres) AddById(O_Id, input int) (int, error) {
	var id int

	query := `INSERT INTO public."Orgs_orgs"(
		o_id, f_o_id)
		VALUES ($1, $2) returning f_o_id;`

	row := r.db.QueryRow(query, O_Id, input)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
