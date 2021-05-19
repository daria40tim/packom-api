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
