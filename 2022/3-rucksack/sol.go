package rucksack

import (
	"fmt"
	"log"

	"github.com/smrnjeet222/aoc/2022/utils"
)

func Solve() {
	println("\nGame 3: Lets go ......\n")

	lines, err := utils.ReadLines("./2022/3-rucksack/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0

	for _, sack := range lines {
		grp1 := sack[0 : len(sack)/2]
		grp2 := sack[len(sack)/2:]
		sum += getPriority([]string{grp1, grp2})
	}

	fmt.Println("Solution 1:", sum)

	sum = 0
	var grp []string

	for i, sack := range lines {
		if i%3 == 0 && i != 0 {
			sum += getPriority(grp)
			grp = []string{}
		}
		grp = append(grp, sack)
	}
	// Have to do it for last elemnent also
	sum += getPriority(grp)

	fmt.Println("Solution 2:", sum)
}

// Return the priority of the items combined
// since ques provides only one unique item
// + and loop could be avoided
func getPriority(grp []string) int {
	duplicates := getCommonChars(grp)
	var priority int

	for _, d := range duplicates {
		if d >= 97 && d <= 122 {
			priority += int(d) - 97 + 1
		}
		if d >= 65 && d <= 90 {
			priority += int(d) - 65 + 27
		}
	}

	return priority
}

// INFO: this approach will ignore duplicates in groups
// for eg:
// Input: words = ["bella","label","roller"]
// Output: ["e","l"]   (not ["e","l","l"])

func getCommonChars(grp []string) []rune {
	allChars := make(map[rune]int)

	for _, line := range grp {
		seen := make(map[rune]bool)
		for _, c := range line {
			if !seen[c] {
				allChars[c] += 1
			}
			seen[c] = true
		}
	}

	result := make([]rune, 0)
	for k, v := range allChars {
		if v == len(grp) {
			result = append(result, k)
		}
	}
	return result
}
