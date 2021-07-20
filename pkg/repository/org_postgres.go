package repository

import (
	"strconv"

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

	query := `SELECT public."Orgs".name, public."Orgs".o_id, case group_id when 1 then 'Клиент' when 2 then 'Поставщик' else 'Клиент, Поставщик' end as group, site, phone, email, public."Countries".name as countries 
	FROM public."Orgs"
	left join public."Org_countries" on public."Orgs".o_id=public."Org_countries".o_id
	left join public."Countries" on public."Countries".country_id=public."Org_countries".country_id`

	if err := r.db.Select(&techs, query); err != nil {
		return nil, err
	}

	for i, v := range techs {
		var specs []string
		var spec string

		query = `SELECT name
		FROM public."Orgs_specs"
		join public."Specs" on public."Specs".spec_id = public."Orgs_specs".spec_id
		where o_id = $1 and active`

		if err := r.db.Select(&specs, query, v.O_id); err != nil {
			return nil, err
		}

		if len(specs) > 1 {
			spec = specs[0]
			for i := 1; i < len(specs); i++ {
				spec = spec + ", " + specs[i]
			}
		} else if len(specs) > 0 {
			spec = specs[0]
		}

		techs[i].Specs = spec
	}

	return techs, nil
}

func (r *OrgPostgres) GetById(O_Id, o_id int) (packom.OrgId, error) {
	var org packom.OrgId
	var orgs []packom.OrgAll
	var e_orgs []packom.OrgAll
	var docs []string
	var e_docs []string
	var specs []string
	var trusted []int

	query := `SELECT public."Orgs".name, public."Orgs".o_id, case group_id when 1 then 'Клиент' when 2 then 'Поставщик' else 'Клиент, Поставщик' end as group, site, phone, email, public."Countries".name as countries
	FROM public."Orgs"
	left join public."Org_countries" on public."Orgs".o_id=public."Org_countries".o_id
	left join public."Countries" on public."Countries".country_id=public."Org_countries".country_id
	where public."Orgs".o_id in (select f_o_id from public."Orgs_orgs" where public."Orgs_orgs".o_id = $1 and active)`

	err := r.db.Select(&orgs, query, o_id)
	if err != nil {
		return org, err
	}

	for i, v := range orgs {
		var spec string
		var specs []string

		query = `SELECT name
		FROM public."Orgs_specs"
		join public."Specs" on public."Specs".spec_id = public."Orgs_specs".spec_id
		where o_id = $1 and active`

		if err := r.db.Select(&specs, query, v.O_id); err != nil {
			return org, err
		}

		if len(specs) > 1 {
			spec = specs[0]
			for i := 1; i < len(specs); i++ {
				spec = spec + ", " + specs[i]
			}
		} else if len(specs) > 0 {
			spec = specs[0]
		}

		orgs[i].Specs = spec
	}

	if orgs == nil {
		org.Orgs = e_orgs
	} else {
		org.Orgs = orgs
	}

	query = `SELECT distinct f_o_id
	FROM public."Orgs_orgs"
	where o_id = $1 and active`

	err = r.db.Select(&trusted, query, O_Id)
	if err != nil {
		return org, err
	}

	org.Trusted = trusted

	query = `SELECT distinct file_name
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

	query = `SELECT distinct public."Specs".name
	FROM public."Orgs_specs"
	join public."Specs" on public."Specs".spec_id = public."Orgs_specs".spec_id 
	where o_id = $1 and active`

	err = r.db.Select(&specs, query, o_id)
	if err != nil {
		return org, err
	}

	if specs == nil {
		org.Specs = nil
	} else {
		org.Specs = specs
	}

	query = `SELECT public."Orgs".o_id, public."Orgs".login, public."Orgs".name, case group_id when 1 then 'Клиент' when 2 then 'Поставщик' else 'Клиент, Поставщик' end as group, site, phone, email, 
	adress, info, status, public."Countries".name as countries, history
		FROM public."Orgs"
		join public."Org_countries" on  public."Org_countries".o_id=public."Orgs".o_id
		join public."Countries" on  public."Countries".country_id=public."Org_countries".country_id
		where public."Orgs".o_id = $1`

	err = r.db.Get(&org, query, o_id)

	return org, err
}

func (r *OrgPostgres) UpdateById(O_Id int, input packom.OrgI) (int, error) {
	var id int

	if input.Pwd != "" {
		query := `UPDATE public."Orgs"
	SET site=$2, phone=$3, email=$4, adress=$5, info=$6, hashed_pwd=$7, history=$8
	WHERE o_id = $1 
	returning o_id;`

		row := r.db.QueryRow(query, O_Id, input.Site, input.Phone, input.Email, input.Adress, input.Info, input.Pwd, input.History)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}

		for _, v := range input.Specs {

			var spec_id int
			var id int

			query = `select spec_id from public."Specs" where name=$1`
			/*err := */ r.db.Get(&spec_id, query, v.Name) /*err != nil {
				return 0, err
			}*/

			if spec_id == 0 {
				query = `INSERT INTO public."Specs" (spec_id, name) VALUES (default, $1) returning  spec_id`
				row := r.db.QueryRow(query, v.Name)
				if err := row.Scan(&spec_id); err != nil {
					return 0, err
				}
			}

			query = `UPDATE public."Orgs_specs" SET  active=$1
			WHERE o_id = $2 and spec_id=$3 returning o_id`
			row = r.db.QueryRow(query, v.Active, O_Id, spec_id)
			/*if err := */ row.Scan(&id) /*err != nil {
				return 0, err
			}*/
			if id == 0 {
				query = `INSERT INTO public."Orgs_specs"(
					o_id, spec_id, active)
					VALUES ($1, $2, $3) returning o_id`
				row = r.db.QueryRow(query, O_Id, spec_id, v.Active)
				if err := row.Scan(&id); err != nil {
					return 0, err
				}
			}

		}
	} else {

		query := `UPDATE public."Orgs"
	SET site=$1, phone=$2, email=$3, adress=$4, info=$5, history=$6
	WHERE o_id = $7 
	returning o_id;`

		row := r.db.QueryRow(query, input.Site, input.Phone, input.Email, input.Adress, input.Info, input.History, O_Id)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}

		for _, v := range input.Specs {

			var spec_id int
			var id int

			query = `select spec_id from public."Specs" where name=$1`
			/*err := */ r.db.Get(&spec_id, query, v.Name) /*err != nil {
				return 0, err
			}*/

			if spec_id == 0 {
				query = `INSERT INTO public."Specs" (spec_id, name) VALUES (default, $1) returning  spec_id`
				row := r.db.QueryRow(query, v.Name)
				if err := row.Scan(&spec_id); err != nil {
					return 0, err
				}
			}

			query = `UPDATE public."Orgs_specs" SET  active=$1
			WHERE o_id = $2 and spec_id=$3 returning o_id`
			row = r.db.QueryRow(query, v.Active, O_Id, spec_id)
			/*if err := */ row.Scan(&id) /*err != nil {
				return 0, err
			}*/
			if id == 0 {
				query = `INSERT INTO public."Orgs_specs"(
					o_id, spec_id, active)
					VALUES ($1, $2, $3) returning o_id`
				row = r.db.QueryRow(query, O_Id, spec_id, v.Active)
				if err := row.Scan(&id); err != nil {
					return 0, err
				}
			}

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

func (r *OrgPostgres) SelectAllSpecs() (packom.Specs, error) {
	var specs []string
	var res packom.Specs

	query := `SELECT name FROM public."Specs"`

	err := r.db.Select(&specs, query)
	if err != nil {
		return res, err
	}

	res.Specs = specs

	return res, nil
}

func (r *OrgPostgres) AddDoc(name string, o_id int) error {
	var id int

	query := `INSERT INTO public."Orgs_docs"(
		doc_id, file_path, file_name, o_id)
		VALUES (default, '', $1, $2) returning doc_id`

	row := r.db.QueryRow(query, name, o_id)
	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}

func (r *OrgPostgres) DeleteTrustedOrg(O_Id, id int) error {
	query := `UPDATE public."Orgs_orgs"
	SET  active=false
	WHERE o_id = $1 and f_o_id=$2`

	row := r.db.QueryRow(query, O_Id, id)
	if err := row.Scan(); err != nil {
		return err
	}
	return nil
}

func (r *OrgPostgres) GetFilterData() (packom.OrgFilterData, error) {
	var res packom.OrgFilterData
	var names []packom.NameFilter
	var countries []packom.CountryFilter
	var specs []packom.SpecFilter

	query := `SELECT distinct o_id, name
	FROM public."Orgs" order by name`

	err := r.db.Select(&names, query)
	if err != nil {
		return res, err
	}

	query = `SELECT distinct country_id, name
	FROM public."Countries" order by name`

	err = r.db.Select(&countries, query)
	if err != nil {
		return res, err
	}

	query = `SELECT distinct spec_id, name
	FROM public."Specs" order by name`

	err = r.db.Select(&specs, query)
	if err != nil {
		return res, err
	}

	res.Countries = countries
	res.Names = names
	res.Specs = specs

	return res, nil
}

func (r *OrgPostgres) GetAllFiltered(O_Id int, names, groups, specs, countries []int) ([]packom.OrgAll, error) {
	var techs []packom.OrgAll

	query := `SELECT public."Orgs".name, public."Orgs".o_id, case group_id when 1 then 'Клиент' when 2 then 'Поставщик' else 'Клиент, Поставщик' end as group, site, phone, email, public."Countries".name as countries
	FROM public."Orgs"
	left join public."Org_countries" on public."Orgs".o_id=public."Org_countries".o_id
	left join public."Countries" on public."Countries".country_id=public."Org_countries".country_id`

	if len(names) == 0 {
		query += ` where public."Orgs".o_id in (select o_id from public."Orgs") `
	} else {
		query += ` where public."Orgs".o_id in ( ` + strconv.Itoa(names[0])
		for i := 1; i < len(names); i++ {
			query += `, ` + strconv.Itoa(names[i])
		}
		query += `)`
	}

	if len(groups) == 0 {
		query += ` and public."Orgs".group_id in (select group_id from public."Orgs") `
	} else {
		query += ` and public."Orgs".group_id in ( ` + strconv.Itoa(groups[0])
		for i := 1; i < len(groups); i++ {
			query += `, ` + strconv.Itoa(groups[i])
		}
		query += `)`
	}
	if len(countries) == 0 {
		query += ` and public."Org_countries".country_id in (select  public."Org_countries".country_id from public."Org_countries") `
	} else {
		query += ` and public."Org_countries".country_id in ( ` + strconv.Itoa(countries[0])
		for i := 1; i < len(countries); i++ {
			query += `, ` + strconv.Itoa(countries[i])
		}
		query += `)`
	}
	if len(specs) == 0 {
		query += ``
	} else {
		query += ` and public."Orgs".o_id in ( SELECT o_id FROM public."Orgs_specs" where spec_id in ( ` + strconv.Itoa(specs[0])
		for i := 1; i < len(specs); i++ {
			query += `, ` + strconv.Itoa(specs[i])
		}
		query += `) and active)`
	}

	if err := r.db.Select(&techs, query); err != nil {
		return nil, err
	}

	for i, v := range techs {
		var specs []string
		var spec string

		query = `SELECT name
		FROM public."Orgs_specs"
		join public."Specs" on public."Specs".spec_id = public."Orgs_specs".spec_id
		where o_id = $1 and active`

		if err := r.db.Select(&specs, query, v.O_id); err != nil {
			return nil, err
		}

		if len(specs) > 1 {
			spec = specs[0]
			for i := 1; i < len(specs); i++ {
				spec = spec + ", " + specs[i]
			}
		} else if len(specs) > 0 {
			spec = specs[0]
		}

		techs[i].Specs = spec
	}

	return techs, nil
}
