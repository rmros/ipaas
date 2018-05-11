package parse

import (
	"strconv"
)

// IntToInt32Pointer translator int to *int32
func IntToInt32Pointer(input int) *int32 {
	output := new(int32)
	*output = int32(input)
	return output
}

// IntToInt64Pointer translator int to *int64
func IntToInt64Pointer(input int) *int64 {
	output := new(int64)
	*output = int64(input)
	return output
}

// BoolToPointer translator bool to *bool
func BoolToPointer(input bool) *bool {
	output := new(bool)
	*output = input
	return output
}

// StringToPointer translator string to *string
func StringToPointer(input string) *string {
	output := new(string)
	*output = input
	return output
}

// StringToInt32Pointer translator string to *int32
func StringToInt32Pointer(input string) *int32 {
	output := new(int32)
	tmp, _ := strconv.Atoi(input)
	*output = int32(tmp)
	return output
}
