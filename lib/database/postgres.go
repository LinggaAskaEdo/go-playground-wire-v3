package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/config"
)

type PostgresImpl struct {
	SQL *sqlx.DB
}

func NewPostgresClient() *PostgresImpl {
	log.Info().Msg("Initialize Postgres connection...")
	var err error

	dbHost := config.Get().Database.Postgres.Host
	dbPort := config.Get().Database.Postgres.Port
	dbName := config.Get().Database.Postgres.Name
	dbUser := config.Get().Database.Postgres.User
	dbPass := config.Get().Database.Postgres.Password

	sHost := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("postgres", sHost)
	if err != nil {
		log.Err(err).Msgf("Error to loading Database %s", err)
		panic(err)
	}

	log.Info().Str("name", dbName).Msg("Success connect to database -->")
	return &PostgresImpl{
		SQL: db,
	}
}
