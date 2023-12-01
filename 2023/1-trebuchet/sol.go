package trebuchet

import (
	"fmt"
	"strings"

	"github.com/smrnjeet222/aoc/2022/utils"
)

func Solve() {
	println("\nDay 1 : Lets Go........\n")

	lines, err := utils.ReadLines("./2023/1-trebuchet/input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0

	for _, line := range lines {
		var str string

		for _, w := range line {
			if w <= 58 && w >= 48 {
				if len(str) == 2 {
					str = str[0:1]
				}
				str += string(w)
			}
		}

		if len(str) == 1 {
			str += str[0:1]
		}

		sum += utils.MustStrConv(str)
	}

	fmt.Printf("Solution 1 : %v\n", sum)

	Solve2(lines)
}

func Solve2(lines []string) {
	sum := 0

	for _, line := range lines {
		var str string

		line = strings.ReplaceAll(line, "one", "o1ne")
		line = strings.ReplaceAll(line, "two", "t2wo")
		line = strings.ReplaceAll(line, "three", "t3hree")
		line = strings.ReplaceAll(line, "four", "f4our")
		line = strings.ReplaceAll(line, "five", "f5ive")
		line = strings.ReplaceAll(line, "six", "s6ix")
		line = strings.ReplaceAll(line, "seven", "s7even")
		line = strings.ReplaceAll(line, "eight", "e8ight")
		line = strings.ReplaceAll(line, "nine", "n9ine")

		for _, w := range line {
			if w <= 58 && w >= 48 {
				if len(str) == 2 {
					str = str[0:1]
				}
				str += string(w)
			}
		}

		if len(str) == 1 {
			str += str[0:1]
		}

		sum += utils.MustStrConv(str)
	}

	fmt.Printf("Solution 2 : %v\n", sum)
}
