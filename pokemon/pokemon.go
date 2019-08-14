package pokemon

import (
	"database/sql"
	"errors"
	"fmt"

	pkgErrors "github.com/pkg/errors"

	"github.com/rs/zerolog/log"

	"github.com/AndersonQ/go1_13_errors/db"
	"github.com/AndersonQ/go1_13_errors/myerrors"
)

var Cheat int

type Escaped struct {
	Pokemon Pokemon
}

type NoPokeballs struct {
	hasDeal bool
}

func (c Escaped) Error() string {
	return fmt.Sprintf("%s escaped", c.Pokemon.Name)
}

func (c Escaped) String() string {
	return c.Pokemon.String()
}

func (n NoPokeballs) Error() string {
	return fmt.Sprintf("no pokeballs")
}

// Deal might be offered to a players when Catch returns NoPokeballs.
// Call Deal() to trigger the deal. If the player accepts the deal, Deal returns true, false otherwise
func (n NoPokeballs) Deal() bool {
	if n.hasDeal {
		log.Info().Msg("Want to buy X Pokeballs for xx,xx?")
		return true
	}

	return false
}

func Catch() error {
	switch Cheat {
	case 0:
		return myerrors.Wrap(Escaped{Pokemon: Pokemon{
			Number: 25,
			Name:   "Pikachu",
			Type:   "eletric",
		}}, "wrapped escaped with myerrors")
	case 1:
		return pkgErrors.Wrap(NoPokeballs{hasDeal: true}, "wrapped escaped with pkgErrors")
	case 2:
		return NoPokeballs{}
	}

	return errors.New("internal server error")
}

var NotFound = errors.New("not found")

func Find(id int) (*Pokemon, error) {
	p, err := db.Find(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NotFound
		}
		return nil, myerrors.Wrap(err, "problem looking for pokemon #%d", id)
	}

	return &Pokemon{
		Number: id,
		Name:   p.Name,
		Type:   p.Type,
	}, nil
}
