package addelements

import "fmt"

func ExampleAddElements() {
	l1 := []int{1, 2, 3, 4, 5}

	fmt.Println(AddElements(l1))

	// Output:
	// 15
}
