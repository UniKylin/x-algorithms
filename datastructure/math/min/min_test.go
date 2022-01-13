package min

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	res := Min(1, 2)
	fmt.Println(res)
}

func TestMaxs(t *testing.T) {
	res := Mins(9, 5, 1)
	fmt.Println(res)
}

func BenchmarkMax(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Min(3, 1)
	}
}

func BenchmarkMaxs(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Mins(3, 11, 2, 5, 8, -1)
	}
}
