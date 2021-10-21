package strings

func internalFunction() {}

func CountOddEven(s string) (odds, evens int) {
	for _, code := range s {
		if code%2 == 0 {
			evens++
			continue
		}
		odds++
	}
	return
}
