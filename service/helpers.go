package service

func ternary(a bool, b, c interface{}) interface{} {
	if a {
		return b
	}

	return c
}
