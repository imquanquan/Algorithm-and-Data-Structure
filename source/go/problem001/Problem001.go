package problem001

type Matrix [][]int

func FindInt(X Matrix, key int)(flag bool) {
	rows := len(X)
	cols := len(X[0])
	if rows < 0 || cols < 0 {
		return false
	}
	row := 0
	col := cols - 1
	for row < rows && col >=0 {
		if X[row][col] == key {
			flag = true
			break
		} else if X[row][col] > key {
			col --
		} else {
			row ++
		}
	}
	return
}
