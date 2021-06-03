package packom

type CP struct {
	Cp_id        int        `json:"cp_id" db:"cp_id"`
	Date         string     `json:"date" db:"date"`
	Tz_Date      string     `json:"tz_date" db:"tz_date"`
	Tz_info      string     `json:"tz_info" db:"tz_info"`
	Cp_st        int        `json:"cp_st" db:"cp_st"`
	Tz_id        int        `json:"tz_id" db:"tz_id"`
	Proj         string     `json:"proj" db:"proj"`
	O_id         string     `json:"o_id" db:"o_id"`
	Tz_O_id      string     `json:"tz_o_id" db:"tz_o_id"`
	Org          string     `json:"org" db:"org"`
	Pay_cond     string     `json:"pay_cond" db:"pay_cond"`
	Task_name    string     `json:"task_name" db:"task_name"`
	Tz_Pay_cond  string     `json:"tz_pay_cond" db:"tz_pay_cond"`
	End_date     string     `json:"end_date" db:"end_date"`
	Tz_End_date  string     `json:"tz_end_date" db:"tz_end_date"`
	Client       string     `json:"client" db:"client"`
	Info         string     `json:"info" db:"info"`
	History      string     `json:"history" db:"history"`
	Group        string     `json:"group" db:"group"`
	Type         string     `json:"type" db:"type"`
	Kind         string     `json:"kind" db:"kind"`
	Private      string     `json:"privacy" db:"private"`
	Costs        []Cost     `json:"costs"`
	Calendars    []Calendar `json:"calendars"`
	Tz_Costs     []Cost     `json:"tz_costs"`
	Tz_Calendars []Calendar `json:"tz_calendars"`
	Docs         []string   `json:"docs"`
	Tz_Docs      []string   `json:"tz_docs"`
}

type CPId struct {
	Cp_id        int          `json:"cp_id" db:"cp_id"`
	Date         string       `json:"date" db:"date"`
	Tz_Date      string       `json:"tz_date" db:"tz_date"`
	Tz_info      string       `json:"tz_info" db:"tz_info"`
	Cp_st        int          `json:"cp_st" db:"cp_st"`
	Tz_id        int          `json:"tz_id" db:"tz_id"`
	Proj         string       `json:"proj" db:"proj"`
	O_id         string       `json:"o_id" db:"o_id"`
	Tz_O_id      string       `json:"tz_o_id" db:"tz_o_id"`
	Org          string       `json:"org" db:"org"`
	Pay_cond     string       `json:"pay_cond" db:"pay_cond"`
	Task_name    string       `json:"task_name" db:"task_name"`
	Tz_Pay_cond  string       `json:"tz_pay_cond" db:"tz_pay_cond"`
	End_date     string       `json:"end_date" db:"end_date"`
	Tz_End_date  string       `json:"tz_end_date" db:"tz_end_date"`
	Client       string       `json:"client" db:"client"`
	Info         string       `json:"info" db:"info"`
	History      string       `json:"history" db:"history"`
	Group        string       `json:"group" db:"group"`
	Type         string       `json:"type" db:"type"`
	Kind         string       `json:"kind" db:"kind"`
	Private      string       `json:"privacy" db:"private"`
	Costs        []Cost       `json:"costs"`
	Calendars    []CPCalendar `json:"calendars"`
	Tz_Costs     []Cost       `json:"tz_costs"`
	Tz_Calendars []Calendar   `json:"tz_calendars"`
	Docs         []string     `json:"docs"`
	Tz_Docs      []string     `json:"tz_docs"`
}

type CPAll struct {
	Cp_id     string `json:"cp_id" db:"cp_id"`
	Date      string `json:"date" db:"date"`
	Cp_st     string `json:"cp_st" db:"cp_st"`
	Tz_id     int    `json:"tz_id" db:"tz_id"`
	Proj      string `json:"proj" db:"proj"`
	Client    string `json:"client" db:"client"`
	O_id      int    `json:"o_id" db:"o_id"`
	Group     string `json:"group" db:"group"`
	Type      string `json:"type" db:"type"`
	Kind      string `json:"kind" db:"kind"`
	Task_name string `json:"task_name" db:"task_name"`
}

type CP_srv struct {
	Date     string `json:"date" db:"date"`
	Tz_id    int    `json:"tz_id" db:"tz_id"`
	O_id     int    `json:"o_id" db:"o_id"`
	End_date string `json:"end_date" db:"end_date"`
	Cp_id    int    `json:"cp_id" db:"cp_id"`
}

type CPIns struct {
	Date      string     `json:"date" db:"date"`
	Tz_id     int        `json:"tz_id" db:"tz_id"`
	Proj      string     `json:"proj" db:"proj"`
	Pay_cond  string     `json:"pay_cond" db:"pay_cond"`
	End_date  string     `json:"end_date" db:"end_date"`
	Info      string     `json:"info" db:"info"`
	History   string     `json:"history" db:"history"`
	Costs     []Cost     `json:"cst"`
	Calendars []Calendar `json:"cal"`
	Docs      []string   `json:"docs"`
}
