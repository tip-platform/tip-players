// Package conn manages database connections and SQL helpers.
package conn

import (
	"database/sql"
	"fmt"
	"strings"
)

func ExecuteSQLFromFile(db *sql.DB, file string, params ...any) (any, error) {
	query, e := ReadSQLStringFromFile(file)

	if e == sql.ErrNoRows {
		return nil, fmt.Errorf("petition not searched exactitudes: %w", e)
	}

	if e != nil {
		return nil, fmt.Errorf("failed to read SQL file %s: %w", file, e)
	}

	if isSelectQuery(query) {
		return executeSelect(db, query, params...)
	}

	return executeNonSelect(db, query, params...)
}

func isSelectQuery(query string) bool {
	trimmed := strings.TrimSpace(query)
	trimmed = stripLeadingComments(trimmed)

	return strings.HasPrefix(strings.ToUpper(trimmed), "SELECT")
}

func stripLeadingComments(query string) string {
	trimmed := strings.TrimSpace(query)

	for strings.HasPrefix(trimmed, "--") {
		newlineIndex := strings.Index(trimmed, "\n")

		if newlineIndex <= 0 {
			return ""
		}

		trimmed = strings.TrimSpace(trimmed[newlineIndex+1:])
	}
	return trimmed
}

func executeSelect(db *sql.DB, query string, params ...any) (map[string]any, error) {
	rows, e := db.Query(query, params...)

	if e != nil {
		return nil, e
	}

	defer rows.Close()

	columns, e := rows.Columns()

	if e != nil {
		return nil, e
	}

	if !rows.Next() {
		if e := rows.Err(); e != nil {
			return nil, e
		}
		return nil, sql.ErrNoRows
	}

	values := make([]any, len(columns))
	valuePtrs := make([]any, len(columns))

	for i := range values {
		valuePtrs[i] = &values[i]
	}

	e = rows.Scan(valuePtrs...)

	if e != nil {
		return nil, e
	}

	result := make(map[string]any)

	for i, col := range columns {
		result[col] = values[i]
	}

	return result, nil
}

func executeNonSelect(db *sql.DB, query string, params ...any) (sql.Result, error) {
	return db.Exec(query, params...)
}
