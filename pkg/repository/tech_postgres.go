package repository

import (
	"github.com/daria40tim/packom"
	"github.com/jmoiron/sqlx"
)

type TechPostgres struct {
	db *sqlx.DB
}

func NewTechPostgres(db *sqlx.DB) *TechPostgres {
	return &TechPostgres{db: db}
}

func (r *TechPostgres) Create(O_Id int, tech packom.Tech) (int, error) {

	var group_id int
	query := `select group_id from public."Pack_groups" where name=$1`
	err := r.db.Get(&group_id, query, tech.Group)
	if err != nil {
		query = `INSERT INTO public."Pack_groups"(group_id, name) VALUES (default, $1) returning  group_id`
		row := r.db.QueryRow(query, tech.Group)
		if err := row.Scan(&group_id); err != nil {
			return 0, err
		}
	}

	var kind_id int
	query = `select kind_id from public."Pack_kinds" where name=$1`
	err = r.db.Get(&kind_id, query, tech.Kind)
	if err != nil {
		query = `INSERT INTO public."Pack_kinds"(kind_id, name) VALUES (default, $1) returning  kind_id`
		row := r.db.QueryRow(query, tech.Kind)
		if err := row.Scan(&kind_id); err != nil {
			return 0, err
		}
	}

	var type_id int
	query = `select type_id from public."Pack_types" where name=$1`
	err = r.db.Get(&type_id, query, tech.Type)
	if err != nil {
		query = `INSERT INTO public."Pack_types"(type_id, name) VALUES (default, $1) returning  type_id`
		row := r.db.QueryRow(query, tech.Type)
		if err := row.Scan(&type_id); err != nil {
			return 0, err
		}
	}

	var pay_cond_id int
	query = `select pay_cond_id from public."Pay_conds" where name=$1`
	err = r.db.Get(&pay_cond_id, query, tech.Pay_cond)
	if err != nil {
		query = `INSERT INTO public."Pay_conds"(pay_cond_id, name) VALUES (default, $1) returning  pay_cond_id`
		row := r.db.QueryRow(query, tech.Pay_cond)
		if err := row.Scan(&pay_cond_id); err != nil {
			return 0, err
		}
	}

	var Tz_Id int
	createTechQuery := `INSERT INTO public."Techs"
	(tz_id, date, o_id, end_date, proj, group_id, kind_id, type_id, tz_st, tender_st, cp_st, pay_cond_id, private, info, history)
	VALUES (default, $1, $2, $3, $4,   $5,       $6,      $7,       $8,    $9,        $10,   $11,         $12,     $13,  $14) returning  tz_id`
	row := r.db.QueryRow(createTechQuery, tech.Date, O_Id, tech.End_date, tech.Proj, group_id, kind_id, type_id, tech.Tz_st, tech.Tender_st, tech.Cp_st, pay_cond_id, tech.Private, tech.Info, tech.History)
	if err := row.Scan(&Tz_Id); err != nil {
		return 0, err
	}

	for _, v := range tech.Calendars {

		var name_id int
		query = `select name_id from public."Task_names" where name=$1`
		err = r.db.Get(&name_id, query, v.Name)
		if err != nil {
			query = `INSERT INTO public."Task_names" (name_id, name) VALUES (default, $1) returning  name_id`
			row := r.db.QueryRow(query, v.Name)
			if err := row.Scan(&name_id); err != nil {
				return 0, err
			}
		}

		var cal_id int
		createTechQuery := `INSERT INTO public."Calendar"(
				cal_id, name_id, period, term, tz_id, cp_id)
		VALUES (default,$1,      $2,      $3,   $4,    $5) returning  cal_id`
		row := r.db.QueryRow(createTechQuery, name_id, v.Period, v.Term, Tz_Id, nil)
		if err := row.Scan(&cal_id); err != nil {
			return 0, err
		}

	}

	for _, v := range tech.Costs {
		var task_id int
		query = `select task_id from public."Tasks" where name=$1`
		err = r.db.Get(&task_id, query, v.Task)
		if err != nil {
			query = `INSERT INTO public."Tasks" (task_id, name) VALUES (default, $1) returning  task_id`
			row := r.db.QueryRow(query, v.Task)
			if err := row.Scan(&task_id); err != nil {
				return 0, err
			}
		}

		var metr_id int
		query = `select metr_id from public."Metrics" where name=$1`
		err = r.db.Get(&metr_id, query, v.Metr)
		if err != nil {
			query = `INSERT INTO public."Metrics"(metr_id, name) VALUES (default, $1) returning  metr_id`
			row := r.db.QueryRow(query, v.Metr)
			if err := row.Scan(&metr_id); err != nil {
				return 0, err
			}
		}

		var cost_id int
		createTechQuery := `INSERT INTO public."Costs"(
			cost_id, metr_id, count, tz_id, cp_id, ppu, info, task_id)
	VALUES (default, $1,      $2,     $3,   $4,    $5,   $6,  $7) returning  cost_id`
		row := r.db.QueryRow(createTechQuery, metr_id, v.Count, Tz_Id, nil, nil, v.Info, task_id)
		if err := row.Scan(&cost_id); err != nil {
			return 0, err
		}
	}

	return Tz_Id, nil
}
