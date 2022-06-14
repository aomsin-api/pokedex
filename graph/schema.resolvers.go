package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"pokedex/graph/generated"
	"pokedex/graph/model"
)

func (r *mutationResolver) CreatePokemon(ctx context.Context, input model.PokemonInput) (*model.Pokemon, error) {
	newPokemon, err := r.pokedex.CreatePokemon(ctx, &input)
	if err != nil {
		return nil, err
	}

	return newPokemon, nil
}

func (r *mutationResolver) UpdatePokemon(ctx context.Context, input model.PokemonInput) (*model.Pokemon, error) {
	if input.ID == nil {
		return nil, fmt.Errorf("id must not be null")
	}

	pokemon := model.Pokemon{
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Abilities:   input.Abilities,
		Type:        input.Abilities,
	}

	err := r.pokedex.UpdatePokemon(ctx, &pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}

func (r *mutationResolver) DeletePokemon(ctx context.Context, id string) (bool, error) {
	err := r.pokedex.DeletePokemon(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) Pokemonbyid(ctx context.Context, id string) (*model.Pokemon, error) {
	return r.pokedex.SearchByID(ctx, id)
}

func (r *queryResolver) Pokemonbyname(ctx context.Context, name string) (*model.Pokemon, error) {
	return r.pokedex.SearchByName(ctx, name)
}

func (r *queryResolver) Pokemons(ctx context.Context) ([]*model.Pokemon, error) {
	return r.pokedex.ListAll(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
