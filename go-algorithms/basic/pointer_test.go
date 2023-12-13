package basic

import (
	"fmt"
	"testing"
)

func setVal(val int) {
	val = 88
}

func setPtr(val *int) {
	*val = 999
}

func TestPointers(t *testing.T) {
	ant := 0
	fmt.Println(">>> ant val: ", ant)

	setVal(ant)
	fmt.Println(">>> set ant val: ", ant)

	setPtr(&ant)
	fmt.Println(">>> set ant val using pointer: ", ant)

	fmt.Println(">>> ant pointer address: ", &ant)
}
