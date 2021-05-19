package packom

type Org struct {
	O_id      int      `json:"o_id" db:"o_id"`
	Name      string   `json:"name"`
	Group     string   `json:"group"`
	Specs     []string `json:"specs"`
	Countries []string `json:"countries"`
	Site      string   `json:"site"`
	Phone     string   `json:"phone"`
	Email     string   `json:"email"`
	Adress    string   `json:"adress"`
	Info      string   `json:"info"`
	Login     string   `json:"login"`
	Pwd       string   `json:"pwd"`
	Status    int      `json:"status"`
	History   string   `json:"history"`
}
