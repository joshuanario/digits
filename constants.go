package digits

const (
	FLOAT       = 24
	DOUBLE_PREC = 53
	QUAD_PREC   = 113
	OCTO_PREC   = 237
	PREC_BITS   = QUAD_PREC
)

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

type Precision int

const Exact = Precision(-1 << 31)

const (
	Trillions Precision = iota - 12
	HundredBillions
	TenBillions
	Billions
	HundredMillions
	TenMillions
	Millions
	HundredThousands
	TenThousands
	Thousands
	Hundreds
	Tens
	Oneth
	Tenth
	Hundredth
	Thousandth
	TenThousandth
	HundredThousandth
	Millionth
	TenMillionth
	HundredMillionth
	Billionth
	TenBillionth
	HundredBillionth
	Trillionth
)
