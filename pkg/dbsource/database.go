package dbsource

import (
	"currencyexchange/config"
	"fmt"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
)

func NewDb(cfg config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.DbName, cfg.Database.Host, cfg.Database.PortDb)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		slog.Error("Failed to ping database", "error", err)
		os.Exit(1)
	}

	return db, nil
}
