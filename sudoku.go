package main

import (
	"github.com/01-edu/z01"
)

func solveSudoku(grid [][]rune) bool {
	row, col := findEmptyCell(grid)
	if row == -1 && col == -1 {
		return true
	}

	for ch := '1'; ch <= '9'; ch++ {
		if isSafe(grid, row, col, ch) {
			grid[row][col] = ch
			if solveSudoku(grid) {
				return true
			}
			grid[row][col] = '.'
		}
	}

	return false
}

func findEmptyCell(grid [][]rune) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func isSafe(grid [][]rune, row, col int, ch rune) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == ch || grid[i][col] == ch || grid[3*(row/3)+i/3][3*(col/3)+i%3] == ch {
			return false
		}
	}
	return true
}

func isValidSudoku(grid [][]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ch := grid[i][j]
			if ch != '.' && (ch < '1' || ch > '9') {
				return false
			}
		}
	}
	return true
}

func printGrid(grid [][]rune) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(grid[i][j])
			if j < 8 {
				z01.PrintRune(' ')
			}
		}

		z01.PrintRune('\n')
	}
}
