package packom

type Tender struct {
	Tender_id   int    `json:"tender_id"`
	Date        string `json:"date"`
	Selected_cp int    `json:"selected_cp"`
	Tz_id       int    `json:"tz_id"`
	History     string `json:"history"`
}

type TenderAll struct {
	Tender_id   int    `json:"tender_id" db:"tender_id"`
	Tz_id       int    `json:"tz_id" db:"tz_id"`
	Tz_st       int    `json:"tz_st" db:"tz_st"`
	Date        string `json:"date" db:"date"`
	Proj        string `json:"proj" db:"proj"`
	Group       string `json:"group" db:"group"`
	Type        string `json:"type" db:"type"`
	Kind        string `json:"kind" db:"kind"`
	Task        string `json:"task" db:"task"`
	Cp_count    string `json:"cp_count" db:"cp_count"`
	Tender_sum  int    `json:"tender_sum" db:"tender_sum"`
	Selected_cp int    `json:"selected_cp" db:"selected_cp"`
}

type TenderById struct {
	Tender_id    int        `json:"tender_id" db:"tender_id"`
	Client       string     `json:"client" db:"client"`
	Date         string     `json:"date" db:"date"`
	Task         string     `json:"task" db:"task"`
	Proj         string     `json:"proj" db:"proj"`
	Tz_id        int        `json:"tz_id" db:"tz_id"`
	Cp_count     string     `json:"cp_count" db:"cp_count"`
	Max_cp       string     `json:"max_cp" db:"max_cp"`
	Min_cp       string     `json:"min_cp" db:"min_cp"`
	Selected_cp  int        `json:"selected_cp" db:"selected_cp"`
	Term         int        `json:"term" db:"term"`
	Tz_Costs     []Cost     `json:"tz_costs"`
	Tz_Calendars []string   `json:"tz_calendars"`
	CPs          []TenderCP `json:"cps"`
	Active       bool       `json:"active" db:"active"`
}

type FullCost struct {
	Sum   float32 `db:"sum"`
	Cp_id int     `db:"cp_id"`
	O_id  int     `db:"o_id"`
}

type TenderCost struct {
	Task     string `json:"task" db:"task"`
	Cost_sum string `json:"cost_sum" db:"cost_sum"`
}

type TenderCP struct {
	Org       string       `json:"org" db:"org"`
	Cp_id     int          `json:"cp_id" db:"cp_id"`
	O_id      int          `json:"o_id" db:"o_id"`
	Sum       float32      `json:"sum" db:"sum"`
	Pay_cond  string       `json:"pay_cond" db:"pay_cond"`
	Calendars []int        `json:"calendars"`
	Costs     []TenderCost `json:"costs"`
}
