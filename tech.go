package packom

type Tech struct {
	Tz_id     int        `json:"tz_id" db:"tz_id"`
	O_id      int        `json:"o_id" db:"o_id"`
	Date      string     `json:"date" db:"date"`
	Proj      string     `json:"proj" db:"proj"`
	Client    string     `json:"client" db:"client"`
	End_date  string     `json:"end_date" db:"end_date"`
	Group     string     `json:"group" db:"group"`
	Type      string     `json:"type" db:"type"`
	Kind      string     `json:"kind" db:"kind"`
	Tz_st     bool       `json:"tz_st" db:"tz_st"`
	Tender_st int        `json:"tender_st" db:"tender_st"`
	Cp_st     int        `json:"cp_st" db:"cp_st"`
	Pay_cond  string     `json:"pay_cond" db:"pay_cond"`
	Private   bool       `json:"private" db:"private"`
	Info      string     `json:"info" db:"info"`
	History   string     `json:"history" db:"history"`
	Costs     []Cost     `json:"costs"`
	Calendars []Calendar `json:"calendars"`
}

type TechAll struct {
	Date      string `json:"date" db:"date"`
	Client    string `json:"client" db:"client"`
	O_id      int    `json:"o_id" db:"o_id"`
	Tz_id     int    `json:"tz_id" db:"tz_id"`
	End_date  string `json:"end_date" db:"end_date"`
	Proj      string `json:"proj" db:"proj"`
	Group     string `json:"group" db:"group"`
	Type      string `json:"type" db:"type"`
	Kind      string `json:"kind" db:"kind"`
	Task      string `json:"task" db:"task"`
	Tz_st     string `json:"tz_st" db:"tz_st"`
	Tender_st string `json:"tender_st" db:"tender_st"`
	CP_count  int    `json:"count" db:"cp_count"`
	CP_st     string `json:"cp_st" db:"cp_st"`
}
