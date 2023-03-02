package mersenne

// IsMersenne erwartet eine Zahl und prüft, ob diese eine Mersenne-Zahl ist.
// Mersenne-Zahlen sind natürliche Zahlen von der Form 2^n + 1.
func IsMersenne(x int) bool {
	// BEGIN-SOLUTION
	for i := 1; i-1 <= x; i *= 2 {
		if i-1 == x {
			return true
		}
	}
	return false
	// END-SOLUTION
}
