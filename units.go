package ethunit

import (
	"fmt"
	"math"
	"math/big"
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
	Kether                 // 21
	Mether                 // 24
	Gether                 // 27
	Tether                 // 30
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
	case "kether":
		u = Kether
	case "mether":
		u = Mether
	case "gether":
		u = Gether
	case "tether":
		u = Tether
	default:
		err = fmt.Errorf("Invalid unit type '%s'", s)
	}
	return u, err
}

func ConvertInt(from *big.Int, fromUnit, toUnit Unit) *big.Float {
	to := new(big.Float)
	to.SetPrec(236) // IEEE 754 octuple-precision binary floating-point format: binary256
	to.SetMode(big.ToNearestEven)
	to.SetFloat64(math.Pow10(int(fromUnit - toUnit)))
	return to.Mul(new(big.Float).SetInt(from), to)
}

func ConvertFloat(from *big.Float, fromUnit, toUnit Unit) *big.Float {
	to := new(big.Float)
	to.SetPrec(236) // IEEE 754 octuple-precision binary floating-point format: binary256
	to.SetMode(big.ToNearestEven)
	to.SetFloat64(math.Pow10(int(fromUnit - toUnit)))
	return to.Mul(from, to)
}
