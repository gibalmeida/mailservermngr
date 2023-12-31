// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package mysql

import (
	"database/sql"
)

type Account struct {
	Name          string         `json:"name"`
	Domain        string         `json:"domain"`
	Password      string         `json:"password"`
	HomeDir       sql.NullString `json:"home_dir"`
	Quota         sql.NullString `json:"quota"`
	ClearPassword sql.NullString `json:"clear_password"`
}

type AddressAlias struct {
	Alias     string         `json:"alias"`
	Addresses sql.NullString `json:"addresses"`
}

type Domain struct {
	Domain string `json:"domain"`
}

type DomainAlias struct {
	Alias  string `json:"alias"`
	Domain string `json:"domain"`
}
