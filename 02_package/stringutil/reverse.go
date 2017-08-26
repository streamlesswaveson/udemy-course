package stringutil

// uppercase function names indicate exportable
func Reverse(s string) string {
	// lowercase function names are NOT exportable
	return reverseTwo(s)
}
