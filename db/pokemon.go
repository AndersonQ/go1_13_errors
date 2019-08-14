package db

import (
	"database/sql"

	"github.com/pkg/errors"
)

type Pokemon struct {
	Name string `db:"name"`
	Type string `db:"type"`
}

// Find returns:
// details for the given pokemon id or
// sql.ErrNoRows if there is no register of the given pokemon id in your pokedex
// error for id < 1
func Find(id int) (*Pokemon, error) {
	if id < 1 {
		return nil, errors.New("there is no pokemon with number <= 0")
	}

	if id == 25 {
		return &Pokemon{Name: "Pikachu", Type: "Eletric"}, nil
	}

	return nil, sql.ErrNoRows
}

// // BetterFind returns:
// // details for the given pokemon id or
// // NotFound if there is no register of the given pokemon id in your pokedex
// // error for id < 1
// func BetterFind(id int) (pokemon.Pokemon, error) {
// 	if id < 1 {
// 		return pokemon.Pokemon{}, errors.New("There is no pokemon with number <= 0!")
// 	}
//
// 	if id == 25 {
// 		return pokemon.Pokemon{Number: "25", Name: "Pikachu", Type: "Eletric"}, nil
// 	}
//
// 	err := someOperation()
// 	return pokemon.Pokemon{}, errors.Wrapf(NotFound, "pokemon #%d not found", id)
// }
//
// func someOperation() error {
// 	return errors.Wrap(sql.ErrNoRows, "some message")
// }
