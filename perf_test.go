package digits_test

import (
	"testing"

	"github.com/joshuanario/digits"
)

func BenchmarkSigFigs(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, testCase := range testCases {
			_, err := digits.New(testCase.stimulus.precision, testCase.stimulus.value, testCase.stimulus.groupSeparator, testCase.stimulus.fractionalPrecision)
			if err != nil {
				b.Fatalf(err.Error())
			}
		}
	}
}
