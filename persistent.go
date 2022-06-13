package main

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type PokedexOp struct {
	db *bun.DB
}

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

func main() {
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

	if err := setSchema(ctx, db); err != nil {
		panic(err)
	}
}

func setSchema(ctx context.Context, db *bun.DB) error {
	if err := db.ResetModel(ctx, (*Pokemon)(nil), (*PokemonToType)(nil), (*Type)(nil)); err != nil {
		return err
	}
	pokemon := []Pokemon{
		{
			Name:        "Bulbasaur",
			Description: "There is a plant seed on its back right from the day this Pokémon is born. The seed slowly grows larger.",
			Category:    "Seed",
			Abilities:   []string{"Overgrow", "Hydropump"},
		},
		{
			Name:        "Ivysaur",
			Description: "When the bulb on its back grows large, it appears to lose the ability to stand on its hind legs.",
			Category:    "Seed",
			Abilities:   []string{"Overgrow", "Rock Throw"},
		},
	}
	if _, err := db.NewInsert().Model(&pokemon).Exec(ctx); err != nil {
		return err
	}

	pokemontype := []Type{
		{TypeName: "GRASS"},
		{TypeName: "DARK"},
		{TypeName: "POISON"},
		{TypeName: "FIRE"},
	}

	if _, err := db.NewInsert().Model(&pokemontype).Exec(ctx); err != nil {
		return err
	}

	pokemontotype := []PokemonToType{
		{PokemonID: 1, TypeID: 1},
		{PokemonID: 1, TypeID: 3},
		{PokemonID: 2, TypeID: 2},
	}

	if _, err := db.NewInsert().Model(&pokemontotype).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (op *PokedexOp) CreatePokemon(ctx context.Context, pokemon Pokemon) error {
	newpokemon := []Pokemon{pokemon}
	if _, err := op.db.NewInsert().Model(&newpokemon).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (op *PokedexOp) SearchByID(ctx context.Context, id int) (*Pokemon, error) {
	pokemons := new(Pokemon)
	if err := op.db.NewSelect().Model(pokemons).Relation("Types").Where("id = ? ", id).Scan(ctx); err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (op *PokedexOp) SearchByName(ctx context.Context, name string) (*Pokemon, error) {
	pokemons := new(Pokemon)
	if err := op.db.NewSelect().Model(pokemons).Relation("Types").Where("name LIKE ? ", name).Scan(ctx); err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (op *PokedexOp) ListAll(ctx context.Context) ([]Pokemon, error) {
	pokemons := make([]Pokemon, 0, 4)
	if err := op.db.NewSelect().Model(&pokemons).Column("name").OrderExpr("id ASC").Scan(ctx); err != nil {
		return nil, err
	}

	// for _, pokemon := range pokemons {
	// 	fmt.Println(pokemon.Name)
	// }
	return pokemons, nil
}

func (op *PokedexOp) UpdatePokemon(ctx context.Context, id int, name string) error {
	pokemon := &Pokemon{ID: id, Name: name}

	_, err := op.db.NewUpdate().Model(pokemon).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (op *PokedexOp) DeletePokemon(ctx context.Context, id int) error {
	pokemon := new(Pokemon)
	_, err := op.db.NewDelete().Model(pokemon).Where("id = ? ", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
