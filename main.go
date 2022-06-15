package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	yun := person{name: "yun", age: 31, favFood: favFood}
	fmt.Println(yun)
}
