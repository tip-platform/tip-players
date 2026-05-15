// Package persistence provides database-backed player adapters.
package persistence

import (
	"context"

	"github.com/tip-platform/tip-players/internal/domain/entity"
	"github.com/tip-platform/tip-players/internal/port/output"
)

type PlayerRepositoryAdapter struct {
	store *PlayerStore
}

func NewPlayerRepositoryAdapter(store *PlayerStore) *PlayerRepositoryAdapter {
	return &PlayerRepositoryAdapter{store: store}
}

func (r *PlayerRepositoryAdapter) Insert(ctx context.Context, player entity.PlayerRecord) error {
	return r.store.Insert(ctx, player)
}

func (r *PlayerRepositoryAdapter) Select(ctx context.Context, id uint32) (entity.PlayerRecord, error) {
	return r.store.Select(ctx, id)
}

func (r *PlayerRepositoryAdapter) Update(ctx context.Context, player entity.PlayerRecord) error {
	return r.store.Update(ctx, player)
}

func (r *PlayerRepositoryAdapter) Delete(ctx context.Context, id uint32) error {
	return r.store.Delete(ctx, id)
}

var _ output.PlayerRepository = (*PlayerRepositoryAdapter)(nil)
