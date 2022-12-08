package max

func Max(ant, bear int) int {
	if ant > bear {
		return ant
	}
	return bear
}

// Return the maximum of all the integers
func Maxs(values ...int) (res int) {
	res = values[0]

	for _, val := range values {
		if val > res {
			res = val
		}
	}

	return
}
