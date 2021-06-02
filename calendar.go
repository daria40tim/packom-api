package packom

type Calendar struct {
	Cal_id int    `json:"cal_id" db:"cal_id"`
	Name   string `json:"task_name" db:"name"`
	Period int    `json:"period" db:"period"`
	Term   int    `json:"term" db:"term"`
	Tz_id  int    `json:"tz_id" db:"tz_id"`
	Cp_id  int    `json:"cp_id" db:"cp_id"`
	Active bool   `json:"active" db:"active"`
}
