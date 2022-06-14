package main

import "github.com/uptrace/bun"

type Pokemon struct {
	bun.BaseModel `bun:"table:pokemon"`

	ID          int `bun:",pk,autoincrement"`
	Name        string
	Description string
	Category    string
	Abilities   []string
	Types       []Type `bun:"m2m:pokemon_to_type,join:Pokemon=Type"`
}

type PokemonToType struct {
	bun.BaseModel `bun:"table:pokemon_to_type"`

	PokemonID int      `bun:",pk"`
	Pokemon   *Pokemon `bun:"rel:belongs-to,join:pokemon_id=id"`
	TypeID    int      `bun:",pk"`
	Type      *Type    `bun:"rel:belongs-to,join:type_id=id"`
}

type Type struct {
	bun.BaseModel `bun:"table:types"`

	ID       int `bun:",pk,autoincrement"`
	TypeName string
	Pokemons []Pokemon `bun:"m2m:pokemon_to_type,join:Type=Pokemon"`
}
