package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func MustStrConv(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func ConvertStrToArr(str string) []int {
	a := strings.Split((strings.Trim(str, " ")), ", ")
	b := make([]int, len(a))

	for i, v := range a {
		b[i], _ = strconv.Atoi(v)
	}

	return b
}
