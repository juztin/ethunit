package ethunit

import (
	"math/big"
	"testing"
)

type stringTest struct {
	unit     Unit
	value    *big.Int
	expected string
}

func newBigInt(s string) *big.Int {
	i, _ := new(big.Int).SetString(s, 10)
	return i
}

var stringTests = []stringTest{
	{Wei, big.NewInt(1), "1"},
	{Wei, big.NewInt(10), "10"},
	{Wei, big.NewInt(100), "100"},
	{Wei, big.NewInt(1000), "1,000"},
	{Wei, big.NewInt(1000000), "1,000,000"},

	{Kwei, big.NewInt(1), "0.001"},
	{Kwei, big.NewInt(1000), "1"},
	{Kwei, big.NewInt(1000000), "1,000"},

	{Gwei, big.NewInt(1), "0.000000001"},
	{Gwei, big.NewInt(100000000000), "100"},
	{Gwei, big.NewInt(1000000000000), "1,000"},
	{Gwei, big.NewInt(100001000000), "100.001"},

	{Ether, big.NewInt(1), "0.000000000000000001"},
	{Ether, big.NewInt(100000000000), "0.0000001"},
	{Ether, big.NewInt(1000000000000000000), "1"},
	{Ether, newBigInt("1000000000000000000000"), "1,000"},
	{Ether, newBigInt("1000000001000000000000000"), "1,000,000.001"},
}

func TestString(t *testing.T) {
	for _, test := range stringTests {
		result := String(test.value, test.unit)
		if result != test.expected {
			t.Errorf("Received %s but expected %s", result, test.expected)
		}
	}
}
