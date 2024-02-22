package ethunit

import (
	"math/big"
	"testing"
)

type stringTest struct {
	value     *big.Float
	precision Unit
	expected  string
}

var stringTests = []stringTest{
	{big.NewFloat(1.0), Wei, "1"},
	{big.NewFloat(1.0), Kwei, "1.00"},
	{big.NewFloat(1.0), Mwei, "1.00000"},
	{big.NewFloat(1.0), Gwei, "1.00000000"},
	{big.NewFloat(1.0), Szabo, "1.00000000000"},

	{big.NewFloat(100.0), Wei, "100"},
	{big.NewFloat(100.0), Kwei, "100.00"},
	{big.NewFloat(100.0), Mwei, "100.00000"},
	{big.NewFloat(100.0), Gwei, "100.00000000"},
	{big.NewFloat(100.0), Szabo, "100.00000000000"},

	{big.NewFloat(1000.0), Wei, "1,000"},
	{big.NewFloat(1000.0), Kwei, "1,000.00"},
	{big.NewFloat(1000.0), Mwei, "1,000.00000"},
	{big.NewFloat(1000.0), Gwei, "1,000.00000000"},
	{big.NewFloat(1000.0), Szabo, "1,000.00000000000"},

	{big.NewFloat(1000000.0), Wei, "1,000,000"},
	{big.NewFloat(1000000.0), Kwei, "1,000,000.00"},
	{big.NewFloat(1000000.0), Mwei, "1,000,000.00000"},
	{big.NewFloat(1000000.0), Gwei, "1,000,000.00000000"},
	{big.NewFloat(1000000.0), Szabo, "1,000,000.00000000000"},
}

func TestString(t *testing.T) {
	for _, test := range stringTests {
		result := String(test.value, test.precision)
		if result != test.expected {
			t.Errorf("Received %s but expected %s", result, test.expected)
		}
	}
}
