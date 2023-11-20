package nospace

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/smrnjeet222/aoc/2022/utils"
)

var (
	MAX_DISK_SPACE    = 70000000
	FREE_SPACE_NEEDED = 30000000
)

type Node struct {
	parent *Node
	name   string
	dirs   []*Node
	files  []int // only storing sizes
	size   int
}

func CreateRoot() *Node {
	return &Node{
		parent: nil,
		name:   "root",
		dirs:   nil,
		files:  nil,
		size:   0,
	}
}

func (n *Node) AddDir(name string) {
	n.dirs = append(n.dirs, &Node{
		parent: n,
		name:   name,
		dirs:   nil,
		files:  nil,
		size:   0,
	})
}

func (n *Node) AddFile(file int) {
	n.files = append(n.files, file)
	n.size += file
	for n.parent != nil {
		n.parent.size += file
		n = n.parent
	}
}

func Solve() {
	println("\nGame 7: Let's go .....\n")

	lines, err := utils.ReadLines("./2022/7-no-space/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	root := CreateRoot()
	var cwd *Node

	for _, line := range lines {
		input := strings.Split(line, " ")

		// process $ -> command
		if input[0] == "$" {
			// process cd into
			if input[1] == "cd" {

				if input[2] == "/" {
					cwd = root
					continue
				}
				if input[2] == ".." {
					cwd = cwd.parent
					continue
				}
				// process cd into dir
				if input[2] != "/" && input[2] != ".." {
					for _, d := range cwd.dirs {
						if d.name == input[2] {
							cwd = d
							break
						}
					}
					continue
				}
			}

			// nothin much to do on ls
			if input[1] == "ls" {
				continue
			}

			// process Folder
		} else if input[0] == "dir" {
			cwd.AddDir(input[1])
			continue
		} else {
			// process File
			size, err := strconv.Atoi(input[0])
			if err != nil {
				log.Fatal(err)
			}
			cwd.AddFile(size)
			continue
		}
	}

	fmt.Printf("Solution 1: %v \n", countTotalSize(root))

	remainingSpace := MAX_DISK_SPACE - root.size

	spaceToDelete := FREE_SPACE_NEEDED - remainingSpace

	// fmt.Printf("Space to delete: %v \n", spaceToDelete)
	fmt.Printf("Solution 2: %v \n", root.dirToClean(spaceToDelete).size)
}

var depth int

func countTotalSize(n *Node) int {
	size := 0

	// fmt.Printf("%s %s ( %v ) \n", strings.Repeat("-", depth), n.name, n.size)
	if n.size <= 100000 {
		size += n.size
	}

	if n.dirs == nil {
		if n.size <= 100000 {
			return n.size
		}
		return 0
	}
	for _, d := range n.dirs {
		depth++
		size += countTotalSize(d)
		depth--
	}

	return size
}

var smallestNode *Node

func (n *Node) dirToClean(spaceToDelete int) *Node {
	if n.dirs == nil {
		if n.size >= spaceToDelete {
			if smallestNode == nil || n.size < smallestNode.size {
				smallestNode = n
			}
		}
		return nil
	}

	for _, d := range n.dirs {

		if n.size >= spaceToDelete {
			if smallestNode == nil || n.size < smallestNode.size {
				smallestNode = n
			}
		}

		d.dirToClean(spaceToDelete)
	}
	return smallestNode
}
