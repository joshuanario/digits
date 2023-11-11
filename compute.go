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

/* Generates a formatted financial value string from the already calculated significant figures.
 */
func computeString(precision Precision, sigFigs, nonSigFigs string, groupSeparator rune, decimalPrecision Decimals, sign bool) (string, error) {
	// TODO: implement error handlers
	decimalSeparator := determineDecimalRune(groupSeparator)
	var result strings.Builder
	var builder strings.Builder
	if sign {
		result.WriteRune('(')
	}
	builder.WriteString(sigFigs)
	if (strings.IndexRune(sigFigs, decimalSeparator) == -1 && strings.IndexRune(nonSigFigs, decimalSeparator) == -1) &&
		(precision == Exact || precision >= Oneth) {
		builder.WriteRune(decimalSeparator)
		builder.WriteString(nonSigFigs)
	} else {
		builder.WriteString(nonSigFigs)
		builder.WriteRune(decimalSeparator)
	}
	postDecimal := len(builder.String()) - strings.IndexRune(builder.String(), decimalSeparator) - 1
	if postDecimal < int(decimalPrecision) {
		builder.WriteString(zeroAppend("", int(decimalPrecision)-postDecimal))
	}
	parts := strings.Split(builder.String(), string(decimalSeparator))
	integerPart, fractionalPart := parts[0], parts[1]
	for i, char := range integerPart {
		if i > 0 && (len(integerPart)-i)%3 == 0 {
			result.WriteRune(groupSeparator)
		}
		result.WriteRune(char)
	}
	result.WriteRune(decimalSeparator)
	result.WriteString(fractionalPart)
	if sign {
		result.WriteRune(')')
	}
	return result.String(), nil
}

/* Generates the Core and Tail from the generated string from computeString().
 */
func computeComponents(precision Precision, sigFig, stringOut string, groupSeparator rune) (string, string, string, error) {
	// TODO: implement error handlers
	var core strings.Builder
	head := ""
	decimalSeparator := determineDecimalRune(groupSeparator)
	if rune(stringOut[0]) == '(' {
		head = "("
		stringOut = stringOut[1:]
	}
	iAdjusted := 0
	for i := 0; i < len(sigFig); i++ {
		if rune(stringOut[iAdjusted]) == groupSeparator {
			core.WriteRune(rune(stringOut[iAdjusted]))
			iAdjusted++
		}
		if rune(stringOut[iAdjusted]) == rune(sigFig[i]) {
			core.WriteRune(rune(sigFig[i]))
		}
		iAdjusted++
	}
	decimalIndex := strings.IndexRune(core.String(), decimalSeparator)
	if decimalIndex == -1 && precision == Exact {
		core.WriteRune(decimalSeparator)
		iAdjusted++
	}
	return head, core.String(), stringOut[iAdjusted:], nil
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
