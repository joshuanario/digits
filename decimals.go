package digits

type Decimals uint

const (
	NoDecimals Decimals = iota
	PreserveUpToTenth
	PreserveUpToHundredth
	PreserveUpToThousandth
	PreserveUpToTenThousandth
	PreserveUpToHundredThousandth
	PreserveUpToMillionth
	PreserveUpToTenMillionth
	PreserveUpToHundredMillionth
	PreserveUpToBillionth
	PreserveUpToTenBillionth
	PreserveUpToHundredBillionth
	PreserveUpToTrillionth
	MaximumDecimals = PreserveUpToTrillionth
)
