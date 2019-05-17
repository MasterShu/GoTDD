package fizzbuzz

import (
	"strconv"
	"strings"
)

func FizzBuzz(num int) string {
	if isFizz(num) {
		return "Fizz"
	}
	if isBuzz(num) {
		return "Buzz"
	}
	return strconv.Itoa(num)
}

func isBuzz(num int) bool {
	return num%5 == 0 || strings.Index(strconv.Itoa(num), "5") > 0
}

func isFizz(num int) bool {
	return num%3 == 0 || strings.Index(strconv.Itoa(num), "3") > -1
}
