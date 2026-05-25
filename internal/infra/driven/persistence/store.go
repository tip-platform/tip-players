// Package persistence provides database-backed player adapters.
package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tip-platform/tip-players/internal/domain/schema"
	"github.com/tip-platform/tip-players/internal/infra/driven/persistence/conn"
)

type PlayerStore struct {
	db *sql.DB
}

func NewPlayerStore() (*PlayerStore, error) {
	db, err := conn.NewClient()
	if err != nil {
		return nil, err
	}

	return &PlayerStore{db: db}, nil
}

func (s *PlayerStore) Ping(c context.Context) error {
	return s.db.PingContext(c)
}

func (s *PlayerStore) Insert(c context.Context, r schema.Player) error {
	_, e := conn.ExecuteSQLFromFile(
		s.db,
		"INSERT_PLAYER.sql",
		r.APIID,
		r.Name,
		r.ShortName,
		r.CountryCode,
		r.CountryName,
		r.Age,
		r.Plays,
		r.TurnedPro,
	)

	if e != nil {
		return e
	}

	return nil
}

func (s *PlayerStore) Select(c context.Context, id uint32) (schema.Player, error) {
	var r schema.Player

	query, e := conn.ReadSQLStringFromFile("SELECT_PLAYER_BY_ID.sql")

	if e != nil {
		return schema.Player{}, fmt.Errorf("failed to read SQL file: %w", e)
	}

	row := s.db.QueryRow(query, id)

	e = row.Scan(
		&r.ID,
		&r.APIID,
		&r.Name,
		&r.ShortName,
		&r.CountryCode,
		&r.CountryName,
		&r.Age,
		&r.Plays,
		&r.TurnedPro,
	)

	if e != nil {
		return schema.Player{}, e
	}

	return r, nil
}

func (s *PlayerStore) Update(c context.Context, r schema.Player) error {
	_, e := conn.ExecuteSQLFromFile(
		s.db,
		"UPDATE_PLAYER.sql",
		r.APIID,
		r.Name,
		r.ShortName,
		r.CountryCode,
		r.CountryName,
		r.Age,
		r.Plays,
		r.TurnedPro,
		r.ID,
	)

	if e != nil {
		return e
	}

	return nil
}

func (s *PlayerStore) Delete(c context.Context, id uint32) error {
	query := "DELETE FROM players WHERE id = @p1"

	_, e := s.db.Exec(query, id)

	if e != nil {
		return e
	}

	return nil
}
