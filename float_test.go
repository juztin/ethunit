package ethunit

import (
	"math/big"
	"testing"
)

type floatTest struct {
	unit     Unit
	value    *big.Int
	expected *big.Float
}

var floatTests = []floatTest{
	{Wei, big.NewInt(1), big.NewFloat(1)},
	{Kwei, big.NewInt(1000), big.NewFloat(1)},
	{Mwei, big.NewInt(1000000), big.NewFloat(1)},
	{Gwei, big.NewInt(1000000000), big.NewFloat(1)},
	{Szabo, big.NewInt(1000000000000), big.NewFloat(1)},
	{Finney, big.NewInt(1000000000000000), big.NewFloat(1)},
	{Ether, big.NewInt(1000000000000000000), big.NewFloat(1)},
}

func TestFloat(t *testing.T) {
	for _, test := range floatTests {
		result := Float(test.value, test.unit)
		cmp := result.Cmp(test.expected)
		if cmp != 0 {
			t.Errorf("Received %f(cmp %d) but expected %f", result, cmp, test.expected)
		}
	}
}
