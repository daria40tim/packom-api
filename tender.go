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
	Date        string `json:"date" db:"date"`
	Proj        string `json:"proj" db:"proj"`
	Group       string `json:"group" db:"group"`
	Type        string `json:"type" db:"type"`
	Kind        string `json:"kind" db:"kind"`
	Selected_cp int    `json:"selected_cp" db:"selected_cp"`
}
