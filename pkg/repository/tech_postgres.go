package repository

import (
	"database/sql"
	"os"
	"strconv"
	"time"

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
	(tz_id, date, o_id, end_date, proj, group_id, kind_id, type_id, tz_st, tender_st, cp_st, pay_cond_id, private, info, history, task_name)
	VALUES (default, $1, $2, $3, $4,   $5,       $6,      $7,       $8,    $9,        $10,   $11,         $12,     $13,  $14, $15) returning  tz_id`
	row := r.db.QueryRow(createTechQuery, tech.Date, O_Id, tech.End_date, tech.Proj, group_id, kind_id, type_id, 0, tech.Tender_st, tech.Cp_st, pay_cond_id, tech.Private, tech.Info, tech.History, tech.Task)
	if err := row.Scan(&Tz_Id); err != nil {
		return 0, err
	}

	err = os.MkdirAll("techs/"+strconv.Itoa(Tz_Id), 0777)
	if err != nil {
		return 0, err
	}

	var tender_id int
	query = `INSERT INTO public."Tenders"(
		tender_id, date, selected_cp, tz_id, history)
VALUES (default,   $1,   $2,          $3,    $4) returning tender_id`

	row = r.db.QueryRow(query, tech.End_date, 0, Tz_Id, "")
	if err := row.Scan(&tender_id); err != nil {
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
				cal_id, name_id, period, term, tz_id, cp_id, active)
		VALUES (default,$1,      $2,      $3,   $4,    $5, $6) returning  cal_id`
		row := r.db.QueryRow(createTechQuery, name_id, v.Period, v.Term, Tz_Id, nil, true)
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
			cost_id, metr_id, count, tz_id, cp_id, ppu, info, task_id, active)
	VALUES (default, $1,      $2,     $3,   $4,    $5,   $6,  $7, $8) returning  cost_id`
		row := r.db.QueryRow(createTechQuery, metr_id, v.Count, Tz_Id, nil, nil, v.Info, task_id, true)
		if err := row.Scan(&cost_id); err != nil {
			return 0, err
		}
	}

	for _, v := range tech.Docs {

		var name_id int

		query = `INSERT INTO public."Tech_docs"(
			tz_id, file_name, active)
			VALUES ($1, $2, true) returning  tz_id`
		row := r.db.QueryRow(query, Tz_Id, v)
		if err := row.Scan(&name_id); err != nil {
			return 0, err
		}

	}

	return Tz_Id, nil
}

func contains(s []packom.CP_srv, tz_id int) int {
	for _, a := range s {
		if a.Tz_id == tz_id {
			return a.Cp_id
		}
	}
	return 0
}

func (r *TechPostgres) GetAll(O_Id int) ([]packom.TechAll, error) {
	var group_id int
	var techs []packom.TechAll

	query := `SELECT group_id
	FROM public."Orgs" where o_id = $1`

	err := r.db.Get(&group_id, query, O_Id)
	if err != nil {
		return techs, err
	}

	var cps []packom.CP_srv

	query = `SELECT cp_id, date, tz_id, o_id, end_date
	FROM public."CP" where o_id=$1`

	err = r.db.Select(&cps, query, O_Id)
	if err != nil {
		return techs, err
	}

	if group_id == 1 {
		query = `SELECT distinct public."Techs".date,public."Techs".selected_cp,  public."Techs".task_name as task,public."Techs".end_date, public."Orgs".name as client, public."Techs".o_id, public."Techs".tz_id,  public."Techs".end_date, 
	public."Techs".proj, public."Pack_groups".name as group, public."Pack_types".name as type, public."Pack_kinds".name as kind, public."Techs".active, 
	count(cp_id) over (partition by public."Techs".tz_id) as cp_count
	FROM public."Techs" join public."Orgs" on public."Techs".o_id=public."Orgs".o_id 
	join public."Pack_groups" on public."Pack_groups".group_id=public."Techs".group_id 
	join public."Pack_types" on public."Pack_types".type_id=public."Techs".type_id 
	join public."Pack_kinds" on public."Pack_kinds".kind_id=public."Techs".kind_id 
	left join public."CP" on public."CP".tz_id=public."Techs".tz_id 
	where public."Techs".o_id=$1`

		err = r.db.Select(&techs, query, O_Id)

	} else {

		query = `SELECT distinct public."Techs".date,public."Techs".selected_cp,  public."Techs".task_name as task,public."Techs".end_date, public."Orgs".name as client, public."Techs".o_id, public."Techs".tz_id,  public."Techs".end_date, 
	public."Techs".proj, public."Pack_groups".name as group, public."Pack_types".name as type, public."Pack_kinds".name as kind, public."Techs".active, 
	count(cp_id) over (partition by public."Techs".tz_id) as cp_count
	FROM public."Techs" join public."Orgs" on public."Techs".o_id=public."Orgs".o_id 
	join public."Pack_groups" on public."Pack_groups".group_id=public."Techs".group_id 
	join public."Pack_types" on public."Pack_types".type_id=public."Techs".type_id 
	join public."Pack_kinds" on public."Pack_kinds".kind_id=public."Techs".kind_id 
	left join public."CP" on public."CP".tz_id=public."Techs".tz_id 
	where not private or $1 in (select distinct f_o_id from public."Orgs_orgs" where public."Orgs_orgs".o_id = public."Techs".o_id)`

		err = r.db.Select(&techs, query, O_Id)

	}

	now := time.Now()

	for i, v := range techs {
		date, _ := time.Parse(time.RFC3339, v.End_date)
		if now.Sub(date) > 0 {
			techs[i].Tz_st = "Архив"
		} else {
			techs[i].Tz_st = "Активно"
		}

		if group_id == 1 {
			techs[i].CP_st = "-"
		} else if contains(cps, v.Tz_id) == 0 && v.Selected_cp == 0 {
			techs[i].CP_st = "Не подано"
		} else if contains(cps, v.Tz_id) == v.Selected_cp {
			techs[i].CP_st = "Принято"
		} else if contains(cps, v.Tz_id) != v.Selected_cp && v.Selected_cp != 0 {
			techs[i].CP_st = "Отклонено"
		} else {
			techs[i].CP_st = "Подано"
		}

		if !v.Active {
			techs[i].Tender_st = "Отменен"
		} else if v.Selected_cp != 0 {
			techs[i].Tender_st = "Принят"
		} else if now.Sub(date) > 0 {
			techs[i].Tender_st = "Ожидает решения"
		} else {
			techs[i].Tender_st = "Сбор КП"
		}

	}

	return techs, err
}

func (r *TechPostgres) GetById(O_Id, tz_id int) (packom.Tech, []packom.Cost, []packom.Calendar, error) {
	var costs []packom.Cost
	var tech packom.Tech
	var calendars []packom.Calendar

	query := `SELECT cost_id, public."Metrics".name as metr, count, tz_id, 0 as cp_id, 0 as ppu, info, public."Tasks".name as task
	FROM public."Costs"
	join public."Metrics" on public."Metrics".metr_id = public."Costs".metr_id
	join public."Tasks" on public."Tasks".task_id = public."Costs".task_id
	where tz_id = $1 and active`

	err := r.db.Select(&costs, query, tz_id)
	if err != nil {
		return tech, nil, nil, err
	}

	query = `SELECT cal_id, public."Task_names".name as name, period, term, tz_id, 0 as cp_id
	FROM public."Calendar"
	join public."Task_names" on public."Task_names".name_id = public."Calendar".name_id
	where tz_id = $1 and active`

	err = r.db.Select(&calendars, query, tz_id)
	if err != nil {
		return tech, nil, nil, err
	}

	var docs []string
	var e_docs []string

	query = `SELECT file_name
	FROM public."Tech_docs"
	where tz_id = $1;`

	err = r.db.Select(&docs, query, tz_id)
	if err != nil {
		return tech, nil, nil, err
	}

	if docs == nil {
		tech.Docs = e_docs
	} else {
		tech.Docs = docs
	}

	query = `SELECT public."Techs".date, public."Techs".o_id, public."Pack_groups".name as group, public."Pack_kinds".name as kind, task_name,
	public."Pack_types".name as type, public."Orgs".name as client, public."Techs".end_date, public."Techs".proj, public."Pay_conds".name as pay_cond,
	tender_st, cp_st, private, public."Techs".info, public."Techs".history, public."Techs".tz_st, public."Techs".tz_id
		FROM public."Techs"
		join public."Orgs" on public."Orgs".o_id = public."Techs".o_id
		join public."Pack_groups" on public."Pack_groups".group_id = public."Techs".group_id
		join public."Pack_kinds" on public."Pack_kinds".kind_id = public."Techs".kind_id
		join public."Pack_types" on public."Pack_types".type_id = public."Techs".type_id
		join public."Pay_conds" on public."Pay_conds".pay_cond_id = public."Techs".pay_cond_id
		where tz_id = $1`

	err = r.db.Get(&tech, query, tz_id)

	return tech, costs, calendars, err
}

func (r *TechPostgres) SelectAll() (packom.Select, error) {
	var metrics, groups, kinds, types, pay_conds, task_names, tasks, task_kinds []string
	var res packom.Select

	query := `SELECT name FROM public."Metrics"`

	err := r.db.Select(&metrics, query)
	if err != nil {
		return res, err
	}

	query = `SELECT name FROM public."Pack_groups"`

	err = r.db.Select(&groups, query)
	if err != nil {
		return res, err
	}

	query = `SELECT name FROM public."Pack_kinds"`

	err = r.db.Select(&kinds, query)
	if err != nil {
		return res, err
	}

	query = `SELECT name FROM public."Pack_types"`

	err = r.db.Select(&types, query)
	if err != nil {
		return res, err
	}

	query = `SELECT name FROM public."Pay_conds"`

	err = r.db.Select(&pay_conds, query)
	if err != nil {
		return res, err
	}

	query = `SELECT name FROM public."Task_names"`

	err = r.db.Select(&task_names, query)
	if err != nil {
		return res, err
	}

	query = `SELECT name FROM public."Tasks"`

	err = r.db.Select(&tasks, query)
	if err != nil {
		return res, err
	}

	query = `SELECT name FROM public."Task_kinds"`

	err = r.db.Select(&task_kinds, query)
	if err != nil {
		return res, err
	}

	res.Groups = groups
	res.Kinds = kinds
	res.Metrics = metrics
	res.Pay_conds = pay_conds
	res.Task_names = task_names
	res.Tasks = tasks
	res.Types = types
	res.Task_kinds = task_kinds

	return res, nil
}

func (r *TechPostgres) DeleteCost(tz_id int, task, history string) (int, error) {
	var task_id int
	var res int

	query := `select task_id from public."Tasks" where name=$1`
	err := r.db.Get(&task_id, query, task)
	if err != nil {
		return 0, err
	}

	query = `UPDATE public."Techs"
	SET history=$1
	WHERE tz_id=$2 returning tz_id`

	row := r.db.QueryRow(query, history, tz_id)
	if err = row.Scan(&res); err != nil {
		return 0, err
	}

	query = `UPDATE public."Costs"
	SET active=false
	WHERE tz_id = $1 and task_id = $2 returning cost_id`

	row = r.db.QueryRow(query, tz_id, task_id)
	if err = row.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (r *TechPostgres) DeleteCal(tz_id int, name, history string) (int, error) {
	var name_id int
	var res int

	query := `select name_id from public."Task_names" where name=$1`
	err := r.db.Get(&name_id, query, name)
	if err != nil {
		return 0, err
	}

	query = `UPDATE public."Calendar"
	SET active=false
	WHERE tz_id = $1 and name_id = $2 returning cal_id`

	row := r.db.QueryRow(query, tz_id, name_id)
	if err = row.Scan(&res); err != nil {
		return 0, err
	}

	query = `UPDATE public."Techs"
	SET history=$1
	WHERE tz_id=$2 returning tz_id`

	row = r.db.QueryRow(query, history, tz_id)
	if err = row.Scan(&res); err != nil {
		return 0, err
	}
	return res, nil
}

func (r *TechPostgres) UpdateById(id int, tech packom.Tech) (int, error) {
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
	query = `DELETE FROM public."Calendar"
	WHERE tz_id = $1 returning tz_id`
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&Tz_Id); err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	query = `DELETE FROM public."Costs"
	WHERE tz_id = $1 returning tz_id`
	row = r.db.QueryRow(query, id)
	if err := row.Scan(&Tz_Id); err != nil && err != sql.ErrNoRows {
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
			cal_id, name_id, period, term, tz_id, cp_id, active)
			VALUES (default, $1, $2, $3, $4, $5, $6) returning  cal_id`
		row := r.db.QueryRow(createTechQuery, name_id, v.Period, v.Term, id, nil, v.Active)
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
			cost_id, metr_id, count, tz_id, cp_id, ppu, info, task_id, active)
	VALUES (default, $1,      $2,     $3,   $4,    $5,   $6,  $7, $8) returning  cost_id`
		row := r.db.QueryRow(createTechQuery, metr_id, v.Count, id, nil, nil, v.Info, task_id, v.Active)
		if err := row.Scan(&cost_id); err != nil {
			return 0, err
		}
	}

	updateTechQuery := `UPDATE public."Techs"
	SET proj=$1, group_id=$2, kind_id=$3, type_id=$4, pay_cond_id=$5, info=$6, history=$7, task_name=$8, end_date = $9
	WHERE tz_id=$10 returning tz_id`
	row = r.db.QueryRow(updateTechQuery, tech.Proj, group_id, kind_id, type_id, pay_cond_id, tech.Info, tech.History, tech.Task, tech.End_date, id)
	if err := row.Scan(&Tz_Id); err != nil {
		return 0, err
	}
	return 0, nil
}

func (r *TechPostgres) AddTechDoc(name string, o_id, tz_id int) error {
	var id int

	query := `INSERT INTO public."Tech_docs"(
		tz_id, file_name)
		VALUES ($1, $2) returning tz_id`

	row := r.db.QueryRow(query, tz_id, name)
	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}

func (r *TechPostgres) GetFilterData() (packom.TechFilterData, error) {
	var res packom.TechFilterData
	var clients []packom.ClientFilter
	var projs []packom.ProjFilter

	query := `SELECT distinct proj as name, tz_id as id
	FROM public."Techs" order by proj`

	err := r.db.Select(&projs, query)
	if err != nil {
		return res, err
	}

	query = `SELECT distinct public."Orgs".name, public."Orgs".o_id as id
	FROM public."Techs"
	join public."Orgs" on public."Orgs".o_id = public."Techs".o_id
	order by name`

	err = r.db.Select(&clients, query)
	if err != nil {
		return res, err
	}

	res.Clients = clients
	res.Projs = projs

	return res, nil
}

func filter(source []int, value int) bool {
	for _, v := range source {
		if v == value {
			return true
		}
	}
	return false
}

func (r *TechPostgres) GetAllTechsFiltered(O_Id int, EDate, SDate string, Clients, Projs, TZ_STS, CP_STS, Tender_STS []int) ([]packom.TechAll, error) {
	var group_id int
	var techs []packom.TechAll
	var res []packom.TechAll

	query := `SELECT group_id
	FROM public."Orgs" where o_id = $1`

	err := r.db.Get(&group_id, query, O_Id)
	if err != nil {
		return techs, err
	}

	var cps []packom.CP_srv

	query = `SELECT cp_id, date, tz_id, o_id, end_date
	FROM public."CP" where o_id=$1`

	err = r.db.Select(&cps, query, O_Id)
	if err != nil {
		return techs, err
	}

	if group_id == 1 {
		query = `SELECT distinct public."Techs".date,public."Techs".selected_cp,  public."Techs".task_name as task,public."Techs".end_date, public."Orgs".name as client, public."Techs".o_id, public."Techs".tz_id,  public."Techs".end_date, 
	public."Techs".proj, public."Pack_groups".name as group, public."Pack_types".name as type, public."Pack_kinds".name as kind, public."Techs".active, 
	count(cp_id) over (partition by public."Techs".tz_id) as cp_count
	FROM public."Techs" join public."Orgs" on public."Techs".o_id=public."Orgs".o_id 
	join public."Pack_groups" on public."Pack_groups".group_id=public."Techs".group_id 
	join public."Pack_types" on public."Pack_types".type_id=public."Techs".type_id 
	join public."Pack_kinds" on public."Pack_kinds".kind_id=public."Techs".kind_id 
	left join public."CP" on public."CP".tz_id=public."Techs".tz_id 
	where public."Techs".o_id=$1`

		if SDate != "" {
			query += ` and public."Techs".date>='` + SDate + `' `
		}
		if EDate != "" {
			query += ` and public."Techs".date<='` + EDate + `' `
		}

		if len(Clients) == 0 {
			query += ` and public."Techs".o_id in (select o_id from public."Orgs") `
		} else {
			query += ` and public."Techs".o_id in ( ` + strconv.Itoa(Clients[0])
			for i := 1; i < len(Clients); i++ {
				query += `, ` + strconv.Itoa(Clients[i])
			}
			query += `) `
		}

		if len(Projs) == 0 {
			query += ` and public."Techs".tz_id in (select tz_id from public."Techs") `
		} else {
			query += ` and public."Techs".tz_id in ( ` + strconv.Itoa(Projs[0])
			for i := 1; i < len(Projs); i++ {
				query += `, ` + strconv.Itoa(Projs[i])
			}
			query += `) `
		}

		err = r.db.Select(&techs, query, O_Id)

	} else {

		query = `SELECT distinct public."Techs".date,public."Techs".selected_cp,  public."Techs".task_name as task,public."Techs".end_date, public."Orgs".name as client, public."Techs".o_id, public."Techs".tz_id,  public."Techs".end_date, 
	public."Techs".proj, public."Pack_groups".name as group, public."Pack_types".name as type, public."Pack_kinds".name as kind, public."Techs".active, 
	count(cp_id) over (partition by public."Techs".tz_id) as cp_count
	FROM public."Techs" join public."Orgs" on public."Techs".o_id=public."Orgs".o_id 
	join public."Pack_groups" on public."Pack_groups".group_id=public."Techs".group_id 
	join public."Pack_types" on public."Pack_types".type_id=public."Techs".type_id 
	join public."Pack_kinds" on public."Pack_kinds".kind_id=public."Techs".kind_id 
	left join public."CP" on public."CP".tz_id=public."Techs".tz_id 
	where (not private or $1 in (select distinct f_o_id from public."Orgs_orgs" where public."Orgs_orgs".o_id = public."Techs".o_id)) `

		if SDate != "" {
			query += ` and public."Techs".date>='` + SDate + `' `
		}
		if EDate != "" {
			query += ` and public."Techs".date<='` + EDate + `' `
		}

		if len(Clients) == 0 {
			query += ` and public."Techs".o_id in (select o_id from public."Orgs") `
		} else {
			query += ` and public."Techs".o_id in ( ` + strconv.Itoa(Clients[0])
			for i := 1; i < len(Clients); i++ {
				query += `, ` + strconv.Itoa(Clients[i])
			}
			query += `) `
		}

		if len(Projs) == 0 {
			query += ` and public."Techs".tz_id in (select tz_id from public."Techs") `
		} else {
			query += ` and public."Techs".tz_id in ( ` + strconv.Itoa(Projs[0])
			for i := 1; i < len(Projs); i++ {
				query += `, ` + strconv.Itoa(Projs[i])
			}
			query += `) `
		}

		err = r.db.Select(&techs, query, O_Id)

	}

	now := time.Now()

	for i, v := range techs {
		date, _ := time.Parse(time.RFC3339, v.End_date)
		if now.Sub(date) > 0 {
			techs[i].Tz_st = "Архив"
			techs[i].Tz_st_id = 2
		} else {
			techs[i].Tz_st = "Активно"
			techs[i].Tz_st_id = 1
		}

		if group_id == 1 {
			techs[i].CP_st = "-"
			techs[i].CP_st_id = 5
		} else if contains(cps, v.Tz_id) == 0 && v.Selected_cp == 0 {
			techs[i].CP_st = "Не подано"
			techs[i].CP_st_id = 1
		} else if contains(cps, v.Tz_id) == v.Selected_cp {
			techs[i].CP_st = "Принято"
			techs[i].CP_st_id = 4
		} else if contains(cps, v.Tz_id) != v.Selected_cp && v.Selected_cp != 0 {
			techs[i].CP_st = "Отклонено"
			techs[i].CP_st_id = 3
		} else {
			techs[i].CP_st = "Подано"
			techs[i].CP_st_id = 2
		}

		if !v.Active {
			techs[i].Tender_st = "Отменен"
			techs[i].Tender_st_id = 4
		} else if v.Selected_cp != 0 {
			techs[i].Tender_st = "Принят"
			techs[i].Tender_st_id = 3
		} else if now.Sub(date) > 0 {
			techs[i].Tender_st = "Ожидает решения"
			techs[i].Tender_st_id = 2
		} else {
			techs[i].Tender_st = "Сбор КП"
			techs[i].Tender_st_id = 1
		}
	}

	tz := []int{1, 2}
	cp := []int{1, 2, 3, 4, 5}
	tender := []int{1, 2, 3, 4}
	var t []int
	var ten []int
	var c []int

	if len(TZ_STS) == 0 {
		t = tz[:]
	} else {
		t = TZ_STS[:]
	}

	if len(Tender_STS) == 0 {
		ten = tender[:]
	} else {
		ten = Tender_STS[:]
	}

	if len(CP_STS) == 0 {
		c = cp[:]
	} else {
		c = CP_STS[:]
	}

	for _, v := range techs {
		if filter(t, v.Tz_st_id) && filter(ten, v.Tender_st_id) && filter(c, v.CP_st_id) {
			res = append(res, v)
		}
	}

	return res, err
}
