package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/vnnyx/betty-BE/internal/config"
	"github.com/vnnyx/betty-BE/internal/log"
)

func NewCockroachDatabase(conf *config.Config) (*sql.DB, error) {
	logger := log.NewLog()

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db, err := sql.Open("postgres", conf.Database.Cockroach.DSN)
	if err != nil {
		logger.Fatal().Caller().Err(err).Msg("Error connecting to database")
	}

	if err := db.PingContext(ctx); err != nil {
		logger.Fatal().Caller().Err(err).Msg("Error connecting to database")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db, nil
}

func NewCockroachContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func RunMigration(cfg *config.Config) error {
	logger := log.NewLog()
	db, err := NewCockroachDatabase(cfg)
	if err != nil {
		return err
	}

	driver, err := cockroachdb.WithInstance(db, &cockroachdb.Config{})
	if err != nil {
		logger.Fatal().Caller().Err(err).Msg("Error running migration")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"cockroachdb", driver)
	if err != nil {
		logger.Fatal().Caller().Err(err).Msg("Error running migration")
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.Fatal().Caller().Err(err).Msg("Error running migration")
	}

	return nil
}
