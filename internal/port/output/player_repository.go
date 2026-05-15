// Package output defines player repository interfaces.
package output

import (
	"context"

	"github.com/tip-platform/tip-players/internal/domain/entity"
)

type PlayerRepository interface {
	Insert(ctx context.Context, player entity.PlayerRecord) error

	Select(ctx context.Context, id uint32) (entity.PlayerRecord, error)

	Update(ctx context.Context, player entity.PlayerRecord) error

	Delete(ctx context.Context, id uint32) error
}
