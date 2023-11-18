package calorieCounting

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aoc2022/utils"
)

// https://adventofcode.com/2022/day/1

func Sovle() int {
	println("\nLets go ......\n")

	lines, err := utils.ReadLines("./1-calorie-counting/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// NOTE: this logic can also be processed while reading the input
	// for simplicity abrstracted the reading logic

	lastsum := 0
	newsum := 0
	for _, line := range lines {
		if line == "" {
			if newsum > lastsum {
				lastsum = newsum
			}
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
	fmt.Println(lastsum)

	return lastsum
}

func Solve2(cacheSize int) int {
	println("\nLets go again ......\n")

	lines, err := utils.ReadLines("./1-calorie-counting/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	maxStore := NewStorage(cacheSize)
	newsum := 0
	for _, line := range lines {
		if line == "" {
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

	fmt.Println(maxStore.stack)

	result := 0
	for _, s := range maxStore.stack {
		result += s
	}

	println(result)
	return result
}

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
