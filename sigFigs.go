package digits

import (
	"math/big"
)

func computeSigFigs(p Precision, value *big.Float, g rune, d Decimals) (string, error) {
	copy := big.NewFloat(0)
	copy = copy.Add(copy, value)
	prec := 0
	if p >= Oneth {
		if p > Oneth || d != NoDecimals {
			prec = int(p)
			dec := int(d)
			if dec > prec {
				prec = dec
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
