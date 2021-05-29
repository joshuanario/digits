package digits

import (
	"math/big"
)

func computeSigFigs(p Precision, value *big.Float, g rune) (string, error) {
	copy := big.NewFloat(0)
	copy = copy.Add(copy, value)
	prec := 0
	if p >= Oneth {
		if p > Oneth {
			prec = int(p)
			if prec > int(MaximumDecimals) {
				prec = int(MaximumDecimals)
			}
		}
		return unsignedtext(copy, Decimals(prec)), nil
	}
	shrunk, err := shrink(p, copy)
	if err != nil {
		return "", err
	}
	return unsignedtext(shrunk, NoDecimals), nil
}
