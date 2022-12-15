package utils

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
