package printsum

import "fmt"

func PrintSum(x int, n int) {
	if n == 0 { // Nur ein = statt ==
		fmt.Println(x) // kein int bei Übergabe von x
		return         // nur return ohne x
	}
	PrintSum(x+1, n-1)
}
