package app

import (
	"currencyexchange/internal/handler"
	"currencyexchange/internal/repo"
	"currencyexchange/internal/usecase"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Run() {
	db, err := dbConnect()
	if err != nil {
		slog.Error("error")
	}
	defer db.Close()

	repos := repo.NewRepository(db)
	usecases := usecase.NewUsecase(repos)
	handlers := handler.NewHandler(usecases)

	router := handlers.SetupRouter(db.DB)

	port := ":8080"
	slog.Info("Starting server", "port", port)

	if err := http.ListenAndServe(port, router); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}

func dbConnect() (*sqlx.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	portDb := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		user, password, dbName, host, portDb)

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
