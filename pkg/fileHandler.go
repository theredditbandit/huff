package pkg

// GetStats returns the character count of the input data
func GetStats(data []byte) map[string]int {
	charCount := make(map[string]int)
	for _, char := range data {
		charCount[string(char)]++
	}
	return charCount
}
