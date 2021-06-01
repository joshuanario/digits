package digits

import (
	"math/big"
	"strings"
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

func computeHead(value *big.Float) string {
	sign := value.Sign()
	if sign < 0 {
		return "("
	}
	return ""
}

func computeCore(p Precision, value *big.Float, g rune, d Decimals) (string, error) {
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
	ret := unsignedtext(shrunk, Decimals(prec))
	trunc := p
	if trunc < Oneth {
		trunc = Oneth
	}
	return DigitGroup(trunc, ret, g, d), nil
}

func computeTail(p Precision, v string, g rune, d Decimals) (string, error) {
	copy, err := lowPrecisionTruncate(p, v, d)
	if err != nil {
		return "", err
	}
	signedTail := signedTail(copy)
	stripper, err := stripper(p, copy)
	if err != nil {
		return "", err
	}
	if p >= Oneth {
		if int(p) < int(d) {
			if stripper.Cmp(big.NewFloat(0.0).SetPrec(PREC_BITS)) == 0 {
				zero := zeroAppend("", int(d)-int(p))
				return zero + signedTail, nil
			}
			stripperTail, err := stripperTail(p, stripper, d)
			if err != nil {
				return "", err
			}
			return DigitGroup(Oneth, stripperTail, g, d) + signedTail, nil
		}
		return "" + signedTail, nil
	}
	stripped := copy.Sub(copy, stripper)
	if stripped.Cmp(big.NewFloat(0.0).SetPrec(PREC_BITS)) == 0 {
		i := strings.IndexRune(v, '.')
		zero := zeroize(p, i, d)
		dd := d
		if i < 0 {
			dd = NoDecimals
		}
		return DigitGroup(Oneth, zero, g, dd) + signedTail, nil
	}
	ret := unsignedtext(stripped, Decimals(d))
	return DigitGroup(Oneth, ret, g, d) + signedTail, nil
}
func zeroize(p Precision, i int, d Decimals) string {
	b := ""
	a := ""
	bLen := -1 * int(p)
	if i >= 0 {
		aLen := int(d)
		if int(p) > int(d) {
			aLen = int(p)
		}
		bLen = i
		a = "."
		for i := 0; i < aLen; i++ {
			a += "0"
		}
	}
	for i := 0; i < bLen; i++ {
		b += "0"
	}
	return b + a
}
