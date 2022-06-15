package database

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type PokedexOp struct {
	Db *bun.DB
}

func PokedexInit() (*bun.DB, error) {
	ctx := context.Background()

	sqlDB, err := sql.Open(sqliteshim.ShimName, "pokedex.db")
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqlDB, sqlitedialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	if err := SetSchema(ctx, db); err != nil {
		return nil, err
	}

	return db, nil
}

func SetSchema(ctx context.Context, Db *bun.DB) error {
	if err := Db.ResetModel(ctx, (*Pokemon)(nil)); err != nil {
		return err
	}
	firstpokemon := Pokemon{
		Name:        "Charmander",
		Description: "It has a preference for hot things. When it rains, steam is said to spout from the tip of its tail.",
		Category:    "Lizard",
		Abilities:   []string{"Blaze", "Flamethrower"},
		Type:        []string{"Fire", "Normal"},
	}

	if _, err := Db.NewInsert().Model(&firstpokemon).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (op *PokedexOp) CreatePokemon(ctx context.Context, input *CreatePokemonInput) (*Pokemon, error) {
	newPokemon := Pokemon{
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Abilities:   input.Abilities,
		Type:        input.Type,
	}
	if _, err := op.Db.NewInsert().Model(&newPokemon).Exec(ctx); err != nil {
		return nil, err
	}
	return &newPokemon, nil
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

func (op *PokedexOp) UpdatePokemon(ctx context.Context, input *UpdatePokemonInput) (*Pokemon, error) {
	updatePokemon := Pokemon{
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Abilities:   input.Abilities,
		Type:        input.Type,
	}
	_, err := op.Db.NewUpdate().Model(&updatePokemon).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &updatePokemon, nil
}

func (op *PokedexOp) DeletePokemon(ctx context.Context, id string) error {
	pokemon := new(Pokemon)
	_, err := op.Db.NewDelete().Model(pokemon).Where("id = ? ", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
