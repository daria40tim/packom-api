package packom

type Tech struct {
	Tz_id     int        `json:"tz_id" db:"tz_id"`
	O_id      int        `json:"o_id" db:"o_id"`
	Date      string     `json:"date" db:"date"`
	Proj      string     `json:"proj" db:"proj"`
	Task      string     `json:"task" db:"task_name"`
	Client    string     `json:"client" db:"client"`
	End_date  string     `json:"end_date" db:"end_date"`
	Group     string     `json:"group" db:"group"`
	Type      string     `json:"type" db:"type"`
	Kind      string     `json:"kind" db:"kind"`
	Tz_st     string     `json:"tz_st" db:"tz_st"`
	Tender_st int        `json:"tender_st" db:"tender_st"`
	Cp_st     int        `json:"cp_st" db:"cp_st"`
	Pay_cond  string     `json:"pay_cond" db:"pay_cond"`
	Private   string     `json:"privacy" db:"private"`
	Info      string     `json:"info" db:"info"`
	History   string     `json:"history" db:"history"`
	Costs     []Cost     `json:"cst"`
	Calendars []Calendar `json:"cal"`
	Docs      []string   `json:"docs"`
}

type TechAll struct {
	Date         string `json:"date" db:"date"`
	Client       string `json:"client" db:"client"`
	O_id         int    `json:"o_id" db:"o_id"`
	Tz_id        int    `json:"tz_id" db:"tz_id"`
	End_date     string `json:"end_date" db:"end_date"`
	Proj         string `json:"proj" db:"proj"`
	Group        string `json:"group" db:"group"`
	Type         string `json:"type" db:"type"`
	Kind         string `json:"kind" db:"kind"`
	Task         string `json:"task" db:"task"`
	Tz_st        string `json:"tz_st"`
	Tz_st_id     int    `json:"tz_st_id"`
	Tender_st    string `json:"tender_st"`
	Tender_st_id int    `json:"tender_st_id"`
	CP_count     int    `json:"count" db:"cp_count"`
	CP_st        string `json:"cp_st"`
	CP_st_id     int    `json:"cp_st_id"`
	Selected_cp  int    `json:"selected_cp" db:"selected_cp"`
	Active       bool   `json:"active" db:"active"`
}

type Select struct {
	Metrics    []string `json:"metrics"`
	Groups     []string `json:"groups"`
	Kinds      []string `json:"kinds"`
	Types      []string `json:"types"`
	Pay_conds  []string `json:"pay_conds"`
	Task_names []string `json:"task_names"`
	Tasks      []string `json:"tasks"`
	Task_kinds []string `json:"task_kinds"`
}

type TechFilterData struct {
	Clients []ClientFilter `json:"clients"`
	Projs   []ProjFilter   `json:"projs"`
}

type ClientFilter struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type ProjFilter struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type TechFilterParams struct {
	EDate      string `json:"e_date"`
	SDate      string `json:"s_date"`
	Clients    []int  `json:"clients"`
	Projs      []int  `json:"projs"`
	TZ_STS     []int  `json:"tz_sts"`
	Tender_STS []int  `json:"tender_sts"`
	CP_STS     []int  `json:"cp_sts"`
}
