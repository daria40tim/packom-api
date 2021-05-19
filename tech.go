package packom

type Tech struct {
	Tz_id     int        `json:"tz_id"`
	Date      string     `json:"date"`
	Proj      string     `json:"proj"`
	Client    string     `json:"client"`
	End_date  string     `json:"end_date"`
	Group     string     `json:"group"`
	Type      string     `json:"type"`
	Kind      string     `json:"kind"`
	Tz_st     bool       `json:"tz_st"`
	Tender_st int        `json:"tender_st"`
	Cp_st     int        `json:"cp_st"`
	Pay_cond  string     `json:"pay_cond"`
	Private   bool       `json:"private"`
	Info      string     `json:"info"`
	History   string     `json:"history"`
	Costs     []Cost     `json:"costs"`
	Calendars []Calendar `json:"calendars"`
}
