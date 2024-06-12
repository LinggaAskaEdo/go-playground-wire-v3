package database

import (
	"github.com/rs/zerolog/log"

	scribble "github.com/nanobox-io/golang-scribble"
)

type ScribleImpl struct {
	db *scribble.Driver
}

func NewScribleClient() *ScribleImpl {
	log.Info().Msg("Initialize Scribble connection...")

	db, err := scribble.New("tmp/db", nil)
	if err != nil {
		log.Err(err).Msg("NewScribleClient")
	}

	log.Info().Msg("Success, scribble is ready")

	return &ScribleImpl{
		db: db,
	}
}

func (db ScribleImpl) DB() *scribble.Driver {
	return db.db
}
