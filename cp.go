package packom

type CP struct {
	Cp_id        int        `json:"cp_id" db:"cp_id"`
	Date         string     `json:"date" db:"date"`
	Cp_st        int        `json:"cp_st" db:"cp_st"`
	Tz_id        int        `json:"tz_id" db:"tz_id"`
	Proj         string     `json:"proj" db:"proj"`
	O_id         string     `json:"o_id" db:"o_id"`
	Pay_cond     string     `json:"pay_cond" db:"pay_cond"`
	End_date     string     `json:"end_date" db:"end_date"`
	Info         string     `json:"info" db:"info"`
	History      string     `json:"history" db:"history"`
	Costs        []Cost     `json:"costs"`
	Calendars    []Calendar `json:"calendars"`
	Tz_Costs     []Cost     `json:"tz_costs"`
	Tz_Calendars []Calendar `json:"tz_calendars"`
}

type CPAll struct {
	Cp_id  string `json:"cp_id" db:"cp_id"`
	Date   string `json:"date" db:"date"`
	Cp_st  string `json:"cp_st" db:"cp_st"`
	Tz_id  int    `json:"tz_id" db:"tz_id"`
	Proj   string `json:"proj" db:"proj"`
	Client string `json:"client" db:"client"`
	O_id   int    `json:"o_id" db:"o_id"`
	Group  string `json:"group" db:"group"`
	Type   string `json:"type" db:"type"`
	Kind   string `json:"kind" db:"kind"`
}

type CP_srv struct {
	Date     string `json:"date" db:"date"`
	Tz_id    int    `json:"tz_id" db:"tz_id"`
	O_id     int    `json:"o_id" db:"o_id"`
	End_date string `json:"end_date" db:"end_date"`
	Cp_id    int    `json:"cp_id" db:"cp_id"`
}
