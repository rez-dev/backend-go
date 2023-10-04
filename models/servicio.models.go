package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Term struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Term_group int    `json:"term_group"`
}

type Terms []Term
