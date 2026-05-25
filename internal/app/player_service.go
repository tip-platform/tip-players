// Package app implements player application services.
package app

import (
	"context"

	"github.com/tip-platform/tip-players/internal/app/port/input"
	"github.com/tip-platform/tip-players/internal/app/port/output"
	"github.com/tip-platform/tip-players/internal/domain/schema"
)

type PlayerService struct {
	repo output.PlayerRepository
}

func NewPlayerService(repo output.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) Add(c context.Context, player schema.Player) (schema.Player, error) {
	p, e := schema.NewPlayer(
		player.APIID,
		player.Name,
		schema.WithCountryCode(player.CountryCode),
		schema.WithCountryName(player.CountryName),
		schema.WithAge(player.Age),
	)

	p.ShortName = player.ShortName
	p.TurnedPro = player.TurnedPro

	if e != nil {
		return schema.Player{}, e
	}

	if e := s.repo.Insert(c, p); e != nil {
		return schema.Player{}, e
	}

	return p, nil
}

func (s *PlayerService) Find(c context.Context, id uint32) (schema.Player, error) {
	p, e := s.repo.Select(c, id)

	if e != nil {
		return schema.Player{}, e
	}

	return p, nil
}

func (s *PlayerService) Modify(c context.Context, player schema.Player) (schema.Player, error) {
	p, e := schema.NewPlayer(
		player.APIID,
		player.Name,
		schema.WithCountryCode(player.CountryCode),
		schema.WithCountryName(player.CountryName),
		schema.WithAge(player.Age),
	)

	p.ShortName = player.ShortName
	p.TurnedPro = player.TurnedPro

	if e != nil {
		return schema.Player{}, e
	}

	if e = s.repo.Update(c, p); e != nil {
		return schema.Player{}, e
	}

	return p, nil
}

func (s *PlayerService) Kill(c context.Context, id uint32) error {
	return s.repo.Delete(c, id)
}

var _ input.PlayerService = (*PlayerService)(nil)
