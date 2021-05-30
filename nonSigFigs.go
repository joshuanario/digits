package digits

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func computeNonSigFigs(p Precision, v string, d Decimals) (string, error) {
	copy, err := highPrecisionTruncate(p, v, d)
	if err != nil {
		return "", err
	}
	stripper, err := stripper(p, copy)
	if err != nil {
		return "", err
	}
	if p >= Oneth {
		if int(p) < int(d) {
			return stripperTail(p, stripper, d)
		}
		return "", nil
	}
	stripped := copy.Sub(copy, stripper)
	ret := unsignedtext(stripped, Decimals(d))
	return ret, nil
}
func stripperTail(p Precision, value *big.Float, d Decimals) (string, error) {
	if p < Oneth {
		return "", nil
	}
	copy := big.NewFloat(0)
	copy = copy.Add(copy, value)
	f := copy.Text('f', int(MaximumDecimals))
	i := strings.IndexRune(f, '.')
	if i < 0 {
		return "", fmt.Errorf("invalid precision value")
	}
	dd := int(d)
	pp := int(p)
	return f[i+pp+1 : i+dd+1], nil
}
func stripper(p Precision, value *big.Float) (*big.Float, error) {
	copy := big.NewFloat(0)
	copy = copy.Add(copy, value)
	shrunk, err := shrink(p, copy)
	if err != nil {
		return nil, err
	}
	return expand(p, shrunk)
}
func expander(p Precision) (*big.Float, error) {
	exp := float64(0)
	if p < Oneth && p != Exact {
		exp = float64(-1 * int(p))
	}
	f := strconv.FormatFloat(math.Pow(10, exp), 'f', 0, 64)
	shifter, _, err := big.ParseFloat(f, 10, PREC_BITS, big.ToZero)
	if err != nil {
		return nil, err
	}
	return shifter, nil
}
func expand(p Precision, value *big.Float) (*big.Float, error) {
	copy := big.NewFloat(0)
	copy = copy.Add(copy, value)
	expander, err := expander(p)
	if err != nil {
		return nil, err
	}
	f := copy.Mul(copy, expander).Text('f', int(MaximumDecimals))
	expanded, _, err := big.ParseFloat(f, 10, PREC_BITS, big.ToZero)
	if err != nil {
		return nil, err
	}
	return expanded, nil
}
