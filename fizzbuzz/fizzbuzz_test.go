package fizzbuzz

import (
	"strconv"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("is Fizz", testFizzBuzz(3, "Fizz"))
	t.Run("is Fizz", testFizzBuzz(13, "Fizz"))
	t.Run("is Buzz", testFizzBuzz(5, "Buzz"))
	t.Run("is Buzz", testFizzBuzz(151, "Buzz"))
	t.Run("is number", testFizzBuzz(11, "11"))
}

func testFizzBuzz(num int, expected string) func(t *testing.T) {
	return func(t *testing.T) {
		if FizzBuzz(num) != expected {
			t.Error("is not pass" + strconv.Itoa(num))
		}
	}
}
