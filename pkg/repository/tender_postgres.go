package repository

import (
	"fmt"
	"strconv"
	"time"

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
	query := `SELECT tender_id, public."Techs".end_date as date, public."Techs".selected_cp, public."Techs".proj , public."Tenders".tz_id, public."Techs".tz_st, public."Techs".task_name as task,
	public."Pack_groups".name as group, public."Pack_types".name as type, public."Pack_kinds".name as kind, public."Techs".active
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

	now := time.Now()

	for i, v := range techs {
		date, err := time.Parse(time.RFC3339, v.Date)
		if err != nil {
			return nil, err
		}
		if !v.Active {
			techs[i].Tender_st = "Отменен"
		} else if v.Active && v.Selected_cp != 0 {
			techs[i].Tender_st = "Архив"
		} else if now.Sub(date) > 0 {
			techs[i].Tender_st = "Ожидает решения"
		} else {
			techs[i].Tender_st = "Сбор КП"
		}
	}

	return techs, err
}

func (r *TenderPostgres) GetById(id int) (packom.TenderById, error) {
	var res packom.TenderById

	query := `SELECT tender_id, public."Tenders".date, public."Techs".selected_cp,public."Techs".active, public."Tenders".tz_id,  public."Orgs".name as client, 
	public."Techs".task_name as task,  public."Techs".proj as proj
	   FROM public."Tenders"
	   join public."Techs" on public."Techs".tz_id = public."Tenders".tz_id
	   join public."Orgs" on public."Orgs".o_id = public."Techs".o_id
	   where tender_id = $1`

	err := r.db.Get(&res, query, id)
	if err != nil {
		return res, err
	}

	var cp []int
	query = `SELECT cp_id
	FROM public."CP"
	where tz_id = $1`
	err = r.db.Select(&cp, query, res.Tz_id)
	if err != nil {
		return res, err
	}

	if len(cp) == 0 {
		return res, nil
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
	query = `SELECT distinct public."Task_names".name
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

		query = `SELECT  cp_period
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

func (r *TenderPostgres) UpdateById(input packom.Tender) (int, error) {
	var tender_id int
	if input.Selected_cp != 0 {
		query := `UPDATE public."Tenders"
	SET selected_cp=$1
	WHERE tender_id=$2 returning  tender_id`
		row := r.db.QueryRow(query, input.Selected_cp, input.Tender_id)
		if err := row.Scan(&tender_id); err != nil {
			return 0, err
		}

		query = `UPDATE public."Techs"
		SET tz_st=3, selected_cp=$1
		WHERE tz_id = $2 returning tz_id`
		row = r.db.QueryRow(query, input.Selected_cp, input.Tender_id)
		if err := row.Scan(&tender_id); err != nil {
			return 0, err
		}

		return tender_id, nil
	} else {
		query := `UPDATE public."Techs"
	SET tz_st=4, active = false
	WHERE tz_id = $1 returning tz_id`
		row := r.db.QueryRow(query, input.Tz_id)
		if err := row.Scan(&tender_id); err != nil {
			return 0, err
		}
		return tender_id, nil
	}
}

func tender_filter(source []int, value int) bool {
	for _, v := range source {
		if v == value {
			return true
		}
	}
	return false
}

func (r *TenderPostgres) GetAllTendersFiltered(O_Id int, EDate, SDate string, Projs, TZ_Ids, Tender_STS []int) ([]packom.TenderAll, error) {
	var techs []packom.TenderAll
	var res []packom.TenderAll
	query := `SELECT tender_id, public."Techs".end_date as date, public."Techs".selected_cp, public."Techs".proj , public."Tenders".tz_id, public."Techs".tz_st, public."Techs".task_name as task,
	public."Pack_groups".name as group, public."Pack_types".name as type, public."Pack_kinds".name as kind, public."Techs".active
	   FROM public."Tenders"
	   join public."Techs" on public."Tenders".tz_id = public."Techs".tz_id
	   join public."Pack_groups" on public."Pack_groups".group_id=public."Techs".group_id 
	   join public."Pack_types" on public."Pack_types".type_id=public."Techs".type_id 
	   join public."Pack_kinds" on public."Pack_kinds".kind_id=public."Techs".kind_id 
	   where public."Techs".o_id = $1`

	if SDate != "" {
		query += ` and public."Techs".end_date>='` + SDate + `' `
	}
	if EDate != "" {
		query += ` and public."Techs".end_date<='` + EDate + `' `
	}

	if len(Projs) == 0 {
		query += ` and public."Tenders".tender_id in (select tender_id from public."Tenders") `
	} else {
		query += ` and public."Tenders".tender_id in ( ` + strconv.Itoa(Projs[0])
		for i := 1; i < len(Projs); i++ {
			query += `, ` + strconv.Itoa(Projs[i])
		}
		query += `) `
	}

	if len(TZ_Ids) == 0 {
		query += ` and public."Tenders".tz_id in (select tz_id from public."Tenders") `
	} else {
		query += ` and public."Tenders".tz_id in ( ` + strconv.Itoa(TZ_Ids[0])
		for i := 1; i < len(TZ_Ids); i++ {
			query += `, ` + strconv.Itoa(TZ_Ids[i])
		}
		query += `) `
	}

	err := r.db.Select(&techs, query, O_Id)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	for i, v := range techs {
		date, err := time.Parse(time.RFC3339, v.Date)
		if err != nil {
			return nil, err
		}
		if !v.Active {
			techs[i].Tender_st = "Отменен"
			techs[i].Tender_st_id = 4
		} else if v.Active && v.Selected_cp != 0 {
			techs[i].Tender_st = "Архив"
			techs[i].Tender_st_id = 3
		} else if now.Sub(date) > 0 {
			techs[i].Tender_st = "Ожидает решения"
			techs[i].Tender_st_id = 2
		} else {
			techs[i].Tender_st = "Сбор КП"
			techs[i].Tender_st_id = 1
		}
	}

	cp := []int{1, 2, 3, 4}
	var c []int

	if len(Tender_STS) == 0 {
		c = cp[:]
	} else {
		c = Tender_STS[:]
	}

	for _, v := range techs {
		if tender_filter(c, v.Tender_st_id) {
			res = append(res, v)
		}
	}

	return res, err
}

func (r *TenderPostgres) GetTenderFilterData() (packom.TenderFilterData, error) {
	var res packom.TenderFilterData
	var projs []packom.ProFilter
	var tz_ids []packom.TZIdsFilter

	query := `SELECT distinct proj as name, tender_id as id
	FROM public."Tenders" 
	join public."Techs" on public."Techs".tz_id = public."Tenders".tz_id
	order by proj`

	err := r.db.Select(&projs, query)
	if err != nil {
		return res, err
	}

	query = `SELECT distinct tz_id as id, tz_id as name
	FROM public."Tenders"
	order by tz_id`

	err = r.db.Select(&tz_ids, query)
	if err != nil {
		return res, err
	}

	res.Projs = projs
	res.Tz_ids = tz_ids

	return res, nil
}
