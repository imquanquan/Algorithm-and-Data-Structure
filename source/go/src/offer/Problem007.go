package offer

func CountOne(number int) int {
	count := 0
	for number != 0 {
		count++
		number = number & (number - 1)
	}
	return count
}
