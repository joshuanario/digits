package digits

import (
	"math/big"
)

func computeHead(value *big.Float) string {
	sign := value.Sign()
	if sign < 0 {
		return "("
	}
	return ""
}
