package digits

import "fmt"

type Expression struct {
	sign       bool
	nonSigFigs string
	sigFigs    string
	stringOut  string
	head       string
	core       string
	tail       string
}

func New(significantPrecision Precision, value string, groupSeparator rune, decimalPrecision Decimals) (*Expression, error) {
	if significantPrecision != Exact && !(significantPrecision > Trillions && significantPrecision < Trillionth) {
		return nil, fmt.Errorf("precision out of bounds")
	}
	ret := Expression{}
	trunc, err := lowPrecisionTruncate(significantPrecision, value, decimalPrecision)
	if err != nil {
		return nil, err
	}
	ret.sign = trunc.Signbit()
	sigFigs, err := computeSigFigs(significantPrecision, value, groupSeparator, decimalPrecision)
	if err != nil {
		return nil, err
	}
	ret.sigFigs = sigFigs
	nonSigFigs, err := computeNonSigFigs(significantPrecision, value, decimalPrecision)
	if err != nil {
		return nil, err
	}
	ret.nonSigFigs = nonSigFigs
	stringOut, err := computeString(sigFigs, nonSigFigs, groupSeparator, decimalPrecision, ret.sign)
	if err != nil {
		return nil, err
	}
	ret.stringOut = stringOut
	ret.head = computeHead(trunc)
	//core, err := computeCore(significantPrecision, value, groupSeparator, decimalPrecision)
	core, tail, err := computeCoreTail(sigFigs, stringOut, groupSeparator)
	if err != nil {
		return nil, err
	}
	ret.core = core
	ret.tail = tail
	return &ret, nil
}
func (d *Expression) SigFigs() string {
	return d.sigFigs
}
func (d *Expression) NonSigFigs() string {
	return d.nonSigFigs
}
func (d *Expression) String() string {
	return d.stringOut
}
func (d *Expression) Head() string {
	return d.head
}
func (d *Expression) Core() string {
	return d.core
}
func (d *Expression) Tail() string {
	return d.tail
}
