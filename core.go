package digits

import (
	"math/big"
)

func computeCore(p Precision, value *big.Float, g rune) (string, error) {
	copy := big.NewFloat(0)
	copy = copy.Add(copy, value)
	shrunk, err := shrink(p, copy)
	if err != nil {
		return "", err
	}
	prec := 0
	if p > Oneth {
		prec = int(p)
	}
	return unsignedtext(shrunk, Decimals(prec)), nil
}
