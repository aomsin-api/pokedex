package graph

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"

	"pokedex/graph/model"
)

type PokedexOp struct {
	db *bun.DB
	// pokemon		[]*model.pokemon
}

func ReSet() {
	ctx := context.Background()

	sqldb, err := sql.Open(sqliteshim.ShimName, "pokedex.db")
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	db.RegisterModel((*PokemonToType)(nil))

	if err := ResetSchema(ctx, db); err != nil {
		panic(err)
	}
}

func ResetSchema(ctx context.Context, db *bun.DB) error {
	if err := db.ResetModel(ctx, (*Pokemon)(nil), (*PokemonToType)(nil), (*Type)(nil)); err != nil {
		return err
	}

	pokemontype := []Type{
		{TypeName: "Bug"},
		{TypeName: "Dark"},
		{TypeName: "Dragon"},
		{TypeName: "Electric"},
		{TypeName: "Fairy"},
		{TypeName: "Fighting"},
		{TypeName: "Fire"},
		{TypeName: "Flying"},
		{TypeName: "Ghost"},
		{TypeName: "Gras"},
		{TypeName: "Ground"},
		{TypeName: "Ice"},
		{TypeName: "Normal"},
		{TypeName: "Poison"},
		{TypeName: "Psychic"},
		{TypeName: "Rock"},
		{TypeName: "Steel"},
		{TypeName: "Water"},
	}

	if _, err := db.NewInsert().Model(&pokemontype).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (op *PokedexOp) CreatePokemon(ctx context.Context, pokemon *model.Pokemon) error {

	if _, err := op.db.NewInsert().Model(&pokemon).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (op *PokedexOp) SearchByID(ctx context.Context, id string) (*model.Pokemon, error) {
	pokemons := new(*model.Pokemon)
	if err := op.db.NewSelect().Model(pokemons).Relation("Types").Where("id = ? ", id).Scan(ctx); err != nil {
		return nil, err
	}

	return *pokemons, nil
}

func (op *PokedexOp) SearchByName(ctx context.Context, name string) (*Pokemon, error) {
	pokemons := new(Pokemon)
	if err := op.db.NewSelect().Model(pokemons).Relation("Types").Where("name LIKE ? ", name).Scan(ctx); err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (op *PokedexOp) ListAll(ctx context.Context) ([]*model.Pokemon, error) {
	pokemons := make([]*model.Pokemon, 0, 4)
	if err := op.db.NewSelect().Model(&pokemons).Column("name").OrderExpr("id ASC").Scan(ctx); err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (op *PokedexOp) UpdatePokemon(ctx context.Context, pokemon *model.Pokemon) error {
	_, err := op.db.NewUpdate().Model(&pokemon).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (op *PokedexOp) DeletePokemon(ctx context.Context, id string) error {
	pokemon := new(Pokemon)
	_, err := op.db.NewDelete().Model(pokemon).Where("id = ? ", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
