/*
Copyright 2018 huangjia.

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

package array

// IntNotIn assert ojb in the des int array or not
func IntNotIn(des []int, obj int) bool {
	for _, item := range des {
		if item == obj {
			return true
		}
	}
	return false
}

// StringNotIn assert ojb in the des  string array or not
func StringNotIn(des []string, obj string) bool {
	for _, item := range des {
		if item == obj {
			return false
		}
	}
	return true
}
