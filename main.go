package main

import (
	db "Product/db/sqlc"
	"Product/util"
	"context"
	"os"

	"github.com/go-redis/redis"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Error().Err(err).Msg("app.env is not found")
		os.Exit(1)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to db")
		os.Exit(1)
	}

	initMigration(config.MigrationURL, config.DBSource)
	defer conn.Close()

	rds := redis.NewClient(
		&redis.Options{
			Addr:     config.RedisAddress,
			Password: "",
			DB:       2,
		},
	)

	store := db.NewStore(conn)
	// server, err := gapi.ServerSetup(config, store, rds)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("Failed to create server")
	// 	os.Exit(1)
	// }
}

func initMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create migration database source")
		os.Exit(1)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("Fail to Up migrate database")
		os.Exit(1)
	}

	log.Info().Msg("Successfully created migration database")
}
