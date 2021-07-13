package digits_test

import (
	"testing"

	"github.com/joshuanario/digits"
)

func BenchmarkSigFigs(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, s := range stimuli {
			_, err := digits.New(s.p, s.v, s.g, s.d)
			if err != nil {
				b.Fatalf(err.Error())
			}
		}
	}
}
