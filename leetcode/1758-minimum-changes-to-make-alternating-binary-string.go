package main

func minOperations(s string) int {
	zeroOnes := 0
	oneZeroes := 0
	for i, char := range s {
		switch {
		case char == '0' && i%2 == 0:
			zeroOnes++
		case char == '0' && i%2 != 0:
			oneZeroes++
		case char == '1' && i%2 == 0:
			oneZeroes++
		case char == '1' && i%2 != 0:
			zeroOnes++
		}
	}
	return len(s) - max(zeroOnes, oneZeroes)
}
