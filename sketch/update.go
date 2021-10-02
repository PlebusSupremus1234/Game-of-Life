package sketch

func countNeighbors(array [100][100]int, y int, x int) int {
	sum := 0

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			col := (y + i + 100) % 100
			row := (x + j + 100) % 100

			sum += array[col][row]
		}
	}

	sum -= array[y][x]

	return sum
}

func Update(array *[100][100]int) {
	next := [100][100]int {}

	for i, subarray := range array {
		for j, e := range subarray {
			amount := countNeighbors(*array, i, j)

			if e == 0 && amount == 3 {
				next[i][j] = 1
			} else if e == 1 && (amount < 2 || amount > 3) {
				next[i][j] = 0
			} else {
				next[i][j] = e
			}
		}
	}

	*array = next
}