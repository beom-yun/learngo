package main

import "fmt"

func superAdd(numbers ...int) int {
	result := 0
	for _, n := range numbers {
		result += n
	}
	return result
}

func main() {
	total := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)
}
