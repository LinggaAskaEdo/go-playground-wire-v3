package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/config"
)

type MysqlImpl struct {
	SQL *sqlx.DB
}

func NewMysqlClient() *MysqlImpl {
	log.Info().Msg("Initialize MySQL connection...")
	var err error

	dbHost := config.Get().Database.MySQL.Host
	dbPort := config.Get().Database.MySQL.Port
	dbName := config.Get().Database.MySQL.Name
	dbUser := config.Get().Database.MySQL.User
	dbPass := config.Get().Database.MySQL.Password

	sHost := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("mysql", sHost)
	if err != nil {
		log.Err(err).Msgf("Error to loading Database %s", err)
		panic(err)
	}

	log.Info().Str("name", dbName).Msg("Success connect to database -->")
	return &MysqlImpl{
		SQL: db,
	}
}
