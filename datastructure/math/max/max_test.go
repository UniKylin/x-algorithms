package max

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	res := Max(1, 2)
	fmt.Println(res)
}

func TestMaxs(t *testing.T) {
	res := Maxs(9, 5, 1)
	fmt.Println(res)
}

func BenchmarkMax(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Max(3, 1)
	}
}

func BenchmarkMaxs(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Maxs(3, 11, 2, 5, 8, -1)
	}
}
