package tuningtrouble

import (
	"fmt"
	"log"
	"slices"

	"github.com/smrnjeet222/aoc/2022/utils"
)

func Solve() {
	println("\nGame 6: Let's go .....\n")

	lines, err := utils.ReadLines("./2022/6-tuning-trouble/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	input := lines[0]

	var ans1, ans2 int

	accumulator := make([]rune, 0)
	for i, c := range input {
		if ans1 != 0 && ans2 != 0 {
			break
		}

		if len(accumulator) == 4 && ans1 == 0 {
			ans1 = i
		}

		if len(accumulator) == 14 && ans2 == 0 {
			ans2 = i
		}

		index := slices.Index(accumulator, c)

		if index != -1 {
			accumulator = accumulator[index+1:]
		}
		accumulator = append(accumulator, c)
	}

	fmt.Println("Solution 1:", ans1)
	fmt.Println("Solution 2:", ans2)
}
