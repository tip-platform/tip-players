// Package persistence provides database-backed player adapters.
package persistence

import (
	"context"

	"github.com/tip-platform/tip-players/internal/app/port/output"
	"github.com/tip-platform/tip-players/internal/domain/schema"
)

// MSSQLPlayerRepositoryAdapter implements the output.PlayerRepository port using MSSQL.
type MSSQLPlayerRepositoryAdapter struct {
	store *PlayerStore
}

func NewMSSQLPlayerRepositoryAdapter(store *PlayerStore) *MSSQLPlayerRepositoryAdapter {
	return &MSSQLPlayerRepositoryAdapter{store: store}
}

func (r *MSSQLPlayerRepositoryAdapter) Insert(ctx context.Context, player schema.Player) error {
	return r.store.Insert(ctx, player)
}

func (r *MSSQLPlayerRepositoryAdapter) Select(ctx context.Context, id uint32) (schema.Player, error) {
	return r.store.Select(ctx, id)
}

func (r *MSSQLPlayerRepositoryAdapter) Update(ctx context.Context, player schema.Player) error {
	return r.store.Update(ctx, player)
}

func (r *MSSQLPlayerRepositoryAdapter) Delete(ctx context.Context, id uint32) error {
	return r.store.Delete(ctx, id)
}

var _ output.PlayerRepository = (*MSSQLPlayerRepositoryAdapter)(nil)
