// Package input defines player service interfaces.
package input

import (
	"context"

	"github.com/tip-platform/tip-players/internal/domain/schema"
)

type PlayerService interface {
	Add(ctx context.Context, player schema.Player) (schema.Player, error)

	Find(ctx context.Context, id uint32) (schema.Player, error)

	Modify(ctx context.Context, player schema.Player) (schema.Player, error)

	Kill(ctx context.Context, id uint32) error
}
