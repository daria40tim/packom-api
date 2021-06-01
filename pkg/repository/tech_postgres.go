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

func (r *TechPostgres) GetAll(O_Id int) ([]packom.TechAll, []packom.CP_srv, error) {
	var cps []packom.CP_srv

	query := `SELECT cp_id, date, tz_id, o_id, end_date
	FROM public."CP" where o_id=$1`

	err := r.db.Select(&cps, query, O_Id)
	if err != nil {
		return nil, nil, err
	}

	var techs []packom.TechAll

	query = `SELECT distinct public."Techs".date,  public."Techs".task_name as task,public."Techs".end_date, public."Orgs".name as client, public."Techs".o_id, public."Techs".tz_id,  public."Techs".end_date, 
	public."Techs".proj, public."Pack_groups".name as group, public."Pack_types".name as type, public."Pack_kinds".name as kind, 
	count(cp_id) over (partition by public."Techs".tz_id) as cp_count
	FROM public."Techs" join public."Orgs" on public."Techs".o_id=public."Orgs".o_id 
	join public."Pack_groups" on public."Pack_groups".group_id=public."Techs".group_id 
	join public."Pack_types" on public."Pack_types".type_id=public."Techs".type_id 
	join public."Pack_kinds" on public."Pack_kinds".kind_id=public."Techs".kind_id 
	left join public."CP" on public."CP".tz_id=public."Techs".tz_id `

	err = r.db.Select(&techs, query)

	return techs, cps, err
}

func (r *TechPostgres) GetById(O_Id, tz_id int) (packom.Tech, []packom.Cost, []packom.Calendar, error) {
	var costs []packom.Cost
	var tech packom.Tech
	var calendars []packom.Calendar

	query := `SELECT cost_id, public."Metrics".name as metr, count, tz_id, 0 as cp_id, 0 as ppu, info, public."Tasks".name as task
	FROM public."Costs"
	join public."Metrics" on public."Metrics".metr_id = public."Costs".metr_id
	join public."Tasks" on public."Tasks".task_id = public."Costs".task_id
	where tz_id = $1`

	err := r.db.Select(&costs, query, tz_id)
	if err != nil {
		return tech, nil, nil, err
	}

	query = `SELECT cal_id, public."Task_names".name as name, period, term, tz_id, 0 as cp_id
	FROM public."Calendar"
	join public."Task_names" on public."Task_names".name_id = public."Calendar".name_id
	where tz_id = $1`

	err = r.db.Select(&calendars, query, tz_id)
	if err != nil {
		return tech, nil, nil, err
	}

	query = `SELECT public."Techs".date, public."Techs".o_id, public."Pack_groups".name as group, public."Pack_kinds".name as kind, 
	public."Pack_types".name as type, public."Orgs".name as client, public."Techs".end_date, public."Techs".proj, public."Pay_conds".name as pay_cond,
	tender_st, cp_st, private, public."Techs".info, public."Techs".history, public."Techs".tz_st, public."Techs".tz_id
		FROM public."Techs"
		join public."Orgs" on public."Orgs".o_id = public."Techs".o_id
		join public."Pack_groups" on public."Pack_groups".group_id = public."Techs".group_id
		join public."Pack_kinds" on public."Pack_kinds".kind_id = public."Techs".kind_id
		join public."Pack_types" on public."Pack_types".type_id = public."Techs".type_id
		join public."Pay_conds" on public."Pay_conds".pay_cond_id = public."Techs".type_id
		where tz_id = $1`

	err = r.db.Get(&tech, query, tz_id)

	return tech, costs, calendars, err
}
