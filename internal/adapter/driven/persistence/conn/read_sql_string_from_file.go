// Package conn manages database connections and SQL helpers.
package conn

import (
	"os"
	"path/filepath"
)

func ReadSQLStringFromFile(filename string) (string, error) {
	path := filepath.Join("internal/adapter/driven/persistence/sql", filename)

	content, err := os.ReadFile(path)

	return string(content), err
}
