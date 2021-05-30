package digits

import "math/big"

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
			stripperTail, err := stripperTail(p, stripper, d)
			if err != nil {
				return "", err
			}
			return DigitGroup(-1, stripperTail, g, d) + signedTail, nil
		}
		return "" + signedTail, nil
	}
	stripped := copy.Sub(copy, stripper)
	ret := unsignedtext(stripped, Decimals(d))
	return DigitGroup(-1, ret, g, d) + signedTail, nil
}
func signedTail(value *big.Float) string {
	sign := value.Sign()
	if sign < 0 {
		return ")"
	}
	return ""
}
