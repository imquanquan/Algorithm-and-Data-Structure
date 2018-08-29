package offer

func IsEven(num int) bool {
	return (num % 2) == 0
}

func ReOrder(nums []int, key func(int) bool) {
	left, right := 0, len(nums) - 1
	for left < right {
		for ! key(nums[left]) {
			left ++
		}
		for key(nums[right]) {
			right --
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
}