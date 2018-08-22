package offer

func Fibonacci() func() int  {
	first, second := 0, 1
	return func() int {
		first, second = second, first + second
		return  first
	}
}
