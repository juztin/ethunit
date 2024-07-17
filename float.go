package ethunit

import (
	"math"
	"math/big"
)

var (
	wei    = new(big.Float)
	kwei   = big.NewFloat(math.Pow10(int(Kwei)))
	mwei   = big.NewFloat(math.Pow10(int(Mwei)))
	gwei   = big.NewFloat(math.Pow10(int(Gwei)))
	szabo  = big.NewFloat(math.Pow10(int(Szabo)))
	finney = big.NewFloat(math.Pow10(int(Finney)))
	ether  = big.NewFloat(math.Pow10(int(Ether)))
)

func pow10(u Unit) *big.Float {
	switch u {
	case Kwei:
		return kwei
	case Mwei:
		return mwei
	case Gwei:
		return gwei
	case Szabo:
		return szabo
	case Finney:
		return finney
	case Ether:
		return ether
	}
	return wei
}

func Float(i *big.Int, u Unit) *big.Float {
	if u == Wei {
		return new(big.Float).SetInt(i)
	}
	f := new(big.Float).
		SetMode(big.ToNearestEven).
		SetPrec(236).
		SetInt(i)
	return f.Quo(f, pow10(u))
}
