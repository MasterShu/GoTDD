package main

import (
	"fizzbuzz/lib"
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(lib.Fizzbuzz(i))
	}
}
