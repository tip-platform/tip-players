// Package output defines player repository interfaces.
package output

import (
	"context"

	"github.com/tip-platform/tip-players/internal/domain/schema"
)

type PlayerRepository interface {
	Insert(ctx context.Context, player schema.Player) error

	Select(ctx context.Context, id uint32) (schema.Player, error)

	Update(ctx context.Context, player schema.Player) error

	Delete(ctx context.Context, id uint32) error
}
