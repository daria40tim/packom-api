package packom

type Calendar struct {
	Cal_id    int    `json:"cal_id" db:"cal_id"`
	Name      string `json:"task_name" db:"name"`
	Period    int    `json:"period" db:"period"`
	Term      int    `json:"term" db:"term"`
	Tz_id     int    `json:"tz_id" db:"tz_id"`
	Cp_id     int    `json:"cp_id" db:"cp_id"`
	Active    bool   `json:"active" db:"active"`
	CP_Period int    `json:"cp_period" db:"cp_period"`
}

type CPCalendar struct {
	Cal_id    int    `json:"cal_id" db:"cal_id"`
	Name      string `json:"task_name" db:"name"`
	Period    int    `json:"period" db:"period"`
	Term      int    `json:"term" db:"term"`
	TZ_Period int    `json:"tz_period" db:"tz_period"`
	TZ_Term   int    `json:"tz_term" db:"tz_term"`
	Tz_id     int    `json:"tz_id" db:"tz_id"`
	Cp_id     int    `json:"cp_id" db:"cp_id"`
	Active    bool   `json:"active" db:"active"`
	CP_Period int    `json:"cp_period" db:"cp_period"`
}

type InputCalendar struct {
	Name      string `json:"task_name" db:"name"`
	Period    int    `json:"period" db:"period"`
	Term      int    `json:"term" db:"term"`
	CP_Period int    `json:"cp_period" db:"cp_period"`
	CP_Term   int    `json:"cp_term"`
	Tz_id     int    `json:"tz_id" db:"tz_id"`
	Cp_id     int    `json:"cp_id" db:"cp_id"`
	Active    bool   `json:"active" db:"active"`
}
