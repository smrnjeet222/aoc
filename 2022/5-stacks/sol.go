package supplyStacks

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/smrnjeet222/aoc/2022/utils"
)

//
//             [J] [Z] [G]
//             [Z] [T] [S] [P] [R]
// [R]         [Q] [V] [B] [G] [J]
// [W] [W]     [N] [L] [V] [W] [C]
// [F] [Q]     [T] [G] [C] [T] [T] [W]
// [H] [D] [W] [W] [H] [T] [R] [M] [B]
// [T] [G] [T] [R] [B] [P] [B] [G] [G]
// [S] [S] [B] [D] [F] [L] [Z] [N] [L]
//  1   2   3   4   5   6   7   8   9
//

var cargo = [][]string{
	{"S", "T", "H", "F", "W", "R"},
	{"S", "G", "D", "Q", "W"},
	{"B", "T", "W"},
	{"D", "R", "W", "T", "N", "Q", "Z", "J"},
	{"F", "B", "H", "G", "L", "V", "T", "Z"},
	{"L", "P", "T", "C", "V", "B", "S", "G"},
	{"Z", "B", "R", "T", "W", "G", "P"},
	{"N", "G", "M", "T", "C", "J", "R"},
	{"L", "G", "B", "W"},
}

func Solve() {
	println("\nGame 5: Let's go .....\n")

	lines, err := utils.ReadLines("./2022/5-stacks/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	newCargo1 := moveCrates(cargo, lines, true)
	result := ""
	for _, cargo := range newCargo1 {
		// fmt.Println(cargo)
		result += cargo[len(cargo)-1]
	}

	fmt.Println("Solution 1:", result)

	newCargo2 := moveCrates(cargo, lines, false)
	result = ""
	for _, cargo := range newCargo2 {
		// fmt.Println(cargo)
		result += cargo[len(cargo)-1]
	}

	fmt.Println("Solution 2:", result)
}

// takes the original and returs modified copy
func moveCrates(cargo [][]string, steps []string, moveOneByOne bool) [][]string {
	duplicate := make([][]string, len(cargo))
	for i := range cargo {
		duplicate[i] = make([]string, len(cargo[i]))
		copy(duplicate[i], cargo[i])
	}

	for _, step := range steps {
		split := strings.Split(step, " ")
		amount, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])

		from = from - 1
		to = to - 1

		toMove := duplicate[from][len(duplicate[from])-amount:]
		duplicate[from] = duplicate[from][:len(duplicate[from])-amount]

		if moveOneByOne {
			slices.Reverse(toMove)
		}
		duplicate[to] = append(duplicate[to], toMove...)
	}

	return duplicate
}
