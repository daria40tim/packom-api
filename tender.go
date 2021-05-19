package packom

import "time"

type Tender struct {
	Tender_id   int       `json:"tender_id"`
	Date        time.Time `json:"date"`
	Selected_cp int       `json:"selected_cp"`
	Tz_id       int       `json:"tz_id"`
	History     string    `json:"history"`
}
