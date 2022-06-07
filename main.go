package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLength, upperName := lenAndUpper("y u n")
	fmt.Println(totalLength, upperName)

	repeatMe("kim", "yun", "song", "park", "jeong")
}
