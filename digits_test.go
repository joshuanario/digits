package digits_test

import (
	"testing"

	"github.com/joshuanario/digits"
)

type stimulus struct {
	p digits.Precision
	v string
	g rune
	d digits.Decimals
}

var stimuli = []*stimulus{
	{
		p: digits.Oneth,
		v: "0.99",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
	{
		p: digits.Hundredth,
		v: "0.02",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
	{
		p: digits.Hundredth,
		v: "-0.02",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
	{
		p: digits.Millions,
		v: "77190000",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
	{
		p: digits.Millions,
		v: "77190000.00009",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
	{
		p: digits.Millions,
		v: "-77190000.00009",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
	{
		p: digits.Hundredth,
		v: "-77190000.00009",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
}

func act() []*digits.Expression {
	var suts = []*digits.Expression{}
	for _, s := range stimuli {
		d, err := digits.New(s.p, s.v, s.g, s.d)
		if err != nil {
			panic(err)
		}
		suts = append(suts, d)
	}
	return suts
}
func Test_SigFigs(t *testing.T) {
	var suts = act()
	expectations := []string{
		"0.",
		"0.02",
		"0.02",
		"77",
		"77",
		"77",
		"77190000.00",
	}
	for i, expectation := range expectations {
		sut := suts[i]
		output := sut.SigFigs()
		if output != expectation {
			t.Fail()
		}
	}
}
func Test_NonSigFigs(t *testing.T) {
	var suts = act()
	expectations := []string{
		"99",
		"",
		"",
		"190000.00",
		"190000.00",
		"190000.00",
		"",
	}
	for i, expectation := range expectations {
		sut := suts[i]
		output := sut.NonSigFigs()
		if output != expectation {
			t.Fail()
		}
	}
}
func Test_String(t *testing.T) {
	var suts = act()
	expectations := []string{
		"0.00",
		"0.02",
		"(0.02)",
		"77,190,000.00",
		"77,190,000.00",
		"(77,190,000.00)",
		"(77,190,000.00)",
	}
	for i, expectation := range expectations {
		sut := suts[i]
		if sut.String() != expectation {
			t.Fail()
		}
	}
}
func Test_Head(t *testing.T) {
	var suts = act()
	expectations := []string{
		"",
		"",
		"(",
		"",
		"",
		"(",
		"(",
	}
	for i, expectation := range expectations {
		sut := suts[i]
		output := sut.Head()
		if output != expectation {
			t.Fail()
		}
	}
}
func Test_Core(t *testing.T) {
	var suts = act()
	expectations := []string{
		"0.",
		"0.02",
		"0.02",
		"77",
		"77",
		"77",
		"77,190,000.00",
	}
	for i, expectation := range expectations {
		sut := suts[i]
		output := sut.Core()
		if output != expectation {
			t.Fail()
		}
	}
}
func Test_Tail(t *testing.T) {
	var suts = act()
	expectations := []string{
		"00",
		"",
		")",
		",190,000.00",
		",190,000.00",
		",190,000.00)",
		")",
	}
	for i, expectation := range expectations {
		sut := suts[i]
		output := sut.Tail()
		if output != expectation {
			t.Fail()
		}
	}
}
