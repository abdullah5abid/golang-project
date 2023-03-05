package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spongeling/admin-api/internal/api"
	"github.com/spongeling/admin-api/internal/auth"
	"github.com/spongeling/admin-api/internal/repo"
	"github.com/spongeling/admin-api/internal/server"
	"github.com/spongeling/admin-api/internal/service"
	"github.com/spongeling/admin-api/shared"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// config
	err := shared.LoadConfig(".env")
	if err != nil {
		log.Fatalf("error loading config %v", err)
	}
	cfg, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	// database
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)

	dbCfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal(err)
	}

	dbCfg.ConnConfig.LogLevel = pgx.LogLevelError

	conn, err := pgxpool.ConnectConfig(ctx, dbCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// migrate database
	if err = migrateUp(cfg.DbMigrationsSource, connStr); err != nil {
		log.Println("Database migrations failed...")
	}

	// service
	db := repo.New(conn)

	svc := service.New(db)
	auth.Users, err = svc.GetAllUsers(ctx)

	// apis
	pingApi := api.NewPing(svc)
	posApi := api.NewCorpus(svc)
	loginApi := api.NewLogin(svc)
	wordClassApi := api.NewWordClass(svc)
	categoryApi := api.NewCategory(svc)

	// server
	srv := server.New(cfg.HttpPort,
		pingApi,
		posApi,
		loginApi,
		categoryApi,
		wordClassApi,
	)

	err = srv.Start()
	if err != nil {
		log.Fatalf("start server err: %v", err)
	}
}
