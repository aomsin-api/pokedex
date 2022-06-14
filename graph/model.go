package graph

import "github.com/uptrace/bun"

type Pokemon struct {
	bun.BaseModel `bun:"table:pokemon"`

	ID          int `bun:",pk,autoincrement"`
	Name        string
	Description string
	Category    string
	Abilities   []string
	Types       []string
}
