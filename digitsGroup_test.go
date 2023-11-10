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
type SUT struct {
	stimulus Stimulus
	output   Output
}

var suts = []*SUT{
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
	for _, sut := range suts {
		output := DigitGroup(Exact, sut.stimulus.value, sut.stimulus.groupingCharRune, sut.stimulus.fractionalPrecision, sut.stimulus.ifSigFig)
		if output != sut.output.final {
			t.Fail()
		}
	}
}

func Test_Triplefy(t *testing.T) {
	for _, sut := range suts {
		var output = triplefy(sut.stimulus.value, sut.stimulus.groupingCharStr, sut.stimulus.ifSigFig)
		if output != sut.output.triplified {
			t.Fail()
		}
	}
}

func Test_ZeroTriplefy(t *testing.T) {
	//TODO: flesh this out in stimulus
	for _, sut := range suts {
		var output = zeroTriplefy(sut.stimulus.value, sut.stimulus.groupingCharStr)
		if output != sut.output.triplified {
			t.Fail()
		}
	}
}
*/
