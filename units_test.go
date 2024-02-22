package ethunit

import (
	"math/big"
	"testing"
)

type intTest struct {
	value            *big.Int
	fromUnit, toUnit Unit
	expected         *big.Float
}

var intTests = []intTest{
	{big.NewInt(1), Wei, Wei, big.NewFloat(1)},
	{big.NewInt(1), Wei, Kwei, big.NewFloat(0.001)},
	{big.NewInt(1), Wei, Mwei, big.NewFloat(0.000001)},
	{big.NewInt(1), Wei, Gwei, big.NewFloat(0.000000001)},
	{big.NewInt(1), Wei, Szabo, big.NewFloat(0.000000000001)},
	{big.NewInt(1), Wei, Finney, big.NewFloat(0.000000000000001)},
	{big.NewInt(1), Wei, Ether, big.NewFloat(0.000000000000000001)},

	{big.NewInt(1), Ether, Wei, big.NewFloat(1000000000000000000)},
	{big.NewInt(1), Ether, Kwei, big.NewFloat(1000000000000000)},
	{big.NewInt(1), Ether, Mwei, big.NewFloat(1000000000000)},
	{big.NewInt(1), Ether, Gwei, big.NewFloat(1000000000)},
	{big.NewInt(1), Ether, Szabo, big.NewFloat(1000000)},
	{big.NewInt(1), Ether, Finney, big.NewFloat(1000)},
	{big.NewInt(1), Ether, Ether, big.NewFloat(1)},
}

func TestConvertInt(t *testing.T) {
	for _, test := range intTests {
		result := ConvertInt(test.value, test.fromUnit, test.toUnit)
		if result.Cmp(test.expected) != 0 {
			t.Errorf("Received %s but expected %s: %d", result.Text('f', 18), test.expected.Text('f', 18), result.Cmp(test.expected))
		}
	}
}

type floatTest struct {
	value            *big.Float
	fromUnit, toUnit Unit
	expected         *big.Float
}

var floatTests = []floatTest{
	{big.NewFloat(1), Wei, Wei, big.NewFloat(1)},
	{big.NewFloat(1), Wei, Kwei, big.NewFloat(0.001)},
	{big.NewFloat(1), Wei, Mwei, big.NewFloat(0.000001)},
	{big.NewFloat(1), Wei, Gwei, big.NewFloat(0.000000001)},
	{big.NewFloat(1), Wei, Szabo, big.NewFloat(0.000000000001)},
	{big.NewFloat(1), Wei, Finney, big.NewFloat(0.000000000000001)},
	{big.NewFloat(1), Wei, Ether, big.NewFloat(0.000000000000000001)},

	{big.NewFloat(1), Ether, Wei, big.NewFloat(1000000000000000000)},
	{big.NewFloat(1), Ether, Kwei, big.NewFloat(1000000000000000)},
	{big.NewFloat(1), Ether, Mwei, big.NewFloat(1000000000000)},
	{big.NewFloat(1), Ether, Gwei, big.NewFloat(1000000000)},
	{big.NewFloat(1), Ether, Szabo, big.NewFloat(1000000)},
	{big.NewFloat(1), Ether, Finney, big.NewFloat(1000)},
	{big.NewFloat(1), Ether, Ether, big.NewFloat(1)},
}

func TestConvertFloat(t *testing.T) {
	for _, test := range floatTests {
		result := ConvertFloat(test.value, test.fromUnit, test.toUnit)
		if result.Cmp(test.expected) != 0 {
			t.Errorf("Received %s but expected %s: %d", result.Text('f', 18), test.expected.Text('f', 18), result.Cmp(test.expected))
		}
	}
}
