package database

import "github.com/uptrace/bun"

type Pokemon struct {
	bun.BaseModel `bun:"table:pokemon"`
	ID            int `bun:",pk,autoincrement"`
	Name          string
	Description   string
	Category      string
	Abilities     []string
	Type          []string
}

type CreatePokemonInput struct {
	Name        *string
	Description *string
	Category    *string
	Abilities   []string
	Type        []string
}

type UpdatePokemonInput struct {
	ID          int
	Name        *string
	Description *string
	Category    *string
	Abilities   []string
	Type        []string
}
