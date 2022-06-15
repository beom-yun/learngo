package main

import "fmt"

func main() {
	names := []string{"yun", "kim", "park"}
	names = append(names, "jeong")
	fmt.Println(names)
}
