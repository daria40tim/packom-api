package packom

import "mime/multipart"

type Org struct {
	O_id      int    `json:"o_id" db:"o_id"`
	Name      string `json:"name" db:"name"`
	Group     string `json:"group_id" db:"group_id"`
	Specs     string `json:"specs" db:"specs"`
	Countries string `json:"countries" db:"countries"`
	Site      string `json:"site" db:"site"`
	Phone     string `json:"phone" db:"phone"`
	Email     string `json:"email_" db:"email"`
	Adress    string `json:"adress" db:"adress"`
	Info      string `json:"info" db:"info"`
	Login     string `json:"email" db:"login"`
	Pwd       string `json:"password" db:"pwd"`
	Status    int    `json:"status" db:"status"`
	History   string `json:"history" db:"history"`
}

type OrgAll struct {
	O_id      int    `json:"o_id" db:"o_id"`
	Name      string `json:"name" db:"name"`
	Group     string `json:"group" db:"group"`
	Specs     string `json:"spec" db:"specs"`
	Countries string `json:"country" db:"countries"`
	Site      string `json:"site" db:"site"`
	Phone     string `json:"phone" db:"phone"`
	Email     string `json:"email" db:"email"`
}

type OrgId struct {
	O_id      int      `json:"o_id" db:"o_id"`
	Name      string   `json:"name" db:"name"`
	Group     string   `json:"group" db:"group"`
	Specs     string   `json:"spec" db:"specs"`
	Countries string   `json:"country" db:"countries"`
	Site      string   `json:"site" db:"site"`
	Phone     string   `json:"phone" db:"phone"`
	Email     string   `json:"email" db:"email"`
	Adress    string   `json:"adress" db:"adress"`
	Info      string   `json:"info" db:"info"`
	Status    bool     `json:"status" db:"status"`
	History   string   `json:"history" db:"history"`
	Login     string   `json:"login" db:"login"`
	Orgs      []OrgAll `json:"orgs"`
	Docs      []string `json:"docs"`
}

type OrgI struct {
	O_id      int      `json:"o_id" db:"o_id"`
	Name      string   `json:"name" db:"name"`
	Group     string   `json:"group" db:"group"`
	Specs     string   `json:"spec" db:"specs"`
	Countries string   `json:"country" db:"countries"`
	Site      string   `json:"site" db:"site"`
	Phone     string   `json:"phone" db:"phone"`
	Email     string   `json:"email" db:"email"`
	Adress    string   `json:"adress" db:"adress"`
	Info      string   `json:"info" db:"info"`
	Status    bool     `json:"status" db:"status"`
	History   string   `json:"history" db:"history"`
	Login     string   `json:"login" db:"login"`
	Pwd       string   `json:"password" db:"pwd"`
	Docs      []string `json:"docs"`
}

type Specs struct {
	Specs []string `json:"specs"`
}

type Org_docs struct {
	Docs *multipart.FileHeader `form:"doc"`
}
