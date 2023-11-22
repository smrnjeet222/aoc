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

	fmt.Printf("Solution 1: %v\n", Approach1(lines))
	fmt.Printf("Solution 2: %v\n", Approach2(lines))
}

func Approach1(lines []string) int {
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

	return len(positionSeen)
}

//NOTE: Real issue with before appraoch is the direct jump from made by T.x = H.x and T.y = H.y
// i.e making thr tail immediately follow the head
func Approach2(lines []string) int {
	positionSeen := make(map[Pos]struct{})
	rope := make([]Pos, 10) // 10 knots

	for _, line := range lines {
		direction := line[0:1]
		steps, err := strconv.Atoi(line[2:])
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(direction, steps)

		for i := 0; i < steps; i++ {

			if direction == "U" {
				rope[0].y++
			} else if direction == "D" {
				rope[0].y--
			} else if direction == "L" {
				rope[0].x--
			} else if direction == "R" {
				rope[0].x++
			}

			for r := 1; r < len(rope); r++ {
				H := &rope[r-1]
				T := &rope[r]

				if H.x-T.x > 1 {
					if H.y > T.y {
						T.y++
					} else if H.y < T.y {
						T.y--
					}
					T.x++
				} else if H.x-T.x < -1 {
					if H.y > T.y {
						T.y++
					} else if H.y < T.y {
						T.y--
					}
					T.x--
				}

				if H.y-T.y > 1 {
					if H.x > T.x {
						T.x++
					} else if H.x < T.x {
						T.x--
					}
					T.y++
				} else if H.y-T.y < -1 {
					if H.x > T.x {
						T.x++
					} else if H.x < T.x {
						T.x--
					}
					T.y--
				}
			}

			positionSeen[rope[len(rope)-1]] = struct{}{}
		}
		// CreatGrid(rope)
	}

	return len(positionSeen)
}

func CreatGrid(rope []Pos) {
	size := 32
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			skip := false
			for k := 0; k < len(rope); k++ {
				if size-i-(size/2) == rope[k].y && j-(size/2) == rope[k].x {
					fmt.Printf("%v ", k)
					skip = true
					break
				}
			}
			if !skip {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
