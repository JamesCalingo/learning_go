package exercises

import (
	"fmt"
	"math/rand"
)

//EXERCISE 1: generating 100 ints
func oneHundredInts() []int {
	var ints []int
	for i := 0; i < 100; i++ {
		ints = append(ints, rand.Intn(100))
	}
	return ints
}

//EXERCISE 2: looping over the ints
func rules(ints []int) []string {
	var result []string
	for _, value := range ints {
		switch {
		case value % 6 == 0:
			result = append(result, "Six!")
		case value % 3 == 0:
			result = append(result, "Three!")
		case value % 2 == 0:
			result = append(result, "Two!")
		default:
			result = append(result, "Never mind.")
		}
	}
	return result
}

//EXERCISE 3: runningSum and the dangers of shadowing
func runningSum() int {
	var total int
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	return total
	//There's a rather obvious problem as written: the := inside the for loop (which was "assigned" by the book). Instead of taking the old total variable and adding to it, we're creating a new one each time, which sets it to 0 and then adds i to it. Also, since the "final" total in the loop is stuck inside the loop, the total we set at the start of the function (i.e. the zero value) is not changed.
}

func Ch4() {
	ints := oneHundredInts()
	fmt.Println(ints)
	fmt.Println(rules(ints))
	fmt.Println(runningSum())
}
