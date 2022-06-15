package main

import "fmt"

func main() {
	yun := map[string]string{"name": "yun", "age": "31"}
	for key, value := range yun {
		fmt.Println(key, value)
	}
}
