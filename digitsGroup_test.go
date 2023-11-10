package digits

type Stimulus struct {
	value               string
	groupingCharRune    rune
	groupingCharStr     string
	ifSigFig            bool
	fractionalPrecision Decimals
}
type Output struct {
	triplified string
	final      string
}
type TestCase struct {
	stimulus Stimulus
	output   Output
}

var testCases = []*TestCase{
	{
		stimulus: Stimulus{
			value:               "45121000000000.00",
			groupingCharRune:    ',',
			groupingCharStr:     ",",
			ifSigFig:            true,
			fractionalPrecision: PreserveUpToHundredth,
		},
		output: Output{
			triplified: "45,121,000,000,000",
			final:      "45,121,000,000,000.00",
		},
	},
	{
		stimulus: Stimulus{
			value:               "80800000.00",
			groupingCharRune:    ',',
			groupingCharStr:     ",",
			ifSigFig:            true,
			fractionalPrecision: PreserveUpToHundredth,
		},
		output: Output{
			triplified: "80,800,000",
			final:      "80,800,000.00",
		},
	},
}

/*
func Test_DigitGroup(t *testing.T) {
	for _, testCase := range testCases {
		sut := DigitGroup(Exact, testCase.stimulus.value, testCase.stimulus.groupingCharRune, testCase.stimulus.fractionalPrecision, testCase.stimulus.ifSigFig)
		if sut != testCase.output.final {
			t.Fail()
		}
	}
}

func Test_Triplefy(t *testing.T) {
	for _, testCase := range testCases {
		sut := triplefy(testCase.stimulus.value, testCase.stimulus.groupingCharStr, testCase.stimulus.ifSigFig)
		if sut != testCase.output.triplified {
			t.Fail()
		}
	}
}

func Test_ZeroTriplefy(t *testing.T) {
	//TODO: flesh this out in stimulus
	for _, testCase := range testCases {
		sut := zeroTriplefy(testCase.stimulus.value, testCase.stimulus.groupingCharStr)
		if sut != testCase.output.triplified {
			t.Fail()
		}
	}
}
*/
