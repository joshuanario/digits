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
		p: digits.Thousands,
		v: "396006000",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
	{
		p: digits.Thousands,
		v: "487000",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
	{
		p: digits.Thousands,
		v: "4059000",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
	{
		p: digits.Thousands,
		v: "45121000",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
	{
		p: digits.Thousands,
		v: "98000",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
	{
		p: digits.Thousands,
		v: "-338863000",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
	{
		p: digits.Exact,
		v: "1038807",
		g: ',',
		d: digits.PreserveUpToHundredth,
	},
	{
		p: digits.HundredThousands,
		v: "80800000",
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
		"396006",
		"487",
		"4059",
		"45121",
		"98",
		"338863",
		"1038807",
		"808",
	}
	for i, expectation := range expectations {
		sut := suts[i]
		output := sut.SigFigs()
		if output != expectation {
			t.Fatalf("SigFigs Failed: expected %s, got %s. Test Index %d", expectation, output, i)
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
		"000",
		"000",
		"000",
		"000",
		"000",
		"000",
		"",
		"00000",
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
		"396,006,000.00",
		"487,000.00",
		"4,059,000.00",
		"45,121,000.00",
		"98,000.00",
		"(338,863,000.00)",
		"1,038,807.00",
		"80,800,000.00",
	}
	for i, expectation := range expectations {
		sut := suts[i]
		output := sut.String()
		if output != expectation {
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
		"",
		"",
		"",
		"",
		"",
		"(",
		"",
		"",
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
		"396,006",
		"487",
		"4,059",
		"45,121",
		"98",
		"338,863",
		"1,038,807.",
		"80,8",
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
		",000.00",
		",000.00",
		",000.00",
		",000.00",
		",000.00",
		",000.00)",
		"00",
		"00,000.00",
	}
	for i, expectation := range expectations {
		sut := suts[i]
		output := sut.Tail()
		if output != expectation {
			t.Fail()
		}
	}
}
