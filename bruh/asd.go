package bruh

func Add(x, y int) int {
	return x + y
}

func Substract(x, y int) int {
	return x - y
}

func Divide(x, y int) int {
	if y == 0 {
		return 0
	}
	return x / y
}

func Product(x, y int) int {
	return x * y
}
