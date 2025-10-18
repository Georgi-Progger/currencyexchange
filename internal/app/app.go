package app

import (
	"currencyexchange/config"
	"currencyexchange/internal/handler"
	"currencyexchange/internal/repo"
	"currencyexchange/internal/usecase"
	"currencyexchange/pkg/dbsource"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func Run() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	db, err := dbsource.NewDb(cfg)
	if err != nil {
		slog.Error("error")
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	repos := repo.NewRepository(db)
	usecases := usecase.NewUsecase(repos)
	handlers := handler.NewHandler(usecases)

	router := handlers.SetupRouter(db.DB)

	slog.Info("Starting server", "port", cfg.Server.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Server.Port), router); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
