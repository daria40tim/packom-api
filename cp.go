package packom

type CP struct {
	Cp_id     int        `json:"cp_id"`
	Date      string     `json:"date"`
	Cp_st     int        `json:"cp_st"`
	Tz_id     int        `json:"tz_id"`
	Proj      string     `json:"proj"`
	O_id      string     `json:"o_id"`
	Pay_cond  string     `json:"pay_cond"`
	End_date  string     `json:"end_date"`
	Info      string     `json:"info"`
	History   string     `json:"history"`
	Costs     []Cost     `json:"costs"`
	Calendars []Calendar `json:"calendars"`
}
