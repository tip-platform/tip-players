// Package rpc provides gRPC adapters for player services.
package rpc

import (
	"context"

	"github.com/tip-platform/tip-players/internal/domain/entity"
	"github.com/tip-platform/tip-players/internal/port/input"
)

type PlayerRepositoryAdapter struct {
	service input.PlayerService
}

func NewPlayerRepositoryAdapter(service input.PlayerService) *PlayerRepositoryAdapter {
	return &PlayerRepositoryAdapter{service: service}
}

func (r *PlayerRepositoryAdapter) Add(c context.Context, player entity.PlayerRecord) (entity.PlayerRecord, error) {
	playerSchema, err := r.service.Add(c, player)
	if err != nil {
		return entity.PlayerRecord{}, err
	}

	return playerSchema, nil
}

func (r *PlayerRepositoryAdapter) Find(c context.Context, id uint32) (entity.PlayerRecord, error) {
	playerSchema, e := r.service.Find(c, id)

	if e != nil {
		return entity.PlayerRecord{}, e
	}

	return playerSchema, nil
}

func (r *PlayerRepositoryAdapter) Modify(c context.Context, player entity.PlayerRecord) (entity.PlayerRecord, error) {
	playerSchema, e := r.service.Modify(c, player)

	if e != nil {
		return entity.PlayerRecord{}, e
	}

	return playerSchema, nil
}

func (r *PlayerRepositoryAdapter) Kill(c context.Context, id uint32) error {
	return r.service.Kill(c, id)
}
