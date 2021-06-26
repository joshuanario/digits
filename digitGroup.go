package digits

import (
	"math/big"
	"strings"
)

func DigitGroup(p Precision, v string, g rune, d Decimals, isSigFig bool) string {
	gChar := string([]rune{g})
	sChar := "." //todo multilingual separator character
	i := strings.IndexRune(v, '.')
	trunc := int(p)
	if trunc > int(d) {
		trunc = int(d)
	}
	if i > -1 {
		truncated := triplefy(v[:i], gChar, isSigFig)
		if d == NoDecimals {
			return truncated
		}
		b := zeroAppend(v[i+1:], trunc-(len(v)-i+-1))
		return truncated + sChar + b
	}
	triplefied := triplefy(v, gChar, isSigFig)
	if d == NoDecimals {
		return triplefied
	}
	return triplefied + zeroAppend(sChar, trunc)
}
func triplefy(v string, g string, isSigFig bool) string {
	i := strings.IndexRune(v, '.')
	if i >= 0 {
		return ""
	}
	z := zeroTriplefy(v, g)
	if z != "" {
		return z
	}
	if len(v) < 3 {
		return v
	}
	modDiff := len(v)%3 + 3
	end := len(v) - modDiff
	lastG := g
	if end == 0 && isSigFig {
		lastG = ""
	}
	return triplefy(v[:end], g, isSigFig) + lastG + v[end:]
}
func zeroTriplefy(v string, g string) string {
	i := strings.IndexRune(v, '.')
	if i >= 0 {
		return ""
	}
	value, _, err := big.ParseFloat(v, 10, PREC_BITS, big.ToZero)
	if err != nil {
		return ""
	}
	if value.Cmp(big.NewFloat(0.0).SetPrec(PREC_BITS)) != 0 {
		return ""
	}
	if len(v) < 3 {
		return v
	}
	modDiff := len(v)%3 + 3
	end := len(v) - modDiff
	return zeroTriplefy(v[:end], g) + g + v[end:]
}
func zeroAppend(v string, i int) string {
	if i <= 0 {
		return v
	}
	ret := v
	for a := 0; a < i; a++ {
		ret += "0"
	}
	return ret
}
