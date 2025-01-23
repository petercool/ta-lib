package utils

// MAType represents the Moving Average type
type MAType int

const (
	SMA   MAType = iota // Simple Moving Average
	EMA                 // Exponential Moving Average
	WMA                 // Weighted Moving Average
	DEMA                // Double Exponential Moving Average
	TEMA                // Triple Exponential Moving Average
	TRIMA               // Triangular Moving Average
	KAMA                // Kaufman Adaptive Moving Average
	MAMA                // MESA Adaptive Moving Average
)

// RetCode represents the return code for indicator calculations
type RetCode int

const (
	Success RetCode = iota
	InvalidParameter
	OutOfRangeStartIndex
	OutOfRangeEndIndex
	AllocError
	InternalError
)

// Result represents the output of an indicator calculation
type Result struct {
	BeginIndex int       // Index of first valid data
	NBElement  int       // Number of elements in the result
	Values     []float64 // Output values
}

// MAResult represents the output of a moving average calculation
type MAResult struct {
	Result
	MAType MAType // Type of moving average used
}

// MACDResult represents the output of MACD calculations
type MACDResult struct {
	Result
	MACDSignal []float64 // MACD signal line
	MACDHist   []float64 // MACD histogram
}

// MinMaxResult represents minimum and maximum values
type MinMaxResult struct {
	Result
	Minimums []float64
	Maximums []float64
}

// Constants for common calculations
const (
	Epsilon = 0.000000001 // Small number for floating point comparisons
)
