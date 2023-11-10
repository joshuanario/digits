package digits_test

import (
	"testing"

	"github.com/joshuanario/digits"
)

type Input struct {
	precision           digits.Precision
	value               string
	radixChar           rune
	fractionalPrecision digits.Decimals
}
type Output struct {
	sigFigs    string
	nonSigFigs string
	strOut     string
	head       string
	core       string
	tail       string
}
type SUT struct {
	input  Input
	output Output
}

var sutsd = []*SUT{
	{
		input: Input{
			precision:           digits.Oneth,
			value:               "0.99",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "0.",
			nonSigFigs: "99",
			strOut:     "0.00",
			head:       "",
			core:       "0.",
			tail:       "00",
		},
	},
	{
		input: Input{
			precision:           digits.Hundredth,
			value:               "0.02",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "0.02",
			nonSigFigs: "",
			strOut:     "0.02",
			head:       "",
			core:       "0.02",
			tail:       "",
		},
	},
	{
		input: Input{
			precision:           digits.Hundredth,
			value:               "-0.02",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "0.02",
			nonSigFigs: "",
			strOut:     "(0.02)",
			head:       "(",
			core:       "0.02",
			tail:       ")",
		},
	},
	{
		input: Input{
			precision:           digits.Millions,
			value:               "77190000",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "77",
			nonSigFigs: "190000.00",
			strOut:     "77,190,000.00",
			head:       "",
			core:       "77",
			tail:       ",190,000.00",
		},
	},
	{
		input: Input{
			precision:           digits.Millions,
			value:               "77190000.00009",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "77",
			nonSigFigs: "190000.00",
			strOut:     "77,190,000.00",
			head:       "",
			core:       "77",
			tail:       ",190,000.00",
		},
	},
	{
		input: Input{
			precision:           digits.Millions,
			value:               "-77190000.00009",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "77",
			nonSigFigs: "190000.00",
			strOut:     "(77,190,000.00)",
			head:       "(",
			core:       "77",
			tail:       ",190,000.00)",
		},
	},
	{
		input: Input{
			precision:           digits.Thousands,
			value:               "396006000",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "396006",
			nonSigFigs: "000",
			strOut:     "396,006,000.00",
			head:       "",
			core:       "396,006",
			tail:       ",000.00",
		},
	},
	{
		input: Input{
			precision:           digits.Thousands,
			value:               "487000",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "487",
			nonSigFigs: "000",
			strOut:     "487,000.00",
			head:       "",
			core:       "487",
			tail:       ",000.00",
		},
	},
	{
		input: Input{
			precision:           digits.Thousands,
			value:               "4059000",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "4059",
			nonSigFigs: "000",
			strOut:     "4,059,000.00",
			head:       "",
			core:       "4,059",
			tail:       ",000.00",
		},
	},
	{
		input: Input{
			precision:           digits.Thousands,
			value:               "45121000",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "45121",
			nonSigFigs: "000",
			strOut:     "45,121,000.00",
			head:       "",
			core:       "45,121",
			tail:       ",000.00",
		},
	},
	{
		input: Input{
			precision:           digits.Thousands,
			value:               "98000",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "98",
			nonSigFigs: "000",
			strOut:     "98,000.00",
			head:       "",
			core:       "98",
			tail:       ",000.00",
		},
	},
	{
		input: Input{
			precision:           digits.Thousands,
			value:               "-338863000",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "338863",
			nonSigFigs: "000",
			strOut:     "(338,863,000.00)",
			head:       "(",
			core:       "338,863",
			tail:       ",000.00)",
		},
	},
	{
		input: Input{
			precision:           digits.Exact,
			value:               "1038807",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "1038807",
			nonSigFigs: "",
			strOut:     "1,038,807.00",
			head:       "",
			core:       "1,038,807.",
			tail:       "00",
		},
	},
	{
		input: Input{
			precision:           digits.HundredThousands,
			value:               "80800000",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "808",
			nonSigFigs: "00000",
			strOut:     "80,800,000.00",
			head:       "",
			core:       "80,8",
			tail:       "00,000.00",
		},
	},
}

var suts = []*SUT{
	/*{
		input: Input{
			precision:           digits.Thousands,
			value:               "45121000",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "45121",
			nonSigFigs: "000",
			strOut:     "45,121,000.00",
			head:       "",
			core:       "45,121",
			tail:       ",000.00",
		},
	},*/
	{
		input: Input{
			precision:           digits.Thousands,
			value:               "80800000",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "80800",
			nonSigFigs: "000",
			strOut:     "80,800,000.00",
			head:       "",
			core:       "80,800",
			tail:       ",000.00",
		},
	},
	{
		input: Input{
			precision:           digits.HundredThousands,
			value:               "80800000",
			radixChar:           ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "808",
			nonSigFigs: "00000",
			strOut:     "80,800,000.00",
			head:       "",
			core:       "80,8",
			tail:       "00,000.00",
		},
	},
}

func process(input Input) *digits.Expression {
	d, err := digits.New(input.precision, input.value, input.radixChar, input.fractionalPrecision)
	if err != nil {
		panic(err)
	}
	return d
}
func Test_SigFigs2(t *testing.T) {
	for _, sut := range suts {
		var expr = process(sut.input)
		uut := expr.SigFigs()
		if uut != sut.output.sigFigs {
			t.Fail()
		}
	}
}
func Test_NonSigFigs2(t *testing.T) {
	for _, sut := range suts {
		var expr = process(sut.input)
		uut := expr.NonSigFigs()
		if uut != sut.output.nonSigFigs {
			t.Fail()
		}
	}
}
func Test_String2(t *testing.T) {
	for _, sut := range suts {
		var expr = process(sut.input)
		uut := expr.String()
		if uut != sut.output.strOut {
			t.Fail()
		}
	}
}
func Test_Head2(t *testing.T) {
	for _, sut := range suts {
		var expr = process(sut.input)
		uut := expr.Head()
		if uut != sut.output.head {
			t.Fail()
		}
	}
}
func Test_Core2(t *testing.T) {
	for _, sut := range suts {
		var expr = process(sut.input)
		uut := expr.Core()
		if uut != sut.output.core {
			t.Fail()
		}
	}
}
func Test_Tail2(t *testing.T) {
	for _, sut := range suts {
		var expr = process(sut.input)
		uut := expr.Tail()
		if uut != sut.output.tail {
			t.Fail()
		}
	}
}

func Test_arb(t *testing.T) {
	test := 5 % 3
	if test < 0 {
		t.Fail()
	}
}
