package digits

type Expression struct {
	sign       bool
	nonSigFigs string
	sigFigs    string
	head       string
	core       string
	tail       string
}

func New(p Precision, v string, g rune, d Decimals) (*Expression, error) {
	ret := Expression{}
	value, err := lowPrecisionTruncate(p, v, d)
	if err != nil {
		return nil, err
	}
	ret.sign = value.Signbit()
	sigFigs, err := computeSigFigs(p, value, g)
	if err != nil {
		return nil, err
	}
	ret.sigFigs = sigFigs
	nonSigFigs, err := computeNonSigFigs(p, v, g, d)
	if err != nil {
		return nil, err
	}
	ret.nonSigFigs = nonSigFigs
	ret.head = computeHead(value)
	core, err := computeCore(p, value, g)
	if err != nil {
		return nil, err
	}
	ret.core = core

	return &ret, nil
}
func (d *Expression) SigFigs() string {
	return d.sigFigs
}
func (d *Expression) NonSigFigs() string {
	return d.nonSigFigs
}
func (d *Expression) String() string {
	return d.head + d.core + d.tail
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
