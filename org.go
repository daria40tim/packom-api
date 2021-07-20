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
	Specs     []string `json:"specs" db:"specs"`
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
	Trusted   []int    `json:"trusted"`
}

type OrgI struct {
	O_id      int      `json:"o_id" db:"o_id"`
	Name      string   `json:"name" db:"name"`
	Group     string   `json:"group" db:"group"`
	Specs     []Spec   `json:"spec" db:"specs"`
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

type OrgFiltered struct {
	O_id      int    `json:"o_id" db:"o_id"`
	Name      string `json:"name" db:"name"`
	Group     string `json:"group" db:"group"`
	Specs     string `json:"spec" db:"specs"`
	Countries string `json:"country" db:"countries"`
	Site      string `json:"site" db:"site"`
	Phone     string `json:"phone" db:"phone"`
	Email     string `json:"email" db:"email"`
}

type Specs struct {
	Specs []string `json:"specs"`
}

type Spec struct {
	Name   string `json:"name" db:"name"`
	Active bool   `json:"active" db:"active"`
}

type Countries struct {
	Countries []string `json:"countries"`
}

type Org_docs struct {
	Docs *multipart.FileHeader `form:"doc"`
}

type OrgFilterData struct {
	Names     []NameFilter    `json:"names"`
	Countries []CountryFilter `json:"countries"`
	Specs     []SpecFilter    `json:"specs"`
}

type NameFilter struct {
	Id   int    `db:"o_id" json:"id"`
	Name string `db:"name" json:"name"`
}

type CountryFilter struct {
	Id      int    `db:"country_id" json:"id"`
	Country string `db:"name" json:"name"`
}

type SpecFilter struct {
	Id   int    `db:"spec_id" json:"id"`
	Spec string `db:"name" json:"name"`
}

type OrgFilterParams struct {
	Names     []int `json:"names"`
	Groups    []int `json:"groups"`
	Specs     []int `json:"specs"`
	Countries []int `json:"countries"`
}
