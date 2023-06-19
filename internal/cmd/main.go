package main

import (
	"cleara/config"
	"cleara/internal/core/services/user_service"
	"cleara/internal/handlers/users"
	"cleara/internal/repositories/usersrepo"
	"cleara/internal/server"
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	var db, err = openDB(ctx)
	if err != nil {
		logger.Fatal(err)
	}

	userRepository := usersrepo.NewRepository(db, ctx)
	usersService := user_service.New(userRepository)
	usersHandler := users.NewUserHandlers(usersService)

	// Server
	httpServer := server.NewServer(
		usersHandler,
	)
	httpServer.Initialize()
}

func openDB(ctx context.Context) (*sqlx.DB, error) {
	var cfg, err = config.Load()
	if err != nil {
		return nil, err
	}

	db, err := sqlx.ConnectContext(ctx, "pgx", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	// seteamos el numero maximo de conexiones abiertas. 0 indica sin limite
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	// seteamos el numero maximo de conexiones inactivas. 0 indica sin limite
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	// usamos time.ParseDuration para convertir el string de duracion a time.Duration
	duration, err := time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	// Seteamos el timeout para las inactivas
	db.SetConnMaxIdleTime(duration)

	// creamos el contexto con 5 segundos de timeout deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// PingContext
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
