package calorieCounting

import (
	"fmt"
	"log"
	"strconv"

	"github.com/smrnjeet222/aoc/2022/utils"
)

// https://adventofcode.com/2022/day/1

type Storage struct {
	cap   int
	stack []int
}

func NewStorage(cap int) *Storage {
	return &Storage{
		cap:   cap,
		stack: make([]int, cap),
	}
}

// LRU type cache stack
func (s *Storage) Push(num int) {
	for i, val := range s.stack {
		if num > val {
			s.stack[i] = num
			return
		}
	}
}

func Solve() {
	println("\nGame 1 : Lets go ......\n")

	lines, err := utils.ReadLines("./2022/1-calorie-counting/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// NOTE: this logic can also be processed while reading the input
	// for simplicity abrstracted the reading logic

	lastsum := 0
	newsum := 0

	// For solution 2
	maxStore := NewStorage(3)

	for _, line := range lines {
		if line == "" {
			if newsum > lastsum {
				lastsum = newsum
			}
			// solution 2
			maxStore.Push(newsum)

			newsum = 0
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("Error converting string to int", err, line)
			}
			newsum += i
		}
	}

	// NOTE: to make sure last sum is greatest add empty line at the end of the input file

	fmt.Println("Solution 1:", lastsum)

	result := 0
	for _, s := range maxStore.stack {
		result += s
	}

	fmt.Println("Solution 2:", result)
}
