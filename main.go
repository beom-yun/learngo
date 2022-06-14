package main

import "fmt"

func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 50:
		return false
	}
	return true

	// switch {
	// case age < 10:
	// 	return false
	// }
	// return true
}

func main() {
	fmt.Println(canIDrink(18))
}
