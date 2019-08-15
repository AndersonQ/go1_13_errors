package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	pkgErr "github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/AndersonQ/go1_13_errors/db"
	"github.com/AndersonQ/go1_13_errors/myerrors"
	"github.com/AndersonQ/go1_13_errors/pokemon"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

func main() {
	log.Info().Msg("-----Is:-----")

	id := 1
	_, err := db.Find(id)
	handle113Error(err, id)
	log.Info().Msg("----------------------------------------")

	id = 2
	_, err = pokemon.Find(id)
	handle113Error(err, id)
	log.Info().Msg("----------------------------------------")

	id = 3
	_, err = pokemon.Find(id)
	handle113Error(pkgErr.Wrap(myerrors.Wrap(err, "myerror wrapping"), "more wrapping"), id)
	log.Info().Msg("----------------------------------------")

	log.Info().Msg("----------------------------------------")
	log.Info().Msg("-----As:-----")

	pokemon.Cheat = 0
	handleCatch(pokemon.Catch())
	log.Info().Msg("----------------------------------------")
	pokemon.Cheat = 1
	handleCatch(pokemon.Catch())
	log.Info().Msg("----------------------------------------")
	pokemon.Cheat = 2
	handleCatch(pokemon.Catch())
	log.Info().Msg("----------------------------------------")
	pokemon.Cheat = -1
	handleCatch(pokemon.Catch())
	log.Info().Msg("----------------------------------------")

	log.Info().Msg("----------------------------------------")
	log.Info().Msg("-----fmt.Errorf:-----")

	err = fmt.Errorf("an error")
	err = myerrors.Wrap(err, "1s wrapping error with myerrors.Wrap")
	err = fmt.Errorf("2nd wrapping with fmt.Errorf: %w", err)
	err = pkgErr.Wrapf(err, "3rd wrapping with pkg/errors", err)

	log.Info().Msgf("%T: err: %s", err, err)

	if err = errors.Unwrap(err); err != nil {
		log.Info().Msgf("1st errors.Unwrap(err): %T: %s", err, err.Error())
	}
	if err = errors.Unwrap(err); err != nil {
		log.Info().Msgf("2nd errors.Unwrap(err): %T: %s", err, err.Error())
	}
	if err = errors.Unwrap(err); err != nil {
		log.Info().Msgf("3rd errors.Unwrap(err): %T: %s", err, err.Error())
	}

	var k Break
	log.Debug().Msgf("%v", k)
	errors.As(errors.New(""), k)
}

func handle113Error(err error, id int) {
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Debug().Msgf("#%d: It's a sql.ErrNoRows error", id)
			log.Error().Err(err).Msgf("Pokemon #%d not registered in your pokedex", id)
			return
		}

		if errors.Is(err, pokemon.NotFound) {
			log.Debug().Msgf("#%d: It's a pokemon.NotFound error", id)
			log.Error().Err(err).Msgf("Pokemon #%d not registered in your pokedex", id)
			return
		}

		log.Error().Err(err).Msgf("generic error finding pokemon #%d", id)
	}
}

func handleCatch(err error) {
	if err != nil {
		var pe pokemon.Escaped
		if errors.As(err, &pe) {
			log.Info().Msgf("%v: %s", pe.Pokemon, pe.Error())
			return
		}

		var np pokemon.NoPokeballs
		if errors.As(err, &np) {
			if !np.Deal() {
				log.Info().Msg("You have no pokeballs, go to a shop to buy more")
			}
			return
		}

		log.Error().Err(err).Msg("problem trying to catch pokemon")
	}
}

type Break interface{}
