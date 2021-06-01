package repository

import (
	"github.com/daria40tim/packom"
	"github.com/jmoiron/sqlx"
)

type CPPostgres struct {
	db *sqlx.DB
}

func NewCPPostgres(db *sqlx.DB) *CPPostgres {
	return &CPPostgres{db: db}
}

func (r *CPPostgres) GetAll(O_Id int /*, filter packom.TechFilter*/) ([]packom.CPAll, error) {
	var techs []packom.CPAll

	query := `SELECT cp_id, public."CP".date, case public."CP".cp_st when 1 then 'Обрабатывается' when 2 then 'Принято' when 3 then 'Отклонено' end as cp_st, public."CP".tz_id, public."CP".proj, public."CP".o_id, public."Orgs".name as client, public."Pack_groups".name as group, public."Pack_types".name as type, public."Pack_kinds".name as kind  
	FROM public."CP" join public."Orgs" on public."Orgs".o_id=public."CP".o_id 
	join public."Techs" on public."Techs".tz_id=public."CP".tz_id
	join public."Pack_groups" on public."Pack_groups".group_id=public."Techs".group_id
	join public."Pack_kinds" on public."Pack_kinds".kind_id=public."Techs".kind_id
	join public."Pack_types" on public."Pack_types".type_id=public."Techs".type_id
	where public."CP".o_id =$1;`

	err := r.db.Select(&techs, query, O_Id)

	return techs, err
}

func (r *CPPostgres) Create(O_Id int, cp packom.CP) (int, error) {
	var pay_cond_id int
	query := `select pay_cond_id from public."Pay_conds" where name=$1`
	err := r.db.Get(&pay_cond_id, query, cp.Pay_cond)
	if err != nil {
		query = `INSERT INTO public."Pay_conds"(pay_cond_id, name) VALUES (default, $1) returning  pay_cond_id`
		row := r.db.QueryRow(query, cp.Pay_cond)
		if err := row.Scan(&pay_cond_id); err != nil {
			return 0, err
		}
	}

	var CP_Id int
	createTechQuery := `INSERT INTO public."CP"(
		cp_id, date, cp_st, tz_id, proj, o_id, pay_cond_id, end_date, info, history)
VALUES (default, $1,  $2,   $3,     $4,  $5,    $6,         $7,       $8,   $9) returning  cp_id`
	row := r.db.QueryRow(createTechQuery, cp.Date, cp.Cp_st, cp.Tz_id, cp.Proj, O_Id, pay_cond_id, cp.End_date, cp.Info, cp.History)
	if err = row.Scan(&CP_Id); err != nil {
		return 0, err
	}

	for _, v := range cp.Calendars {

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
		row := r.db.QueryRow(createTechQuery, name_id, v.Period, v.Term, nil, CP_Id)
		if err := row.Scan(&cal_id); err != nil {
			return 0, err
		}

	}

	for _, v := range cp.Costs {
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
		row := r.db.QueryRow(createTechQuery, metr_id, nil, nil, CP_Id, v.PPU, v.Info, task_id)
		if err := row.Scan(&cost_id); err != nil {
			return 0, err
		}
	}

	return CP_Id, nil
}

func (r *CPPostgres) GetById(O_Id, cp_id int) (packom.CP, error) {
	var costs []packom.Cost
	var tz_costs []packom.Cost
	var calendars []packom.Calendar
	var tz_calendars []packom.Calendar
	var cp packom.CP

	query := `SELECT cp_id, date, tz_id, proj, o_id, public."Pay_conds".name as pay_cond, end_date, info, history, cp_st
	FROM public."CP"
	join public."Pay_conds" on public."Pay_conds".pay_cond_id=public."CP".pay_cond_id
	where o_id = $1 and cp_id = $2`

	err := r.db.Get(&cp, query, O_Id, cp_id)
	if err != nil {
		return cp, err
	}

	query = `SELECT cost_id, public."Metrics".name as metr, count, tz_id, 0 as cp_id, 0 as ppu, info, public."Tasks".name as task
	FROM public."Costs"
	join public."Metrics" on public."Metrics".metr_id = public."Costs".metr_id
	join public."Tasks" on public."Tasks".task_id = public."Costs".task_id
	where tz_id = $1`

	err = r.db.Select(&tz_costs, query, cp.Tz_id)
	if err != nil {
		return cp, err
	}

	query = `SELECT cal_id, public."Task_names".name as name, period, term, tz_id, 0 as cp_id
	FROM public."Calendar"
	join public."Task_names" on public."Task_names".name_id = public."Calendar".name_id
	where tz_id = $1`

	err = r.db.Select(&tz_calendars, query, cp.Tz_id)
	if err != nil {
		return cp, err
	}

	query = `SELECT cost_id, public."Metrics".name as metr, 0 as count, 0 as tz_id, cp_id, ppu, info, public."Tasks".name as task
	FROM public."Costs"
	join public."Metrics" on public."Metrics".metr_id = public."Costs".metr_id
	join public."Tasks" on public."Tasks".task_id = public."Costs".task_id
	where cp_id = $1`

	err = r.db.Select(&costs, query, cp_id)
	if err != nil {
		return cp, err
	}

	query = `SELECT cal_id, public."Task_names".name as name, period, term, 0 as tz_id,  cp_id
	FROM public."Calendar"
	join public."Task_names" on public."Task_names".name_id = public."Calendar".name_id
	where cp_id = $1`

	err = r.db.Select(&calendars, query, cp_id)
	if err != nil {
		return cp, err
	}

	cp.Calendars = calendars
	cp.Costs = costs
	cp.Tz_Calendars = tz_calendars
	cp.Tz_Costs = tz_costs

	return cp, err
}
