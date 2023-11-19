package campCleanup

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/smrnjeet222/aoc/2022/utils"
)

func Solve() {
	println("\nGame 4 : Lets go ......\n")
	lines, err := utils.ReadLines("./2022/4-camp-cleanup/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	completeCount := 0
	partialCount := 0
	for _, line := range lines {
		sections := strings.Split(line, ",")
		a := strings.Split(sections[0], "-")
		b := strings.Split(sections[1], "-")

		if checkCompleteOverlap(a, b) || checkCompleteOverlap(b, a) {
			completeCount += 1
		}

		if checkPartialOverlap(a, b) || checkPartialOverlap(b, a) {
			partialCount += 1
		}

	}

	fmt.Println("Solution 1:", completeCount)
	fmt.Println("Solution 2:", partialCount)
}

func checkCompleteOverlap(a, b []string) bool {
	l1, err := strconv.Atoi(a[0])
	h1, err := strconv.Atoi(a[1])
	l2, err := strconv.Atoi(b[0])
	h2, err := strconv.Atoi(b[1])
	if err != nil {
		log.Fatal(err)
		return false
	}

	if h1 < l2 {
		return false
	}

	if l1 >= l2 && h1 <= h2 {
		return true
	}

	return false
}

func checkPartialOverlap(a, b []string) bool {
	l1, err := strconv.Atoi(a[0])
	h1, err := strconv.Atoi(a[1])
	l2, err := strconv.Atoi(b[0])
	h2, err := strconv.Atoi(b[1])
	if err != nil {
		log.Fatal(err)
		return false
	}

	if l2 >= l1 && l2 <= h1 {
		return true
	}
	if h2 >= l1 && h2 <= h1 {
		return true
	}

	return false
}
