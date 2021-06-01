package packom

type TechFilter struct {
	Tech_date_from string `json:"tech_date_from"`
	Tech_date_to   string `json:"tech_date_to"`
	Tech_o_id      int    `json:"tech_o_id"`
	Tech_id        int    `json:"tech_id"`
	Tech_proj      string `json:"tech_proj"`
	Tech_st        int    `json:"tech_st"`
	Tech_tender_st int    `json:"tech_tender_st"`
	Tech_cp_st     int    `json:"tech_cp_st"`
}

type CPFilter struct {
	CP_id        int    `json:"cp_id"`
	CP_date_from string `json:"cp_date_from"`
	CP_date_to   string `json:"cp_date_to"`
	CP_st        int    `json:"cp_st"`
	CP_tz_id     int    `json:"cp_tz_id"`
	CP_proj      string `json:"cp_proj"`
	CP_o_id      int    `json:"cp_o_id"`
}

type OrgFilter struct {
	O_id      int    `json:"o_id"`
	O_name    string `json:"o_name"`
	O_group   int    `json:"o_group"`
	O_country string `json:"o_country"`
	O_spec    string `json:"o_spec"`
}

type TenderFilter struct {
	Tender_id        int    `json:"tender_id"`
	Tender_tz_id     int    `json:"tender_tz_id"`
	Tender_date_from string `json:"tender_date_from"`
	Tender_date_to   string `json:"tender_date_to"`
	Tender_proj      string `json:"tender_proj"`
	Tender_st        int    `json:"tender_st"`
}
