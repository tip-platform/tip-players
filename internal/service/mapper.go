// Package service implements player application services.
package service

import (
	"github.com/tip-platform/tip-players/internal/domain/entity"
	"github.com/tip-platform/tip-players/internal/domain/schema"
)

func SchemaToEntity(p schema.Player) entity.PlayerRecord {
	return entity.PlayerRecord{
		ID:          p.ID,
		APIID:       p.APIID,
		Name:        p.Name,
		ShortName:   p.ShortName,
		CountryCode: p.CountryCode,
		CountryName: p.CountryName,
		Age:         p.Age,
		Plays:       p.Plays,
		TurnedPro:   p.TurnedPro,
	}
}

func EntityToSchema(r entity.PlayerRecord) schema.Player {
	return schema.Player{
		ID:          r.ID,
		APIID:       r.APIID,
		Name:        r.Name,
		ShortName:   r.ShortName,
		CountryCode: r.CountryCode,
		CountryName: r.CountryName,
		Age:         r.Age,
		Plays:       r.Plays,
		TurnedPro:   r.TurnedPro,
	}
}
