package repository

import (
	"fmt"

	"github.com/daria40tim/packom"
	"github.com/jmoiron/sqlx"
)

type TenderPostgres struct {
	db *sqlx.DB
}

func NewTenderPostgres(db *sqlx.DB) *TenderPostgres {
	return &TenderPostgres{db: db}
}

func (r *TenderPostgres) Create(O_Id int, tender packom.Tender) (int, error) {
	var tender_id int
	query := `INSERT INTO public."Tenders"(
		tender_id, date, selected_cp, tz_id, history)
	VALUES (default, $1, $2,          $3,    $4) returning  tender_id`
	row := r.db.QueryRow(query, tender.Date, tender.Selected_cp, tender.Tz_id, tender.History)
	if err := row.Scan(&tender_id); err != nil {
		return 0, err
	}
	return tender_id, nil
}

func (r *TenderPostgres) GetAll(O_Id int) ([]packom.TenderAll, error) {
	var techs []packom.TenderAll
	query := `SELECT tender_id, public."Tenders".date, selected_cp, public."Techs".proj , public."Tenders".tz_id, public."Techs".task_name as task,
	public."Pack_groups".name as group, public."Pack_types".name as type, public."Pack_kinds".name as kind
	   FROM public."Tenders"
	   join public."Techs" on public."Tenders".tz_id = public."Techs".tz_id
	   join public."Pack_groups" on public."Pack_groups".group_id=public."Techs".group_id 
	   join public."Pack_types" on public."Pack_types".type_id=public."Techs".type_id 
	   join public."Pack_kinds" on public."Pack_kinds".kind_id=public."Techs".kind_id 
	   where public."Techs".o_id = $1`

	err := r.db.Select(&techs, query, O_Id)
	if err != nil {
		return nil, err
	}

	return techs, err
}

func (r *TenderPostgres) GetById(id int) (packom.TenderById, error) {
	var res packom.TenderById

	query := `SELECT tender_id, public."Tenders".date, selected_cp, public."Tenders".tz_id,  public."Orgs".name as client, 
	public."Techs".task_name as task,  public."Techs".proj as proj
	   FROM public."Tenders"
	   join public."Techs" on public."Techs".tz_id = public."Tenders".tz_id
	   join public."Orgs" on public."Orgs".o_id = public."Techs".o_id
	   where tender_id = $1`

	err := r.db.Get(&res, query, id)
	if err != nil {
		return res, err
	}

	var fcst []float32
	query = `with a as (select count, task_id from public."Costs" where tz_id = $1 and active)
	SELECT distinct sum(a.count*public."Costs".ppu) over (partition by public."CP".cp_id)
		FROM public."Costs" 
		join a on a.task_id = public."Costs".task_id 
		join public."CP" on public."CP".cp_id = public."Costs".cp_id 
		where public."CP".cp_id in (select cp_id from public."CP" where tz_id = $1 and active)
		order by sum`

	err = r.db.Select(&fcst, query, res.Tz_id)
	if err != nil {
		return res, err
	}

	res.Min_cp = fmt.Sprintf("%f", fcst[0])
	res.Max_cp = fmt.Sprintf("%f", fcst[len(fcst)-1])

	var term int
	query = `SELECT sum(period)
	FROM public."Calendar"
	where tz_id = $1 and active`

	err = r.db.Get(&term, query, res.Tz_id)
	if err != nil {
		return res, err
	}

	res.Term = term

	var tz_cal []string
	query = `SELECT public."Task_names".name
	FROM public."Calendar"
	join public."Task_names" on public."Task_names".name_id = public."Calendar".name_id
	where tz_id = $1 and active`

	err = r.db.Select(&tz_cal, query, res.Tz_id)
	if err != nil {
		return res, err
	}

	res.Tz_Calendars = tz_cal

	var tz_cst []packom.Cost
	query = `SELECT cost_id, public."Metrics".name as metr, count, tz_id, public."Tasks".name as task, 0 as cp_id
	FROM public."Costs" 
	join public."Tasks" on public."Tasks".task_id = public."Costs".task_id
	join public."Metrics" on public."Metrics".metr_id = public."Costs".metr_id
	where tz_id = $1 and active`

	err = r.db.Select(&tz_cst, query, res.Tz_id)
	if err != nil {
		return res, err
	}

	res.Tz_Costs = tz_cst

	var cps []packom.TenderCP
	query = `SELECT cp_id, public."CP".o_id,public."Pay_conds".name as pay_cond, public."Orgs".name as org
	FROM public."CP"
	join public."Orgs" on public."Orgs".o_id = public."CP".o_id
	join public."Pay_conds" on public."Pay_conds".pay_cond_id = public."CP".pay_cond_id
	where tz_id =$1`

	err = r.db.Select(&cps, query, res.Tz_id)
	if err != nil {
		return res, err
	}

	for i, v := range cps {
		var costs []packom.TenderCost
		var cals []int
		var sum int

		query = `with a as (select count, task_id from  public."Costs"  where active and tz_id = $1)
		SELECT sum(a.count * ppu)
			FROM public."Costs" 
			join a on a.task_id =public."Costs".task_id  
			where cp_id = $2 and active`

		err = r.db.Get(&sum, query, res.Tz_id, v.Cp_id)
		if err != nil {
			return res, err
		}

		cps[i].Sum = float32(sum)

		query = `with a as (select count, task_id from  public."Costs" where tz_id = $1 and active)
		SELECT  public."Tasks".name as task, a.count*ppu as cost_sum
			FROM public."Costs"
			join public."Tasks" on public."Tasks".task_id =  public."Costs".task_id 
			join a on public."Costs".task_id = a.task_id
			where cp_id = $2 and active`

		err = r.db.Select(&costs, query, res.Tz_id, v.Cp_id)
		if err != nil {
			return res, err
		}

		cps[i].Costs = costs

		query = `SELECT  period
		FROM public."Calendar"
		where cp_id = $1 and active`

		err = r.db.Select(&cals, query, v.Cp_id)
		if err != nil {
			return res, err
		}
		cps[i].Calendars = cals
	}

	res.CPs = cps

	return res, nil
}

func (r *TenderPostgres) GetFullCosts(id int) ([]packom.FullCost, error) {
	var res []packom.FullCost

	query := `with a as (select count, task_id from public."Costs" where tz_id = $1 and active)
	SELECT distinct sum(a.count*public."Costs".ppu) over (partition by public."CP".cp_id), public."Costs".cp_id, public."CP".o_id
		FROM public."Costs" 
		join a on a.task_id = public."Costs".task_id 
		join public."CP" on public."CP".cp_id = public."Costs".cp_id 
		where public."CP".cp_id in (select cp_id from public."CP" where tz_id = $1 and active)`

	err := r.db.Select(&res, query, id)
	if err != nil {
		return res, err
	}

	return res, nil

}
