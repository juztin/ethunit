package ethunit

import (
	"fmt"
	"math/big"
	"strings"
)

func commafy(s string) string {
	l := len(s)
	if l < 4 {
		return s
	}
	lb := ((l - 1) / 3) + l
	b := make([]byte, lb, lb)
	for i, j := 0, 0; i < l; i, j = i+1, j+1 {
		if i > 0 && (l-i)%3 == 0 {
			b[j] = ','
			j++
		}
		b[j] = s[i]
	}
	return string(b)
}

func String(d *big.Int, u Unit) string {
	s := d.String()
	if len(s) <= int(u) {
		padded := fmt.Sprintf("0.%s%s", strings.Repeat("0", int(u)-len(s)), s)
		return strings.TrimRight(padded, "0")
	} else {
		decimal := len(s) - int(u)
		integer := s[:decimal]
		if decimal < len(s) {
			mantissa := strings.TrimRight(s[decimal:], "0")
			if len(mantissa) > 0 {
				return fmt.Sprintf("%s.%s", commafy(integer), mantissa)
			} else {
				return fmt.Sprintf("%s", commafy(integer))
			}
		}
		return commafy(integer)
	}
	return s
}
