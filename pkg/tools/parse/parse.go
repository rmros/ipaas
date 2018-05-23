/*
Copyright [yyyy] [name of copyright owner]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
