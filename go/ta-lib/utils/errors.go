package utils

import "errors"

// Common errors
var (
	ErrInvalidParameter       = errors.New("invalid parameter")
	ErrOutOfRangeStartIndex   = errors.New("start index out of range")
	ErrOutOfRangeEndIndex     = errors.New("end index out of range")
	ErrAllocFailed            = errors.New("memory allocation failed")
	ErrInternalError          = errors.New("internal error")
	ErrEmptyInputData         = errors.New("empty input data")
	ErrMismatchedInputLengths = errors.New("mismatched input lengths")
)
