package basic

import (
	"fmt"
	"testing"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func TestRecursion(t *testing.T) {
	fmt.Println(">>> fact: ", fact(8))

	// closure
	var fib func(n int) int

	fib = func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(">>> fib :", fib(8))
}
