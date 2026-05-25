// Package conn manages database connections and SQL helpers.
package conn

import (
	"os"
	"path/filepath"
)

func ReadSQLStringFromFile(filename string) (string, error) {
	// Prefer resolving from the executable directory (works in containers and multi-env).
	if exe, err := os.Executable(); err == nil {
		base := filepath.Dir(exe)

		abs := filepath.Join(base, "internal/adapter/driven/persistence/sql", filename)

		if b, e := os.ReadFile(abs); e == nil {
			return string(b), nil
		}
	}

	// Fallback to relative path (local dev).
	rel := filepath.Join("internal/adapter/driven/persistence/sql", filename)

	b, err := os.ReadFile(rel)

	return string(b), err
}
