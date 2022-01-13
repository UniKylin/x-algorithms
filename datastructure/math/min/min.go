package min

func Min(ant, bear int) int {
	if ant > bear {
		return bear
	}
	return ant
}

// Return the minimum of all the integers
func Mins(values ...int) (res int) {
	res = values[0]

	for _, val := range values {
		if val < res {
			res = val
		}
	}

	return
}
