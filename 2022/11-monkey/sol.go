package monkey

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/smrnjeet222/aoc/2022/utils"
)

type Monkey struct {
	items     []int
	operation Operation
	test      Test
	inspected int
}

type Operation struct {
	op  string
	val string
}

type Test struct {
	val int
	t   int
	f   int
}

var monkeys []Monkey

func Solve() {
	println("\nGame 11: Let's go .....\n")

	lines, err := utils.ReadLines("./2022/11-monkey/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	prepareData(lines)

	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := &monkeys[j]

			for _, item := range monkey.items {
				var worry int
				if monkey.operation.op == "*" {
					if monkey.operation.val == "old" {
						worry = item * item
					} else {
						worry = item * utils.MustStrConv(monkey.operation.val)
					}
				}
				if monkey.operation.op == "+" {
					if monkey.operation.val == "old" {
						worry = item + item
					} else {
						worry = item + utils.MustStrConv(monkey.operation.val)
					}
				}

				worry /= 3
				if worry%monkey.test.val == 0 {
					monkeys[monkey.test.t].items = append(monkeys[monkey.test.t].items, worry)
					monkey.inspected++
				} else {
					monkeys[monkey.test.f].items = append(monkeys[monkey.test.f].items, worry)
					monkey.inspected++
				}
			}
			monkey.items = []int{}
		}

		// for j := 0; j < len(monkeys); j++ {
		// 	fmt.Println(monkeys[j].items)
		// }
		// fmt.Println()
	}

	inspeceted := make([]int, len(monkeys))
	for i := 0; i < len(monkeys); i++ {
		inspeceted[i] = -1 * monkeys[i].inspected
	}

	sort.Ints(inspeceted)

	fmt.Printf("Solution 1: %v\n", inspeceted[0]*inspeceted[1])
}

func prepareData(lines []string) {
	for i := 0; i < len(lines); i++ {
		itemSplit := strings.Split(lines[i], ":")

		if strings.HasSuffix(itemSplit[0], "items") {

			opSplit := strings.Split(lines[i+1], " ")
			testSplit := strings.Split(lines[i+2], " ")
			testTrueSplit := strings.Split(lines[i+3], " ")
			testFalseSplit := strings.Split(lines[i+4], " ")

			itms := utils.ConvertStrToArr(itemSplit[1])

			monkeys = append(monkeys, Monkey{
				items: itms,
				operation: Operation{
					op:  opSplit[len(opSplit)-2],
					val: opSplit[len(opSplit)-1],
				},
				test: Test{
					val: utils.MustStrConv(testSplit[len(testSplit)-1]),
					t:   utils.MustStrConv(testTrueSplit[len(testTrueSplit)-1]),
					f:   utils.MustStrConv(testFalseSplit[len(testFalseSplit)-1]),
				},
				inspected: 0,
			})
		}

	}
}
