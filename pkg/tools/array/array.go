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
