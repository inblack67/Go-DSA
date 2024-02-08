package recursion

func Factorial(num int32) int32 {
	if num <= 1 {
		return 1
	}

	return num * Factorial(num-1)
}
