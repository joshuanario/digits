package digits_test

import (
	"testing"

	"github.com/joshuanario/digits"
)

func BenchmarkSigFigs(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, sut := range suts {
			_, err := digits.New(sut.stimulus.precision, sut.stimulus.value, sut.stimulus.groupSeparator, sut.stimulus.fractionalPrecision)
			if err != nil {
				b.Fatalf(err.Error())
			}
		}
	}
}
