package problem008

func InOrderFind(a []int) int {
	result := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] < result {
			result = a[i]
		}
	}
	return result
}

func FindMinInRotateArray(a []int) int {
	if len(a) == 0 {
		return -1
	}
	mid, left, right := 0, 0, len(a) - 1

	for a[left] >= a[right] {
		if right - left == 1 {
			return a[right]
		}
		mid := (left + right) / 2
		if a[left] == a[mid] && a[mid] == a[right] {
			return InOrderFind(a)
		}
		if a[left] <= a[mid] {
			left = mid
		}
		if a[right] >= a[mid] {
			right = mid
		}
	}
	return a[mid]
}
