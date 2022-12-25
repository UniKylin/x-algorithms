package basic

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	ant := make(map[string]string)
	ant["k1"] = "Rob Pike"
	ant["k2"] = "Dannis Ritchie"
	ant["k3"] = "Ken Thompson"
	ant["t1"] = "Test1"
	fmt.Println(">>> map ant: ", ant)

	delete(ant, "t1")
	fmt.Println(">>> map ant: ", ant)

	if val, ok := ant["k3"]; ok {
		fmt.Println(">>> k3: ", val)
	}

	for k, v := range ant {
		fmt.Println(k, v)
	}

	bear := map[string]int{"ant": 88, "bear": 99}
	fmt.Println(">>> bear: ", bear)

}
