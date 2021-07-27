package packom

type Cost struct {
	Cost_id int    `json:"cost_id" db:"cost_id"`
	Task    string `json:"task" db:"task"`
	Metr    string `json:"metr" db:"metr"`
	Count   int    `json:"count" db:"count"`
	Tz_id   int    `json:"tz_id" db:"tz_id"`
	Cp_id   int    `json:"cp_id" db:"cp_id"`
	PPU     int    `json:"ppu" db:"ppu"`
	Info    string `json:"info" db:"info"`
	Active  bool   `json:"active" db:"active"`
}
