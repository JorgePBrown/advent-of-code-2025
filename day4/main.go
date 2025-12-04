package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	part := os.Args[1]
	switch part {
	case "1":
		part1()
	case "2":
		part2()
	}
}

func part1() {
	f, err := os.Open("./day4/input")
	if err != nil {
		log.Fatal(err)
	}
	solvePart1(f)
	solvePart1(strings.NewReader(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`))
}

const (
	PAPER = '@'
	EMPTY = '.'
)

func solvePart1(f io.Reader) {
	mat := [][]byte{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mat = append(mat, bytes.Clone(scanner.Bytes()))
	}

	isPaper := func(x, y int) int {
		if x < 0 || x >= len(mat) || y < 0 || y >= len(mat[x]) || mat[x][y] != PAPER {
			return 0
		}
		return 1
	}

	sum := 0

	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == EMPTY {
				continue
			}
			paperCount := isPaper(i-1, j-1)
			paperCount += isPaper(i-1, j)
			paperCount += isPaper(i-1, j+1)

			paperCount += isPaper(i, j-1)
			paperCount += isPaper(i, j+1)

			paperCount += isPaper(i+1, j-1)
			paperCount += isPaper(i+1, j)
			paperCount += isPaper(i+1, j+1)

			if paperCount < 4 {
				sum += 1
			}
		}
	}

	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("./day4/input")
	if err != nil {
		log.Fatal(err)
	}
	solvePart2(f)
	solvePart2(strings.NewReader(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`))
}
func solvePart2(f io.Reader) {
	mat := [][]byte{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mat = append(mat, bytes.Clone(scanner.Bytes()))
	}

	isPaper := func(x, y int) int {
		if x < 0 || x >= len(mat) || y < 0 || y >= len(mat[x]) || mat[x][y] != PAPER {
			return 0
		}
		return 1
	}

	papersToRemove := [][2]int{}

	canRemove := func(i, j int) bool {
		if i < 0 || i >= len(mat) || j < 0 || j >= len(mat[i]) || mat[i][j] != PAPER {
			return false
		}
		paperCount := isPaper(i-1, j-1)
		paperCount += isPaper(i-1, j)
		paperCount += isPaper(i-1, j+1)

		paperCount += isPaper(i, j-1)
		paperCount += isPaper(i, j+1)

		paperCount += isPaper(i+1, j-1)
		paperCount += isPaper(i+1, j)
		paperCount += isPaper(i+1, j+1)

		return paperCount < 4
	}

	for i := range mat {
		for j := range mat[i] {
			if canRemove(i, j) {
				papersToRemove = append(papersToRemove, [2]int{i, j})
			}
		}
	}

	sum := 0
	for len(papersToRemove) > 0 {
		sum += len(papersToRemove)

		for _, p := range papersToRemove {
			mat[p[0]][p[1]] = EMPTY
		}

		newPapersToRemove := [][2]int{}
		appendToNewPapers := func(x, y int) {
			if canRemove(x, y) && !slices.ContainsFunc(newPapersToRemove, func(pair [2]int) bool { return pair[0] == x && pair[1] == y }) {
				newPapersToRemove = append(newPapersToRemove, [2]int{x, y})
			}
		}
		for _, p := range papersToRemove {
			appendToNewPapers(p[0]-1, p[1]-1)
			appendToNewPapers(p[0]-1, p[1])
			appendToNewPapers(p[0]-1, p[1]+1)

			appendToNewPapers(p[0], p[1]-1)
			appendToNewPapers(p[0], p[1]+1)

			appendToNewPapers(p[0]+1, p[1]-1)
			appendToNewPapers(p[0]+1, p[1])
			appendToNewPapers(p[0]+1, p[1]+1)
		}

		papersToRemove = newPapersToRemove
	}

	fmt.Println(sum)
}
