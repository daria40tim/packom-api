package packom

type Cost struct {
	Cost_id int     `json:"cost_id"`
	Task    string  `json:"task"`
	Metr    string  `json:"metr"`
	Count   int     `json:"count"`
	Tz_id   int     `json:"tz_id"`
	Cp_id   int     `json:"cp_id"`
	PPU     float32 `json:"ppu"`
	Info    string  `json:"info"`
}
