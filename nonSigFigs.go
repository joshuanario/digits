package digits

import (
	"math"
	"math/big"
	"strconv"
)

func computeNonSigFigs(p Precision, value *big.Float, g rune, d Decimals) (string, error) {
	copy := big.NewFloat(0)
	copy = copy.Add(copy, value)
	if p >= Oneth || p == Exact {
		return text(value, d), nil
	}
	expander, err := expander(p)
	if err != nil {
		return "", err
	}
	shrunk, err := shrink(p, copy)
	if err != nil {
		return "", err
	}
	f := big.NewFloat(0).Mul(shrunk, expander)
	return text(f, d), nil
}
func expander(p Precision) (*big.Float, error) {
	f := strconv.FormatFloat(math.Pow(10, float64(-1*int(p))), 'f', 0, 64)
	shifter, _, err := big.ParseFloat(f, 10, QUAD_PREC, big.ToZero)
	if err != nil {
		return nil, err
	}
	return shifter, nil
}
