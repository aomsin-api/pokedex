package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"pokedex/database"
	"pokedex/graph/generated"
	"pokedex/graph/gqlmodel"
)

func (r *mutationResolver) CreatePokemon(ctx context.Context, input gqlmodel.PokemonInput) (*database.Pokemon, error) {
	newpokemon, err := r.Pokedex.CreatePokemon(ctx, &database.CreatePokemonInput{
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Abilities:   input.Abilities,
		Type:        input.Type,
	})
	if err != nil {
		return nil, err
	}

	return newpokemon, nil
}

func (r *mutationResolver) UpdatePokemon(ctx context.Context, input gqlmodel.PokemonInput) (*database.Pokemon, error) {
	if input.ID == nil {
		return nil, fmt.Errorf("id must not be null")
	}
	pokemon, err := r.Pokedex.UpdatePokemon(ctx, &database.UpdatePokemonInput{
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Abilities:   input.Abilities,
		Type:        input.Type,
	})
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

func (r *mutationResolver) DeletePokemon(ctx context.Context, id string) (bool, error) {
	err := r.Pokedex.DeletePokemon(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) SearchPokemonByID(ctx context.Context, id string) (*database.Pokemon, error) {
	return r.Pokedex.SearchByID(ctx, id)
}

func (r *queryResolver) SearchPokemonByName(ctx context.Context, name string) (*database.Pokemon, error) {
	return r.Pokedex.SearchByName(ctx, name)
}

func (r *queryResolver) Pokemons(ctx context.Context) ([]*database.Pokemon, error) {
	return r.Pokedex.ListAll(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
