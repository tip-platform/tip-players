// Package entity defines persistence layer player records.
package entity

import (
	"time"
)

type PlayerRecord struct {
	ID          uint32    `db:"id"`
	APIID       uint32    `db:"api_id"`
	Name        string    `db:"name"`
	ShortName   string    `db:"short_name"`
	CountryCode string    `db:"country_code"`
	CountryName string    `db:"country_name"`
	Age         uint8     `db:"age"`
	Plays       string    `db:"plays"`
	TurnedPro   time.Time `db:"turned_pro"`
}
