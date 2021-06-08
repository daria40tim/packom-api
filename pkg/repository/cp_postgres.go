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

	query := `SELECT cp_id, public."CP".date, case public."CP".cp_st when 1 then 'Обрабатывается' when 2 then 'Принято' when 3 then 'Отклонено' end as cp_st, 
	public."CP".tz_id, public."CP".proj, public."CP".o_id, public."Orgs".name as client, public."Pack_groups".name as group, 
	public."Pack_types".name as type, public."Pack_kinds".name as kind, public."Techs".task_name as task_name  
	FROM public."CP" join public."Orgs" on public."Orgs".o_id=public."CP".o_id 
	join public."Techs" on public."Techs".tz_id=public."CP".tz_id
	join public."Pack_groups" on public."Pack_groups".group_id=public."Techs".group_id
	join public."Pack_kinds" on public."Pack_kinds".kind_id=public."Techs".kind_id
	join public."Pack_types" on public."Pack_types".type_id=public."Techs".type_id
	where public."CP".o_id =$1;`

	err := r.db.Select(&techs, query, O_Id)

	return techs, err
}

func (r *CPPostgres) Create(O_Id int, cp packom.CPIns) (int, error) {
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
	row := r.db.QueryRow(createTechQuery, cp.Date, 1, cp.Tz_id, cp.Proj, O_Id, pay_cond_id, cp.End_date, cp.Info, cp.History)
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
				cal_id, name_id, period, term, tz_id, cp_id, active)
		VALUES (default,$1,      $2,      $3,   $4,    $5, $6) returning  cal_id`
		row := r.db.QueryRow(createTechQuery, name_id, v.Period, v.Term, nil, CP_Id, true)
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

		var cost_id int
		createTechQuery := `INSERT INTO public."Costs"(
			cost_id, metr_id, count, tz_id, cp_id, ppu, info, task_id, active)
	VALUES (default, $1,      $2,     $3,   $4,    $5,   $6,  $7, true) returning  cost_id`
		row := r.db.QueryRow(createTechQuery, nil, nil, nil, CP_Id, v.PPU, v.Info, task_id)
		if err := row.Scan(&cost_id); err != nil {
			return 0, err
		}
	}

	for _, v := range cp.Docs {

		var name_id int

		query = `INSERT INTO public."CP_docs"(
			file_name, cp_id, active)
			VALUES ($1, $2, true) returning  cp_id`
		row := r.db.QueryRow(query, v, CP_Id)
		if err := row.Scan(&name_id); err != nil {
			return 0, err
		}
	}

	return CP_Id, nil
}

func (r *CPPostgres) GetById(O_Id, cp_id int) (packom.CPId, error) {
	var costs []packom.Cost
	var tz_costs []packom.Cost
	var calendars []packom.CPCalendar
	var tz_calendars []packom.Calendar
	var tz_docs []string
	var docs []string
	var e_docs []string
	var cp packom.CPId

	query := `SELECT public."CP".cp_id, public."CP".date, public."CP".tz_id, public."Techs".proj, public."CP".o_id, public."Pay_conds".name as pay_cond, 
	public."Pack_groups".name as group, public."Pack_kinds".name as kind, public."CP".end_date, public."CP".info, public."CP".history, 
	public."Pack_types".name as type, public."Orgs".name as client, public."Techs".date as tz_date, public."Techs".end_date as tz_end_date, 
	case public."Techs".private when true then 'Закрыт' else 'Открыт' end as private, public."Techs".o_id as tz_o_id,  public."Techs".task_name as task_name,  
	public."Techs".info as tz_info
		FROM public."CP"
		join public."Pay_conds" on public."Pay_conds".pay_cond_id=public."CP".pay_cond_id
		join public."Techs" on public."Techs".tz_id = public."CP".tz_id
		join public."Orgs" on public."Techs".o_id = public."Orgs".o_id
		join public."Pack_groups" on public."Pack_groups".group_id = public."Techs".group_id
		join public."Pack_kinds" on public."Pack_kinds".kind_id = public."Techs".kind_id
		join public."Pack_types" on public."Pack_types".type_id = public."Techs".type_id
	where cp_id = $1`

	err := r.db.Get(&cp, query, cp_id)
	if err != nil {
		return cp, err
	}

	query = `SELECT public."Pay_conds".name as tz_pay_cond
	FROM public."CP"
	join public."Techs" on public."Techs".tz_id = public."CP".tz_id
	join public."Pay_conds" on public."Pay_conds".pay_cond_id=public."Techs".pay_cond_id
	where cp_id = $1`

	err = r.db.Get(&cp.Tz_Pay_cond, query, cp_id)
	if err != nil {
		return cp, err
	}

	query = `SELECT public."Orgs".name as org
	FROM public."CP"
	join public."Orgs" on public."Orgs".o_id = public."CP".o_id
	where cp_id = $1`

	err = r.db.Get(&cp.Org, query, cp_id)
	if err != nil {
		return cp, err
	}

	query = `SELECT cost_id, public."Metrics".name as metr, count, tz_id, 0 as cp_id, 0 as ppu, info, public."Tasks".name as task, active
	FROM public."Costs"
	join public."Metrics" on public."Metrics".metr_id = public."Costs".metr_id
	join public."Tasks" on public."Tasks".task_id = public."Costs".task_id
	where tz_id = $1 and active`

	err = r.db.Select(&tz_costs, query, cp.Tz_id)
	if err != nil {
		return cp, err
	}

	query = `SELECT cal_id, public."Task_names".name as name, period, term, tz_id, 0 as cp_id, active
	FROM public."Calendar"
	join public."Task_names" on public."Task_names".name_id = public."Calendar".name_id
	where tz_id = $1 and active`

	err = r.db.Select(&tz_calendars, query, cp.Tz_id)
	if err != nil {
		return cp, err
	}

	query = `with a as (SELECT cost_id, metr_id, count, tz_id, cp_id, ppu, info, task_id, sum, active
		FROM public."Costs" where tz_id = $1 and active)
	SELECT a.count, a.tz_id, public."Costs".cp_id, public."Costs".ppu, public."Costs".info, public."Tasks".name as task, a.active, 
	public."Metrics".name as metr
	FROM public."Costs"
		join a on a.task_id = public."Costs".task_id
		join public."Tasks" on a.task_id=public."Tasks".task_id
		join public."Metrics" on a.metr_id=public."Metrics".metr_id
		where public."Costs".cp_id = $2 and public."Costs".active`

	err = r.db.Select(&costs, query, cp.Tz_id, cp_id)
	if err != nil {
		return cp, err
	}

	query = `with a as (SELECT cal_id, name_id, period, term, tz_id, cp_id, active
		FROM public."Calendar" where tz_id = $1 and active)
		
	SELECT public."Calendar".cal_id, public."Calendar".period, public."Calendar".term, a.tz_id , a.period as tz_period, a.active,
	a.term as tz_term, public."Calendar".cp_id, public."Task_names".name as name
		FROM public."Calendar" 
		right join a on public."Calendar".name_id =  a.name_id
		join public."Task_names" on public."Calendar".name_id = public."Task_names".name_id
		where public."Calendar".cp_id = $2 and public."Calendar".active`

	err = r.db.Select(&calendars, query, cp.Tz_id, cp_id)
	if err != nil {
		return cp, err
	}

	query = `SELECT file_name
	FROM public."Tech_docs"
	where tz_id = $1;`

	err = r.db.Select(&tz_docs, query, cp.Tz_id)
	if err != nil {
		return cp, err
	}

	if docs == nil {
		cp.Tz_Docs = e_docs
	} else {
		cp.Tz_Docs = tz_docs
	}

	query = `SELECT file_name
	FROM public."CP_docs"
	where cp_id = $1;`

	err = r.db.Select(&docs, query, cp.Cp_id)
	if err != nil {
		return cp, err
	}

	if docs == nil {
		cp.Docs = e_docs
	} else {
		cp.Docs = docs
	}

	cp.Calendars = calendars
	cp.Costs = costs
	cp.Tz_Calendars = tz_calendars
	cp.Tz_Costs = tz_costs

	return cp, err
}

func (r *CPPostgres) UpdateById(cp_id int, cp packom.CPIns) (int, error) {
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
	createTechQuery := `UPDATE public."CP"
	SET pay_cond_id=$1, end_date=$2, info=$3, history=$4
	WHERE cp_id = $5 returning  cp_id`
	row := r.db.QueryRow(createTechQuery, pay_cond_id, cp.End_date, cp.Info, cp.History, cp_id)
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
				cal_id, name_id, period, term, tz_id, cp_id, active)
		VALUES (default,$1,      $2,      $3,   $4,    $5, $6) returning  cal_id`
		row := r.db.QueryRow(createTechQuery, name_id, v.Period, v.Term, nil, CP_Id, true)
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

		var cost_id int
		createTechQuery := `INSERT INTO public."Costs"(
			cost_id, metr_id, count, tz_id, cp_id, ppu, info, task_id, active)
	VALUES (default, $1,      $2,     $3,   $4,    $5,   $6,  $7, true) returning  cost_id`
		row := r.db.QueryRow(createTechQuery, nil, nil, nil, CP_Id, v.PPU, v.Info, task_id)
		if err := row.Scan(&cost_id); err != nil {
			return 0, err
		}
	}

	for _, v := range cp.Docs {

		var name_id int

		query = `INSERT INTO public."CP_docs"(
			file_name, cp_id, active)
			VALUES ($1, $2, true) returning  cp_id`
		row := r.db.QueryRow(query, v, CP_Id)
		if err := row.Scan(&name_id); err != nil {
			return 0, err
		}
	}

	return CP_Id, nil
}

func (r *CPPostgres) DeleteCal(id int) (int, error) {
	query := `UPDATE public."Calendar"
	SET active=false
	WHERE cp_id = $1`

	row := r.db.QueryRow(query, id)
	if err := row.Scan(); err != nil {
		return 0, err
	}

	return 0, nil
}

func (r *CPPostgres) DeleteCst(id int) (int, error) {
	query := `UPDATE public."Costs"
	SET active=false
	WHERE cp_id = $1`

	row := r.db.QueryRow(query, id)
	if err := row.Scan(); err != nil {
		return 0, err
	}

	return 0, nil
}
