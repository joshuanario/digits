package digits

type stimulus struct {
	value               string
	groupingCharRune    rune
	groupingCharStr     string
	ifSigFig            bool
	fractionalPrecision Decimals
}
type output struct {
	triplified string
	final      string
}
type testCase struct {
	stimulus stimulus
	output   output
}

var testCases = []*testCase{
	{
		stimulus: stimulus{
			value:               "45121000000000.00",
			groupingCharRune:    ',',
			groupingCharStr:     ",",
			ifSigFig:            true,
			fractionalPrecision: PreserveUpToHundredth,
		},
		output: output{
			triplified: "45,121,000,000,000",
			final:      "45,121,000,000,000.00",
		},
	},
	{
		stimulus: stimulus{
			value:               "80800000.00",
			groupingCharRune:    ',',
			groupingCharStr:     ",",
			ifSigFig:            true,
			fractionalPrecision: PreserveUpToHundredth,
		},
		output: output{
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
