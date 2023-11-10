package digits_test

import (
	"testing"

	"github.com/joshuanario/digits"
)

type Stimulus struct {
	precision           digits.Precision
	value               string
	groupSeparator      rune
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
	stimulus Stimulus
	output   Output
}

var suts = []*SUT{
	{
		stimulus: Stimulus{
			precision:           digits.Oneth,
			value:               "0.99",
			groupSeparator:      ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "0.",
			nonSigFigs: "99",
			strOut:     "0.99",
			head:       "",
			core:       "0.",
			tail:       "99",
		},
	},
	{
		stimulus: Stimulus{
			precision:           digits.Hundredth,
			value:               "0.02",
			groupSeparator:      ',',
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
		stimulus: Stimulus{
			precision:           digits.Hundredth,
			value:               "-0.02",
			groupSeparator:      ',',
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
		stimulus: Stimulus{
			precision:           digits.Millions,
			value:               "77190000",
			groupSeparator:      ',',
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
		stimulus: Stimulus{
			precision:           digits.Millions,
			value:               "77190000.00009",
			groupSeparator:      ',',
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
		stimulus: Stimulus{
			precision:           digits.Millions,
			value:               "-77190000.00009",
			groupSeparator:      ',',
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
		stimulus: Stimulus{
			precision:           digits.Thousands,
			value:               "396006000",
			groupSeparator:      ',',
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
		stimulus: Stimulus{
			precision:           digits.Thousands,
			value:               "487000",
			groupSeparator:      ',',
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
		stimulus: Stimulus{
			precision:           digits.Thousands,
			value:               "4059000",
			groupSeparator:      ',',
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
		stimulus: Stimulus{
			precision:           digits.Thousands,
			value:               "45121000",
			groupSeparator:      ',',
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
		stimulus: Stimulus{
			precision:           digits.Thousands,
			value:               "98000",
			groupSeparator:      ',',
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
		stimulus: Stimulus{
			precision:           digits.Thousands,
			value:               "-338863000",
			groupSeparator:      ',',
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
		stimulus: Stimulus{
			precision:           digits.Exact,
			value:               "1038807",
			groupSeparator:      ',',
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
		stimulus: Stimulus{
			precision:           digits.HundredThousands,
			value:               "80800000",
			groupSeparator:      ',',
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
	{
		stimulus: Stimulus{
			precision:           digits.HundredThousands,
			value:               "-80800000",
			groupSeparator:      ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "808",
			nonSigFigs: "00000",
			strOut:     "(80,800,000.00)",
			head:       "(",
			core:       "80,8",
			tail:       "00,000.00)",
		},
	},
	{
		stimulus: Stimulus{
			precision:           digits.Tenth,
			value:               "80800000.99090909090",
			groupSeparator:      ',',
			fractionalPrecision: digits.PreserveUpToHundredth,
		},
		output: Output{
			sigFigs:    "80800000.9",
			nonSigFigs: "9",
			strOut:     "80,800,000.99",
			head:       "",
			core:       "80,800,000.9",
			tail:       "9",
		},
	},
}

func process(stimulus Stimulus) *digits.Expression {
	d, err := digits.New(stimulus.precision, stimulus.value, stimulus.groupSeparator, stimulus.fractionalPrecision)
	if err != nil {
		panic(err)
	}
	return d
}
func Test_SigFigs(t *testing.T) {
	for _, sut := range suts {
		var expr = process(sut.stimulus)
		uut := expr.SigFigs()
		if uut != sut.output.sigFigs {
			t.Fail()
		}
	}
}
func Test_NonSigFigs(t *testing.T) {
	for _, sut := range suts {
		var expr = process(sut.stimulus)
		uut := expr.NonSigFigs()
		if uut != sut.output.nonSigFigs {
			t.Fail()
		}
	}
}
func Test_String(t *testing.T) {
	for _, sut := range suts {
		var expr = process(sut.stimulus)
		uut := expr.String()
		if uut != sut.output.strOut {
			t.Fail()
		}
	}
}
func Test_Head(t *testing.T) {
	for _, sut := range suts {
		var expr = process(sut.stimulus)
		uut := expr.Head()
		if uut != sut.output.head {
			t.Fail()
		}
	}
}
func Test_Core(t *testing.T) {
	for _, sut := range suts {
		var expr = process(sut.stimulus)
		uut := expr.Core()
		if uut != sut.output.core {
			t.Fail()
		}
	}
}
func Test_Tail(t *testing.T) {
	for _, sut := range suts {
		var expr = process(sut.stimulus)
		uut := expr.Tail()
		if uut != sut.output.tail {
			t.Fail()
		}
	}
}
