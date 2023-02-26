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

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, receiverID *string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// DeliverOrder is the resolver for the deliverOrder field.
func (r *mutationResolver) DeliverOrder(ctx context.Context, orderID *string, amount *int) (*model.Response, error) {
	panic(fmt.Errorf("not implemented: DeliverOrder - deliverOrder"))
}

// UpdateWeeklyOrder is the resolver for the updateWeeklyOrder field.
func (r *mutationResolver) UpdateWeeklyOrder(ctx context.Context, playerID *string, amount *int) (*model.Response, error) {
	panic(fmt.Errorf("not implemented: UpdateWeeklyOrder - updateWeeklyOrder"))
}

// Sender is the resolver for the sender field.
func (r *orderResolver) Sender(ctx context.Context, obj *model.Order) (*model.Player, error) {
	panic(fmt.Errorf("not implemented: Sender - sender"))
}

// Receiver is the resolver for the receiver field.
func (r *orderResolver) Receiver(ctx context.Context, obj *model.Order) (*model.Player, error) {
	panic(fmt.Errorf("not implemented: Receiver - receiver"))
}

// Board is the resolver for the board field.
func (r *orderResolver) Board(ctx context.Context, obj *model.Order) (*model.Board, error) {
	panic(fmt.Errorf("not implemented: Board - board"))
}

// NewOrder is the resolver for the newOrder field.
func (r *subscriptionResolver) NewOrder(ctx context.Context, playerID *string) (<-chan *model.Order, error) {
	panic(fmt.Errorf("not implemented: NewOrder - newOrder"))
}

// OrderDelivery is the resolver for the orderDelivery field.
func (r *subscriptionResolver) OrderDelivery(ctx context.Context, playerID *string) (<-chan *model.Order, error) {
	panic(fmt.Errorf("not implemented: OrderDelivery - orderDelivery"))
}

// Order returns graph.OrderResolver implementation.
func (r *Resolver) Order() graph.OrderResolver { return &orderResolver{r} }

type orderResolver struct{ *Resolver }