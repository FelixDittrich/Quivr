package main

import (
	"github.com/FelixDittrich/Quivr/internal/db"
	"github.com/FelixDittrich/Quivr/internal/env"
	"github.com/FelixDittrich/Quivr/internal/store"
	"go.uber.org/zap"
)

const version = "0.0.1"

func main() {
	env.LoadEnvFile("./.env")

	// Config
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),
		env: env.GetString("ENV", "development"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Database
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store: store,
		logger: logger,
	}

	mux := app.mount()
	logger.Fatal(app.run(mux))
}