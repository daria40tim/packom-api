package repository

import (
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

func (r *TenderPostgres) GetAll(O_Id int /*, filter packom.TechFilter*/) ([]packom.TenderAll, error) {
	var techs []packom.TenderAll
	query := `SELECT public."Tenders".tender_id, public."Tenders".tz_id, public."Tenders".date, public."Techs".proj, selected_cp, public."Pack_groups".name as group, public."Pack_kinds".name as kind, public."Pack_types".name as type
	FROM public."Tenders"
	join public."Techs" on public."Techs".tz_id=public."Tenders".tz_id
	join public."Pack_groups" on public."Pack_groups".group_id=public."Techs".group_id
	join public."Pack_kinds" on public."Pack_kinds".kind_id=public."Techs".kind_id
	join public."Pack_types" on public."Pack_types".type_id=public."Techs".type_id
	
	where public."Techs".o_id =$1;`

	err := r.db.Select(&techs, query, O_Id)

	return techs, err
}
