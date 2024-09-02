package utils

// StringPtr returns a pointer to the given string value
func StringPtr(s string) *string {
	return &s
}

// IntPtr returns a pointer to the given int value
func IntPtr(i int) *int {
	return &i
}

// Int64Ptr returns a pointer to the given int64 value
func Int64Ptr(i int64) *int64 {
	return &i
}

// UintPtr returns a pointer to the given uint value
func UintPtr(u uint) *uint {
	return &u
}
