package utils

// ValidateParams validates common parameters for technical indicators
func ValidateParams(startIdx, endIdx int, inReal []float64, optInTimePeriod int) (RetCode, int, int) {
	// Check for minimum number of elements
	if len(inReal) == 0 {
		return InvalidParameter, 0, 0
	}

	// Check for valid index range
	if startIdx < 0 {
		return OutOfRangeStartIndex, 0, 0
	}
	if endIdx < startIdx {
		return OutOfRangeEndIndex, 0, 0
	}
	if endIdx >= len(inReal) {
		return OutOfRangeEndIndex, 0, 0
	}

	// Check for valid time period
	if optInTimePeriod <= 0 {
		return InvalidParameter, 0, 0
	}

	// Calculate output start index and number of elements
	outBegIdx := startIdx
	outNBElement := endIdx - startIdx + 1

	return Success, outBegIdx, outNBElement
}

// ValidatePrice validates price data arrays for technical indicators
func ValidatePrice(startIdx, endIdx int, high, low, close []float64) (RetCode, int, int) {
	// Check for minimum number of elements
	if len(high) == 0 || len(low) == 0 || len(close) == 0 {
		return InvalidParameter, 0, 0
	}

	// Check for equal length arrays
	if len(high) != len(low) || len(high) != len(close) {
		return InvalidParameter, 0, 0
	}

	// Check for valid index range
	if startIdx < 0 {
		return OutOfRangeStartIndex, 0, 0
	}
	if endIdx < startIdx {
		return OutOfRangeEndIndex, 0, 0
	}
	if endIdx >= len(high) {
		return OutOfRangeEndIndex, 0, 0
	}

	// Calculate output start index and number of elements
	outBegIdx := startIdx
	outNBElement := endIdx - startIdx + 1

	return Success, outBegIdx, outNBElement
}

// ValidateVolume validates price and volume data arrays for technical indicators
func ValidateVolume(startIdx, endIdx int, price, volume []float64) (RetCode, int, int) {
	// Check for minimum number of elements
	if len(price) == 0 || len(volume) == 0 {
		return InvalidParameter, 0, 0
	}

	// Check for equal length arrays
	if len(price) != len(volume) {
		return InvalidParameter, 0, 0
	}

	// Check for valid index range
	if startIdx < 0 {
		return OutOfRangeStartIndex, 0, 0
	}
	if endIdx < startIdx {
		return OutOfRangeEndIndex, 0, 0
	}
	if endIdx >= len(price) {
		return OutOfRangeEndIndex, 0, 0
	}

	// Calculate output start index and number of elements
	outBegIdx := startIdx
	outNBElement := endIdx - startIdx + 1

	return Success, outBegIdx, outNBElement
}

// ValidateMAType validates the moving average type
func ValidateMAType(maType MAType) RetCode {
	if maType < SMA || maType > MAMA {
		return InvalidParameter
	}
	return Success
}
