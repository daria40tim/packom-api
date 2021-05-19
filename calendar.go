package packom

type Calendar struct {
	Cal_id int    `json:"cal_id"`
	Name   string `json:"name"`
	Period int    `json:"period"`
	Term   int    `json:"term"`
	Tz_id  int    `json:"tz_id"`
	Cp_id  int    `json:"cp_id"`
}
