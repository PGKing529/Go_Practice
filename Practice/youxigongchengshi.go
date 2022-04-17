package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Scanln()
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &matrix[i][j])
		}
		fmt.Scanln()
	}

	for len(matrix) != 0 {
		i, j := core(matrix)
		matrix = minus(matrix, i, j)
		fmt.Println(i+1, j+1)
	}
}

func core(matrix [][]int) (int, int) {
	row := len(matrix)
	column := len(matrix[0])
	if row == 1 && column == 1 {
		return 0, 0
	}
	matrix_row := make([]int, row)
	matrix_column := make([]int, column)
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			matrix_row[i] += matrix[i][j]
		}
	}
	for j := 0; j < column; j++ {
		for i := 0; i < row; i++ {
			matrix_column[j] += matrix[i][j]
		}
	}

	res := [2]int{}
	sum := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if sum < (matrix_row[i] + matrix_column[j] - matrix[i][j]) {
				sum = (matrix_row[i] + matrix_column[j] - matrix[i][j])
				res[0], res[1] = i, j
			}
		}
	}
	return res[0], res[1]
}

func minus(matrix [][]int, h, s int) [][]int {
	row := len(matrix)
	column := len(matrix[0])
	if row == 1 && column == 1 {
		return nil
	}
	res := make([][]int, row-1)
	for i := 0; i < row-1; i++ {
		res[i] = make([]int, column-1)
	}

	for i := 0; i < row; i++ {
		x := 0
		if i == h {
			continue
		} else if i > h {
			x = i - 1
		} else if i < h {
			x = i
		}
		for j := 0; j < column; j++ {
			if j == s {
				continue
			} else if j < s {
				res[x][j] = matrix[i][j]
			} else if j > s {
				res[x][j-1] = matrix[i][j]
			}
		}
	}
	return res
}
