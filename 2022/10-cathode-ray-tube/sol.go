package cathoderaytube

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/smrnjeet222/aoc/2022/utils"
)

var result1 int

type CPU struct {
	laps     []int
	cycle    int
	register int
	crt      [6][40]string
}

func NewCPU() *CPU {
	return &CPU{
		laps:     []int{20, 60, 100, 140, 180, 220},
		cycle:    0,
		register: 1,
		crt:      [6][40]string{},
	}
}

func (c *CPU) IncCycle() {
	c.cycle++
	if slices.Contains(c.laps, c.cycle) {
		result1 += c.cycle * c.register
	}
}

func (c *CPU) IncCycleWithCRT() {
	row := c.cycle / 40
	col := c.cycle % 40

	if c.register-1 <= col && col <= c.register+1 {
		c.crt[row][col] = "#"
	} else {
		c.crt[row][col] = "."
	}

	c.cycle++
}

func (c *CPU) PrintCRT() {
	for _, row := range c.crt {
		fmt.Println(row)
	}
}

func Solve() {
	println("\nGame 10: Lets go .....\n")

	lines, err := utils.ReadLines("./2022/10-cathode-ray-tube/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	cpu := NewCPU()
	cpu2 := NewCPU()

	for _, line := range lines {

		if line == "noop" {
			cpu.IncCycle()
			cpu2.IncCycleWithCRT()
			continue
		}

		cmd := strings.Split(line, " ")

		var buffer string

		if cmd[0] == "addx" {
			cpu.IncCycle()
			cpu2.IncCycleWithCRT()

			buffer = cmd[1]
		}

		if buffer != "" {
			cpu.IncCycle()
			cpu2.IncCycleWithCRT()

			rInt, _ := strconv.Atoi(buffer)
			cpu.register += rInt
			cpu2.register += rInt

			buffer = ""

		}

	}

	fmt.Println("Solution 1:", result1)

	fmt.Println("Solution 2:")
	cpu2.PrintCRT()
}
