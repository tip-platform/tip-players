// Package conn manages database connections and SQL helpers.
package conn

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

func pingWithRetry(db *sql.DB, attempts int, delay time.Duration) error {
	var lastErr error

	for i := 0; i < attempts; i++ {
		if e := db.Ping(); e == nil {
			return nil
		} else {
			lastErr = e
			time.Sleep(delay)
		}
	}

	return lastErr
}

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

// buildConnString ahora acepta el nombre de la base de datos dinámicamente
func buildConnString(config dbConfig, dbname string) string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", config.user, config.password, config.host, config.port, dbname)
}

func ensureDatabaseAndTables(config dbConfig) (*sql.DB, error) {
	// 1. Conectamos inicialmente a 'master' para verificar/crear la base de datos del proyecto
	masterConnString := buildConnString(config, "master")

	masterDB, e := sql.Open("sqlserver", masterConnString)

	if e != nil {
		return nil, fmt.Errorf("failed to connect to master: %w", e)
	}

	// Usamos un flag manual o cerramos masterDB explícitamente en lugar de defer defer,
	// para asegurarnos de que libere la conexión antes de que la app intente migrar las tablas.
	if e := pingWithRetry(masterDB, 12, 2*time.Second); e != nil {
		masterDB.Close()
		return nil, fmt.Errorf("failed to ping master: %w", e)
	}

	// Crear la base de datos (Ejecuta en el contexto de 'master')
	if _, e := ExecuteSQLFromFile(masterDB, "CREATE_DATABASE.sql", config.dbname); e != nil {
		masterDB.Close()
		return nil, fmt.Errorf("failed to create database: %w", e)
	}
	masterDB.Close() // Ya no necesitamos la sesión master, cerramos de forma limpia

	// 2. Ahora abrimos una conexión enfocada DIRECTAMENTE en la base de datos de la app para las tablas
	appConnString := buildConnString(config, config.dbname)

	appDB, e := sql.Open("sqlserver", appConnString)

	if e != nil {
		return nil, fmt.Errorf("failed to connect to app database: %w", e)
	}

	if e := pingWithRetry(appDB, 12, 2*time.Second); e != nil {
		appDB.Close()
		return nil, fmt.Errorf("failed to ping app database: %w", e)
	}

	// Crear las tablas (Ejecuta ya posicionado en tu base de datos destino)
	if _, e := ExecuteSQLFromFile(appDB, "CREATE_TABLES.sql"); e != nil {
		appDB.Close()
		return nil, fmt.Errorf("failed to create tables: %w", e)
	}

	// 3. Devolvemos el pool 'appDB' que está VIVO, listo y conectado a la DB final
	return appDB, nil
}

func NewClient() (*sql.DB, error) {
	config, e := readDBConfigFromEnv()
	if e != nil {
		return nil, e
	}

	// Pasamos toda la configuración a la función encargada de inicializar
	return ensureDatabaseAndTables(config)
}

func (c *Client) Close() error {
	return c.DB.Close()
}
