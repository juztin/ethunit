package ethunit

import (
	"fmt"
	"strings"
)

type Unit int

const (
	Wei    Unit = iota * 3 // 0
	Kwei                   // 3
	Mwei                   // 6
	Gwei                   // 9
	Szabo                  // 12
	Finney                 // 15
	Ether                  // 18
)

func Atou(s string) (Unit, error) {
	var err error
	var u Unit
	switch strings.ToLower(s) {
	case "wei":
		u = Wei
	case "kwei":
		u = Kwei
	case "mwei":
		u = Mwei
	case "gwei":
		u = Gwei
	case "szabo":
		u = Szabo
	case "finney":
		u = Finney
	case "ether":
		u = Ether
	default:
		err = fmt.Errorf("Invalid unit type '%s'", s)
	}
	return u, err
}
