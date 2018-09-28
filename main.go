// 8 QUEENS PUZZLE
package main

import (
	"fmt"
	"strings"
)

const n int = 8

var total int

func main() {
	var queens [][2]int
	place(queens)
}

func place(queens [][2]int) {
	row := len(queens)
	if row < n {
		chessboard := render(queens)
		for col := 0; col < n; col++ {
			if chessboard[row][col] == 0 {
				queens = append(queens, [2]int{row, col})
				place(queens)
				queens = queens[:row]
			}
		}
	} else {
		total++
		visualize(total, queens)
	}
}

func render(queens [][2]int) [n][n]int {
	var chessboard [n][n]int
	for _, queen := range queens {
		for row := 0; row < n; row++ {
			for col := 0; col < n; col++ {
				if row == queen[0] && col == queen[1] {
					chessboard[row][col] = 1
				} else if row == queen[0] ||
					col == queen[1] ||
					row-queen[0] == col-queen[1] ||
					row+col == queen[0]+queen[1] {
					chessboard[row][col] = -1
				}
			}
		}
	}
	return chessboard
}

func visualize(index int, solution [][2]int) {
	fmt.Println("Solution ", index)
	fmt.Println(strings.Repeat("-", 2*n-1))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == solution[i][0] && j == solution[i][1] {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		println()
	}
	fmt.Println(strings.Repeat("-", 2*n-1))
}
