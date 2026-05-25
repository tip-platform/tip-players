// Package di wires player adapters and services.
package di

import (
	"github.com/tip-platform/tip-players/internal/app"
	"github.com/tip-platform/tip-players/internal/app/port/input"
	"github.com/tip-platform/tip-players/internal/infra/driven/persistence"
)

type Container struct {
	PlayerService input.PlayerService
	PlayerStore   *persistence.PlayerStore
}

func NewContainer() (*Container, error) {
	store, e := persistence.NewPlayerStore()

	if e != nil {
		return nil, e
	}

	repo := persistence.NewMSSQLPlayerRepositoryAdapter(store)

	playerService := app.NewPlayerService(repo)

	return &Container{PlayerService: playerService, PlayerStore: store}, nil
}
