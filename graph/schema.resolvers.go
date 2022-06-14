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
	newPokemon := model.Pokemon{
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Abilities:   input.Abilities,
		Type:        input.Abilities,
	}

	err := r.pokedex.CreatePokemon(ctx, &newPokemon)
	if err != nil {
		return nil, err
	}

	return &newPokemon, nil
}

func (r *mutationResolver) UpdatePokemon(ctx context.Context, input model.PokemonInput) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePokemon(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pokemon(ctx context.Context, id string) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pokemons(ctx context.Context) ([]*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
