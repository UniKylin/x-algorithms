package basic

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	ant := make([]string, 3)
	fmt.Println(">>> empty ant: ", ant)

	ant[0] = "A"
	ant[1] = "N"
	ant[2] = "T"
	fmt.Println(">>> Upper ant: ", ant)
	fmt.Println(">>> len ant: ", len(ant))

	ant = append(ant, "b")
	ant = append(ant, "e", "a", "r")
	fmt.Println(">>> ant append : ", ant)

	bear := make([]string, len(ant))
	copy(bear, ant)
	fmt.Println(">>> bear: ", bear)

	cat := bear[2:4]
	fmt.Println(">>> bear slice: ", cat)

	dog := bear[1:]
	fmt.Println(">>> dog slice: ", dog)

	egg := []string{"ant", "dog", "cat"}
	fmt.Println(">>> string slice egg: ", egg)
}
