package digits

import (
	"math/big"
	"strings"
)

func computeSigFigs(p Precision, v string, g rune, d Decimals) (string, error) {
	if p == Exact {
		dot := strings.IndexRune(v, '.')
		if dot < 0 {
			return computeSigFigs(Oneth, v, g, d)
		}
		highPrec := len(v) - dot - 1
		return computeSigFigs(Precision(highPrec), v, g, d)
	}
	value, err := lowPrecisionTruncate(p, v, d)
	if err != nil {
		return "", err
	}
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
	if p == Exact {
		return "", nil
	}
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
			if stripper.Cmp(big.NewFloat(0.0).SetPrec(PREC_BITS)) == 0 {
				zero := zeroAppend("", int(d)-int(p))
				return zero, nil
			}
			return stripperTail(p, stripper, d)
		}
		return "", nil
	}
	stripped := copy.Sub(copy, stripper)
	if stripped.Cmp(big.NewFloat(0.0).SetPrec(PREC_BITS)) == 0 {
		i := strings.IndexRune(v, '.')
		zero := zeroize(p, i, d)
		return zero, nil
	}
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

func computeCore(p Precision, v string, g rune, d Decimals) (string, error) {
	if p == Exact {
		dot := strings.IndexRune(v, '.')
		if dot < 0 {
			return computeCore(Oneth, v, g, d)
		}
		highPrec := len(v) - dot - 1
		return computeCore(Precision(highPrec), v, g, d)
	}
	value, err := lowPrecisionTruncate(p, v, d)
	if err != nil {
		return "", err
	}
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
	if p < Oneth {
		return DigitGroup(Oneth, ret, g, NoDecimals, true), nil
	}
	return DigitGroup(p, ret, g, d, true), nil
}

func computeTail(p Precision, v string, g rune, d Decimals) (string, error) {
	if p == Exact {
		dot := strings.IndexRune(v, '.')
		if dot < 0 {
			return zeroAppend("", int(d)), nil
		}
		trunc := len(v) - dot - 1
		if trunc < int(d) {
			return zeroAppend("", int(d)-trunc), nil
		}
		return "", nil
	}
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
			return DigitGroup(Oneth, stripperTail, g, d, false) + signedTail, nil
		}
		return "" + signedTail, nil
	}
	stripped := copy.Sub(copy, stripper)
	if stripped.Cmp(big.NewFloat(0.0).SetPrec(PREC_BITS)) == 0 {
		i := strings.IndexRune(v, '.')
		zero := zeroize(p, i, d)
		return DigitGroup(Precision(d), zero, g, d, false) + signedTail, nil
	}
	ret := unsignedtext(stripped, Decimals(d))
	return DigitGroup(Oneth, ret, g, d, false) + signedTail, nil
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
