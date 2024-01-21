package main

import (
	"fmt"
	"os"
)

func SolveSudoku(grid [][]int) [][]int {
	// check if the grid is valid
	if !isValid(grid) {
		fmt.Println("Invalid grid")
		os.Exit(1)
	}
	// solve the grid
	solve(grid)
	return grid
}

func isValid(grid [][]int) bool {
	// check if the grid is 9x9
	if len(grid) != 9 {
		return false
	}
	for _, row := range grid {
		if len(row) != 9 {
			return false
		}
	}
	// check if the grid contains only numbers from 0 to 9
	for _, row := range grid {
		for _, num := range row {
			if num < 0 || num > 9 {
				return false
			}
		}
	}
	return true
}

func solve(grid [][]int) bool {
	// find empty cell
	row, col := findEmptyCell(grid)
	// if there is no empty cell, the grid is solved
	if row == -1 && col == -1 {
		return true
	}
	// try to fill the empty cell with a number
	for num := 1; num <= 9; num++ {
		// check if the number is valid
		if isValidNum(grid, row, col, num) {
			// fill the cell
			grid[row][col] = num
			// solve the grid recursively
			if solve(grid) {
				return true
			}
			// if the grid can't be solved, reset the cell
			grid[row][col] = 0
		}
	}
	return false
}

func findEmptyCell(grid [][]int) (int, int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}

const BoxSize = 3

func isValidNum(grid [][]int, row, col, num int) bool {
	// check if the number is already in the row
	for _, currentNum := range grid[row] {
		if currentNum == num {
			return false
		}
	}
	// check if the number is already in the column
	for _, currentRow := range grid {
		if currentRow[col] == num {
			return false
		}
	}
	// check if the number is already in the box
	boxRowStart := row - row%BoxSize
	boxColStart := col - col%BoxSize
	for i := boxRowStart; i < boxRowStart+BoxSize; i++ {
		for j := boxColStart; j < boxColStart+BoxSize; j++ {
			if grid[i][j] == num {
				return false
			}
		}
	}
	return true
}
