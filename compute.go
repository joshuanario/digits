package digits

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

const (
	FLOAT       = 24
	DOUBLE_PREC = 53
	QUAD_PREC   = 113
	OCTO_PREC   = 237
	PREC_BITS   = QUAD_PREC
)

func significandBits(k int) uint {
	switch k {
	case 32:
		return FLOAT
	case 64:
		return DOUBLE_PREC
	case 128:
		return QUAD_PREC
	case 256:
		return OCTO_PREC
	}
	ret := k - int(math.Round(4*math.Log2(float64(k)))) + 13
	return uint(ret)
}

func text(v *big.Float, d Decimals) string {
	copy := big.NewFloat(0)
	copy = copy.Add(copy, v)
	dec := int(d)
	f := copy.Text('f', dec)
	if f == "0" {
		f = f + "."
		for i := 0; i < dec; i++ {
			f = f + "0"
		}
	}
	return f
}
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
