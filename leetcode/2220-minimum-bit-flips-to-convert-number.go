package main

func minBitFlips(start int, goal int) int {
	bitflips := 0
	for start != 0 || goal != 0 {
		bitflips += (start & 1) ^ (goal & 1)
		start = start >> 1
		goal = goal >> 1
	}
	return bitflips
}
