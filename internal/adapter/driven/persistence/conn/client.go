// Package conn manages database connections and SQL helpers.
package conn

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

type Client struct {
	DB *sql.DB
}

type dbConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func readDBConfigFromEnv() (dbConfig, error) {
	config := dbConfig{
		host:     os.Getenv("DB_PJ_HOST"),
		port:     os.Getenv("DB_PJ_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("SQL_PASSWORD"),
		dbname:   os.Getenv("DB_PJ_NAME"),
	}

	if config.host == "" || config.port == "" || config.user == "" || config.password == "" || config.dbname == "" {
		return dbConfig{}, fmt.Errorf("missing required environment variables")
	}

	return config, nil
}

func buildConnString(config dbConfig) string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", config.user, config.password, config.host, config.port, config.dbname)
}

func ensureDatabaseAndTables(connString, dbname string) (*sql.DB, error) {
	masterDB, e := sql.Open("sqlserver", connString)

	if e != nil {
		return nil, fmt.Errorf("failed to connect to master: %w", e)
	}

	defer masterDB.Close()

	if e := masterDB.Ping(); e != nil {
		return nil, fmt.Errorf("failed to ping master: %w", e)
	}

	if _, e := ExecuteSQLFromFile(masterDB, "CREATE_DATABASE.sql", dbname); e != nil {
		return nil, fmt.Errorf("failed to create database: %w", e)
	}

	if _, e := ExecuteSQLFromFile(masterDB, "CREATE_TABLES.sql"); e != nil {
		return nil, fmt.Errorf("failed to create table: %w", e)
	}

	return masterDB, nil
}

func NewClient() (*sql.DB, error) {
	config, e := readDBConfigFromEnv()

	if e != nil {
		return nil, e
	}

	connString := buildConnString(config)

	return ensureDatabaseAndTables(connString, config.dbname)
}

func (c *Client) Close() error {
	return c.DB.Close()
}
