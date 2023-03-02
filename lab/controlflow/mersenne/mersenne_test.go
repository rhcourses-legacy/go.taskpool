package mersenne

import "fmt"

func ExampleIsMersenne() {
	fmt.Println(IsMersenne(0))
	fmt.Println(IsMersenne(1))
	fmt.Println(IsMersenne(3))
	fmt.Println(IsMersenne(7))
	fmt.Println(IsMersenne(15))
	fmt.Println(IsMersenne(31))
	fmt.Println(IsMersenne(63))
	fmt.Println(IsMersenne(2))
	fmt.Println(IsMersenne(4))
	fmt.Println(IsMersenne(6))
	fmt.Println(IsMersenne(8))

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// false
	// false
	// false
	// false
}
