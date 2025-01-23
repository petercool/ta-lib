package utils

import "math"

// Abs returns the absolute value of x
func Abs(x float64) float64 {
	return math.Abs(x)
}

// RoundPos rounds a positive number
func RoundPos(x float64) float64 {
	return math.Floor(x + 0.5)
}

// RoundNeg rounds a negative number
func RoundNeg(x float64) float64 {
	return math.Ceil(x - 0.5)
}

// RoundPos2 rounds a positive number to 2 decimal places
func RoundPos2(x float64) float64 {
	return math.Floor(x*100.0+0.5) / 100.0
}

// RoundNeg2 rounds a negative number to 2 decimal places
func RoundNeg2(x float64) float64 {
	return math.Ceil(x*100.0-0.5) / 100.0
}

// IsZero checks if a float64 is approximately zero
func IsZero(x float64) bool {
	return math.Abs(x) < Epsilon
}

// AreEqual checks if two float64 values are approximately equal
func AreEqual(a, b float64) bool {
	return math.Abs(a-b) < Epsilon
}

// Max returns the maximum of two float64 values
func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of two float64 values
func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// MaxInSlice returns the maximum value in a slice of float64
func MaxInSlice(values []float64) float64 {
	if len(values) == 0 {
		return math.NaN()
	}
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// MinInSlice returns the minimum value in a slice of float64
func MinInSlice(values []float64) float64 {
	if len(values) == 0 {
		return math.NaN()
	}
	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// Sum returns the sum of all values in a slice
func Sum(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum
}

// Mean calculates the arithmetic mean of a slice of values
func Mean(values []float64) float64 {
	if len(values) == 0 {
		return math.NaN()
	}
	return Sum(values) / float64(len(values))
}

// Variance calculates the variance of a slice of values
func Variance(values []float64, mean float64) float64 {
	if len(values) == 0 {
		return math.NaN()
	}

	var sum float64
	for _, v := range values {
		diff := v - mean
		sum += diff * diff
	}
	return sum / float64(len(values))
}

// StdDev calculates the standard deviation of a slice of values
func StdDev(values []float64) float64 {
	if len(values) == 0 {
		return math.NaN()
	}
	mean := Mean(values)
	return math.Sqrt(Variance(values, mean))
}
