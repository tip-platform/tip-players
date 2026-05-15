// Package di wires player adapters and services.
package di

import (
	"github.com/tip-platform/tip-players/internal/adapter/driven/persistence"
	"github.com/tip-platform/tip-players/internal/service"
)

type Container struct {
	PlayerService *service.PlayerService
}

func NewContainer() (*Container, error) {
	store, e := persistence.NewPlayerStore()

	if e != nil {
		return nil, e
	}

	repo := persistence.NewPlayerRepositoryAdapter(store)

	playerService := service.NewPlayerService(repo)

	return &Container{PlayerService: playerService}, nil
}
