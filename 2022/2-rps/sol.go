package rps

import (
	"fmt"
	"log"

	"github.com/aoc2022/utils"
)

type ScoreRules struct {
	R    int
	P    int
	S    int
	Lose int
	Win  int
	Draw int
}

func StartGame() *ScoreRules {
	return &ScoreRules{
		R:    1,
		P:    2,
		S:    3,
		Lose: 0,
		Win:  6,
		Draw: 3,
	}
}

func (s *ScoreRules) Play(a, b byte) int {
	// fmt.Printf("a: %c, b: %c \n", a, b)

	switch b {
	// X -> rock
	case 'X':
		var resultScore int
		if a == 'A' {
			resultScore = s.Draw
		} else if a == 'B' {
			resultScore = s.Lose
		} else {
			resultScore = s.Win
		}
		return resultScore + s.R
	// Y -> paper
	case 'Y':
		var resultScore int
		if a == 'A' {
			resultScore = s.Win
		} else if a == 'B' {
			resultScore = s.Draw
		} else {
			resultScore = s.Lose
		}
		return resultScore + s.P
	// Z -> scissors
	case 'Z':
		var resultScore int
		if a == 'A' {
			resultScore = s.Lose
		} else if a == 'B' {
			resultScore = s.Win
		} else {
			resultScore = s.Draw
		}
		return resultScore + s.S
	default:
		log.Fatalf("invalid input %c", b)
		return 0
	}
}

func (s *ScoreRules) Play2(a, b byte) int {
	switch b {
	// X means loose
	case 'X':
		var resultScore int
		if a == 'A' {
			// to losr with rock choose scissors
			resultScore = s.S
		} else if a == 'B' {
			// to lose with paper choose rock
			resultScore = s.R
		} else {
			// to lose with scissors choose paper
			resultScore = s.P
		}
		return resultScore + s.Lose
	// Y means draw
	case 'Y':
		var resultScore int
		if a == 'A' {
			// to draw with rock choose rock
			resultScore = s.R
		} else if a == 'B' {
			// to draw with paper choose paper
			resultScore = s.P
		} else {
			// to draw with scissors choose scissors
			resultScore = s.S
		}
		return resultScore + s.Draw
	// Z means win
	case 'Z':
		var resultScore int
		if a == 'A' {
			// to win with rock choose paper
			resultScore = s.P
		} else if a == 'B' {
			// to win with paper choose scissors
			resultScore = s.S
		} else {
			// to win with scissors choose rock
			resultScore = s.R
		}
		return resultScore + s.Win
	default:
		log.Fatalf("invalid input %c", b)
		return 0
	}
}

func Solve() {
	println("\nLets go ......\n")

	lines, err := utils.ReadLines("./2-rps/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	game := StartGame()

	game1Score, game2Score := 0, 0

	for _, line := range lines {
		game1Score += game.Play(line[0], line[2])
		game2Score += game.Play2(line[0], line[2])
	}

	fmt.Println("Game 1:", game1Score)
	fmt.Println("Game 2:", game2Score)
}
