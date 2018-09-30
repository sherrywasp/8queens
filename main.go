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
	y := len(queens)
	if y < n {
		nextRank := render(queens)
		for x := 0; x < n; x++ {
			if !nextRank[x] {
				queens = append(queens, [2]int{x, y})
				place(queens)
				queens = queens[:y]
			}
		}
	} else {
		total++
		visualize(total, queens)
	}
}

func render(queens [][2]int) [n]bool {
	var rank [n]bool
	y := len(queens)
	for x := 0; x < n; x++ {
		for _, queen := range queens {
			if x-queen[0] == y-queen[1] ||
				x == queen[0] ||
				x+y == queen[0]+queen[1] {
				rank[x] = true
				break
			}
		}
	}
	return rank
}

func visualize(index int, solution [][2]int) {
	fmt.Println("Solution ", index)
	fmt.Println(strings.Repeat("-", 2*n-1))
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			if x == solution[y][0] && y == solution[y][1] {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		println()
	}
	fmt.Println(strings.Repeat("-", 2*n-1))
}
