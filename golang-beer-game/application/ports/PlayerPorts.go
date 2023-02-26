package ports

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
)

type IPlayerApi interface {
	AddPlayer(ctx context.Context, boardId string, role string) (*model.Player, error)
	Get(ctx context.Context, id string) (*model.Player, error)
}

type IPlayerService interface {
	AddPlayer(ctx context.Context, boardId string, role string) (*domain.Player, error)
	Get(ctx context.Context, id string) (*domain.Player, *string, error)
}

type IPlayerRepository interface {
	AddPlayer(ctx context.Context, boardId string, player domain.Player) (*domain.Player, error)
	Get(ctx context.Context, id string) (*domain.Player, *string, error)
	DeleteAll(ctx context.Context)
}
