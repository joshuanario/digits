package digits

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func unsignedtext(v *big.Float, d Decimals) string {
	copy := big.NewFloat(0)
	copy = copy.Add(copy, v)
	dec := int(d)
	f := copy.Abs(copy).Text('f', dec)
	if f == "0" {
		f = f + "."
		for i := 0; i < dec; i++ {
			f = f + "0"
		}
	}
	return f
}
func low(p Precision, d Decimals) int {
	prec := 0
	if p >= Oneth {
		if p > Oneth || d != NoDecimals {
			prec = int(p)
			dec := int(d)
			if dec < prec {
				prec = dec
			}
		}
	}
	return prec
}
func high(p Precision, d Decimals) int {
	prec := 0
	if p >= Oneth {
		if p > Oneth || d != NoDecimals {
			prec = int(p)
			dec := int(d)
			if dec > prec {
				prec = dec
			}
		}
	}
	return prec
}
func lowPrecisionTruncate(p Precision, v string, d Decimals) (*big.Float, error) {
	prec := low(p, d)
	i := strings.IndexRune(v, '.')
	f := v
	if i > -1 {
		f = v[:i+prec+1]
	}
	truncated, _, err := big.ParseFloat(f, 10, PREC_BITS, big.ToZero)
	return truncated, err
}
func highPrecisionTruncate(p Precision, v string, d Decimals) (*big.Float, error) {
	prec := high(p, d)
	i := strings.IndexRune(v, '.')
	f := v
	// TODO: this is where the error occurs, with index out of bounds.
	// Consider calculating the number of appended zeros needed, then parse float, or see if parseFloat can be passed a value to set the precision.
	if i > -1 {
		f = v[:i+prec+1]
	}
	truncated, _, err := big.ParseFloat(f, 10, PREC_BITS, big.ToZero)
	return truncated, err
}
func shrinker(p Precision) (*big.Float, error) {
	if p >= Oneth || p == Exact {
		one, _, err := big.ParseFloat("1", 10, PREC_BITS, big.ToZero)
		return one, err
	}
	num := math.Pow(10, float64(int(p)))
	prec := 0
	if p < Oneth {
		prec = -1 * int(p)
	}
	f := strconv.FormatFloat(num, 'f', prec, 64)
	shifter, _, err := big.ParseFloat(f, 10, PREC_BITS, big.ToZero)
	if err != nil {
		return nil, err
	}
	return shifter, nil
}
func shrink(p Precision, value *big.Float) (*big.Float, error) {
	copy := big.NewFloat(0)
	copy = copy.Add(copy, value)
	if p >= Oneth || p == Exact {
		return copy, nil
	}
	shrinker, err := shrinker(p)
	if err != nil {
		return nil, err
	}
	f := copy.Mul(copy, shrinker).Text('f', 0)
	shrunk, _, err := big.ParseFloat(f, 10, PREC_BITS, big.ToZero)
	if err != nil {
		return nil, err
	}
	return shrunk, nil
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

/* Determines the rune used for the decimal, for localization.
 */
func determineDecimalRune(groupSeparator rune) rune {
	if groupSeparator == '.' {
		return ','
	}
	return '.'
}

/* Quick function for appending zeros.
 */
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
