package iteration

const defaultRepeateCount = 5

// Repeat takes a string and a repeat times and returns the repeated string.
func Repeat(character string, expectedCount int) string {
	var repeated string
	var repeateCount int
	if expectedCount > 0 {
		repeateCount = expectedCount
	} else {
		repeateCount = defaultRepeateCount
	}
	for i := 0; i < repeateCount; i++ {
		repeated += character
	}
	return repeated
}
