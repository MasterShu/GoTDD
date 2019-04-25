package lib

import (
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	if Fizzbuzz(15) != "FizzBuzz" {
		t.Error("is not pass 15")
	}
	if Fizzbuzz(5) != "Buzz" {
		t.Error("is not pass 5")
	}
	if Fizzbuzz(3) != "Fizz" {
		t.Error("is not pass 3")
	}

}
