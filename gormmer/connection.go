package gormmer

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseOption struct {
	Driver          string
	Host            string
	Port            int
	SSLMode         bool
	DBName          string
	Username        string
	Password        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func NewConnection(cfg DatabaseOption) *gorm.DB {

	// set default configuration
	if cfg.Host == "" {
		cfg.Host = "localhost"
	}

	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = 5
	}

	if cfg.MaxOpenConns == 0 {
		cfg.MaxOpenConns = 10
	}

	if cfg.ConnMaxLifetime == 0 {
		cfg.ConnMaxLifetime = time.Hour
	}

	switch dbDriver := cfg.Driver; dbDriver {
	case "postgresql":
		return newPostgreSQL(cfg)
	case "pgx":
		return newPostgreSQL(cfg)
	default:
		return newSQLLite(cfg)
	}
}

func newPostgreSQL(config DatabaseOption) *gorm.DB {
	// set default config for postgresql
	if config.Port == 0 {
		config.Port = 5432
	}

	var sslMode = "disable"
	if !config.SSLMode {
		sslMode = "require"
	}

	// generate postgresql DSN
	dsnPostgreSQL := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, sslMode)

	connection, err := sql.Open(config.Driver, dsnPostgreSQL)
	if err != nil {
		log.Fatal(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	connection.SetMaxIdleConns(config.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	connection.SetMaxOpenConns(config.MaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	connection.SetConnMaxLifetime(config.ConnMaxLifetime)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: connection,
	}), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func newSQLLite(config DatabaseOption) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.DBName), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// TODO: support for mySQL
