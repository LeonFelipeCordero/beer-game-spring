package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/LeonFelipeCordero/golang-beer-game/graph"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
)

// AddPlayer is the resolver for the addPlayer field.
func (r *mutationResolver) AddPlayer(ctx context.Context, boardID *string, role *model.Role) (*model.Player, error) {
	return r.PlayerApiAdapter.AddPlayer(ctx, *boardID, role.String())
}

// Board is the resolver for the board field.
func (r *playerResolver) Board(ctx context.Context, obj *model.Player) (*model.Board, error) {
	return r.BoardApiAdapter.GetByPlayer(ctx, obj.ID)
}

// Orders is the resolver for the orders field.
func (r *playerResolver) Orders(ctx context.Context, obj *model.Player) ([]*model.Order, error) {
	return r.OrderApiAdapter.LoadByPlayer(ctx, obj.ID)
}

// GetPlayer is the resolver for the getPlayer field.
func (r *queryResolver) GetPlayer(ctx context.Context, playerID *string) (*model.Player, error) {
	return r.PlayerApiAdapter.Get(ctx, *playerID)
}

// GetPlayersByBoard is the resolver for the getPlayersByBoard field.
func (r *queryResolver) GetPlayersByBoard(ctx context.Context, boardID *string) ([]*model.Player, error) {
	return r.PlayerApiAdapter.GetPlayersByBoard(ctx, *boardID)
}

// UpdateWeeklyOrder is the resolver for the updateWeeklyOrder field.
func (r *queryResolver) UpdateWeeklyOrder(ctx context.Context, playerID *string, amount *int) (*model.Response, error) {
	return r.PlayerApiAdapter.UpdateWeeklyOrder(ctx, *playerID, *amount)
}

// Player is the resolver for the player field.
func (r *subscriptionResolver) Player(ctx context.Context, playerID *string) (<-chan *model.Player, error) {
	panic(fmt.Errorf("not implemented: Player - player"))
}

// Player returns graph.PlayerResolver implementation.
func (r *Resolver) Player() graph.PlayerResolver { return &playerResolver{r} }

type playerResolver struct{ *Resolver }
