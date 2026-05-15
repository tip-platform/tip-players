// Package schema defines player domain models and validation.
package schema

import (
	"strings"
	"time"

	"github.com/tip-platform/tip-players/internal/domain/entity"
)

// Player is the domain model exposed by the schema layer.
//
// NOTE: It's an alias to the persistence/entity representation so ports can use
// schema.Player without duplicating structs.
type Player = entity.PlayerRecord

func WithShortName(shortName string) Option {
	return func(p *entity.PlayerRecord) error {
		if p.ShortName != "" {
			return DomainError{Operation: "create", Entity: "player", Field: "short_name", Reason: ErrEmpty}
		}
		if len(shortName) > len(p.Name) {
			return DomainError{Operation: "create", Entity: "player", Field: "short_name", Reason: ErrTooLong}
		}
		p.ShortName = shortName
		return nil
	}
}

func WithCountryName(countryName string) Option {
	return func(p *entity.PlayerRecord) error {
		if countryName == "" {
			return DomainError{Operation: "create", Entity: "player", Field: "country_name", Reason: ErrEmpty}
		}
		p.CountryName = countryName
		return nil
	}
}

func WithCountryCode(code string) Option {
	return func(p *entity.PlayerRecord) error {
		if len(code) < 2 || len(code) > 3 || code != strings.ToUpper(code) {
			return DomainError{Operation: "create", Entity: "player", Field: "country_code", Reason: ErrInvalidFormat}
		}
		p.CountryCode = code
		return nil
	}
}

func WithAge(age uint8) Option {
	return func(p *entity.PlayerRecord) error {
		if age <= 0 || age > 120 {
			return DomainError{Operation: "create", Entity: "player", Field: "age", Reason: ErrMustBePositive}
		}
		p.Age = age
		return nil
	}
}

func WithTurnedPro(turnedPro time.Time) Option {
	return func(p *entity.PlayerRecord) error {
		if turnedPro.After(time.Now().UTC()) {
			return DomainError{Operation: "create", Entity: "player", Field: "turned_pro", Reason: ErrMustBePast}
		}
		p.TurnedPro = turnedPro
		return nil
	}
}

func NewPlayer(apiID uint32, name string, opts ...Option) (entity.PlayerRecord, error) {
	if name == "" {
		return entity.PlayerRecord{}, DomainError{Operation: "update", Entity: "player", Field: "name", Reason: ErrEmpty}
	}

	p := entity.PlayerRecord{
		APIID:     apiID,
		Name:      name,
		TurnedPro: time.Now().UTC(),
	}

	for _, opt := range opts {
		if err := opt(&p); err != nil {
			return entity.PlayerRecord{}, err
		}
	}
	return p, nil
}

type Option func(*entity.PlayerRecord) error
