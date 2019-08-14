package pokemon

import "fmt"

type Pokemon struct {
	Number int
	Name   string
	Type   string
}

func (p Pokemon) String() string {
	return fmt.Sprintf("#%d %s (%s)", p.Number, p.Name, p.Type)
}
