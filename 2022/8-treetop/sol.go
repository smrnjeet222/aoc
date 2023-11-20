package treetop

import (
	"fmt"
	"log"

	"github.com/smrnjeet222/aoc/2022/utils"
)

func Solve() {
	println("\nGame 8: Let's go .....\n")

	lines, err := utils.ReadLines("./2022/8-treetop/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	trees, seenTrees := ParseTrees(lines)

	// Store the MaxTree for each row
	maxFromTop := make([]int, len(trees))
	maxFromBottom := make([]int, len(trees))

	for i, a := 0, len(trees)-1; i < len(trees) && a >= 0; i, a = i+1, a-1 {
		// Store the MaxTree for each column
		// each look will create its new instrance
		// so array not required
		var maxFromLeft, maxFromRight int

		for j, b := 0, len(trees[i])-1; j < len(trees[i]) && b >= 0; j, b = j+1, b-1 {
			// Start looking from left
			if j == 0 || (trees[i][j] > maxFromLeft) {
				if trees[i][j] > maxFromLeft {
					maxFromLeft = trees[i][j]
				}
				seenTrees[i][j] = true
			}

			// Start looking from right
			if b == (len(trees[i])-1) || (trees[i][b] > maxFromRight) {
				if trees[i][b] > maxFromRight {
					maxFromRight = trees[i][b]
				}
				seenTrees[i][b] = true
			}

			// Start looking from top
			if i == 0 || (trees[i][j] > maxFromTop[j]) {
				if trees[i][j] > maxFromTop[j] {
					maxFromTop[j] = trees[i][j]
				}
				seenTrees[i][j] = true
			}

			// Start looking from bottom
			if a == (len(trees)-1) || (trees[a][j] > maxFromBottom[j]) {
				if trees[a][j] > maxFromBottom[j] {
					maxFromBottom[j] = trees[a][j]
				}
				seenTrees[a][j] = true
			}
		}
	}

	count := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			if seenTrees[i][j] {
				// fmt.Print("T")
				count++
			} else {
				// fmt.Print("F")
			}
		}
		// fmt.Println()
	}

	fmt.Println("Solution 1:", count)
}

// Convert the input into a 2D array
func ParseTrees(lines []string) ([][]int, [][]bool) {
	grid := make([][]int, len(lines))
	seenTrees := make([][]bool, len(lines))

	for i, line := range lines {
		grid[i] = make([]int, len(line))
		seenTrees[i] = make([]bool, len(line))
		for j, t := range line {
			grid[i][j] = int(t - '0')
			seenTrees[i][j] = false
		}
	}
	return grid, seenTrees
}
