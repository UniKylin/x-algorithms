// link		: https://leetcode-cn.com/problems/divide-two-integers/
// Author	: Kylin
// Date		: 2022-02-02

package leetcode

import (
	"math"
)

func Divide(dividend int, divisor int) int {
	// fmt.Println(math.MinInt32)
	// fmt.Println(math.MaxInt32)

	// 当 -2^31 / -1 会溢出，这是因为最大的int32是 2^31 - 1
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	negative := 2

	if dividend > 0 {
		negative--
		dividend = -dividend
	}

	if divisor > 0 {
		negative--
		divisor = -divisor
	}

	result := div(dividend, divisor)

	if negative == 1 {
		return -result
	} else {
		return result
	}
}

func div(dividend, divisor int) int {
	res := 0

	for dividend <= divisor {
		value := divisor
		quotient := 1

		for value >= math.MinInt32/2 && dividend <= value+value {
			quotient += quotient
			value += value
		}

		res += quotient
		dividend -= value
	}

	return res
}
