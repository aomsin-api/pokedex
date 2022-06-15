package database

import (
	"context"
	"database/sql"
	"fmt"
	"pokedex/graph/gqlmodel"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type PokedexOp struct {
	Db *bun.DB
}

func PokedexInit() (*bun.DB, error) {
	sqlDB, err := sql.Open(sqliteshim.ShimName, "pokedex.db")
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqlDB, sqlitedialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db, nil
}

func (op *PokedexOp) CreatePokemon(ctx context.Context, input *CreatePokemonInput) (*Pokemon, error) {
	newPokemon := Pokemon{
		Name:        *input.Name,
		Description: *input.Description,
		Category:    *input.Category,
		Abilities:   input.Abilities,
		Type:        input.Type,
	}
	if _, err := op.Db.NewInsert().Model(&newPokemon).Exec(ctx); err != nil {
		return nil, err
	}
	return &newPokemon, nil
}

func (op *PokedexOp) UpdatePokemon(ctx context.Context, id *string, input *UpdatePokemonInput) (*Pokemon, error) {
	pokemon, err := op.SearchByID(ctx, *id)
	if err != nil {
		return nil, err
	}
	if input.Name != nil {
		pokemon.Name = *input.Name
	}
	if input.Description != nil {
		pokemon.Description = *input.Description
	}
	if input.Category != nil {
		pokemon.Category = *input.Category
	}
	if input.Abilities != nil {
		pokemon.Abilities = input.Abilities
	}
	if input.Type != nil {
		pokemon.Type = input.Type
	}
	_, err = op.Db.NewUpdate().Model(pokemon).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return op.SearchByID(ctx, *id)
}

func (op *PokedexOp) DeletePokemon(ctx context.Context, id string) error {
	pokemon := new(Pokemon)
	_, err := op.Db.NewDelete().Model(pokemon).Where("id = ? ", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (op *PokedexOp) SearchByID(ctx context.Context, id string) (*Pokemon, error) {
	pokemons := new(Pokemon)
	if err := op.Db.NewSelect().Model(pokemons).Where("id = ? ", id).Scan(ctx); err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (op *PokedexOp) SearchByName(ctx context.Context, name string) (*Pokemon, error) {
	pokemons := new(Pokemon)
	if err := op.Db.NewSelect().Model(pokemons).Where("name LIKE ? ", name).Scan(ctx); err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (op *PokedexOp) ListAll(ctx context.Context) ([]*Pokemon, error) {
	pokemons := make([]*Pokemon, 0)
	if err := op.Db.NewSelect().Model(&pokemons).OrderExpr("id ASC").Scan(ctx); err != nil {
		return nil, err
	}

	return pokemons, nil
}

func CheckInput(input gqlmodel.PokemonInput) error {
	if input.Name == nil {
		return fmt.Errorf("name must not be null")
	}
	if input.Description == nil {
		return fmt.Errorf("description must not be null")
	}
	if input.Category == nil {
		return fmt.Errorf("category must not be null")
	}
	if input.Abilities == nil {
		return fmt.Errorf("abilities must not be null")
	}
	if input.Type == nil {
		return fmt.Errorf("type must not be null")
	}

	return nil
}
