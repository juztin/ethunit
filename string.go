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

func String(f *big.Float, u Unit) string {
	precision := int(u)
	if precision > 0 {
		precision--
	}
	s := fmt.Sprintf(f.Text('f', precision))
	i := strings.Index(s, ".")
	if i == -1 {
		return commafy(s)
	}
	return fmt.Sprintf("%s.%s", commafy(s[:i]), s[i+1:])
}
