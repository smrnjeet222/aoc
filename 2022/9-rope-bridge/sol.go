package ropebridge

import (
	"fmt"
	"log"
	"strconv"

	"github.com/smrnjeet222/aoc/2022/utils"
)

type Pos struct {
	x int
	y int
}

func Solve() {
	println("\nGame 9: Let's go .....\n")

	lines, err := utils.ReadLines("./2022/9-rope-bridge/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	positionSeen := make(map[Pos]struct{})
	T := Pos{0, 0}
	H := Pos{0, 0}

	for _, line := range lines {
		direction := line[0:1]
		steps, err := strconv.Atoi(line[2:])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < steps; i++ {

			if direction == "U" {
				H.y++
			} else if direction == "D" {
				H.y--
			} else if direction == "L" {
				H.x--
			} else if direction == "R" {
				H.x++
			}

			if H.x-T.x > 1 {
				if H.y != T.y {
					T.y = H.y
				}
				T.x++
			} else if H.x-T.x < -1 {
				if H.y != T.y {
					T.y = H.y
				}
				T.x--
			}

			if H.y-T.y > 1 {
				if H.x != T.x {
					T.x = H.x
				}
				T.y++
			} else if H.y-T.y < -1 {
				if H.x != T.x {
					T.x = H.x
				}
				T.y--
			}

			positionSeen[T] = struct{}{}
		}
	}

	fmt.Printf("Solution 1: %v\n", len(positionSeen))
}
