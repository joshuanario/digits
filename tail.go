package digits

func computeTail(p Precision, v string, g rune, d Decimals) (string, error) {
	copy, err := lowPrecisionTruncate(p, v, d)
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
