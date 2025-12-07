package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
	f, err := os.Open("./day7/input")
	if err != nil {
		log.Fatal(err)
	}
	solvePart1(f)
	solvePart1(strings.NewReader(`.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`))
}
func solvePart1(f io.Reader) {
	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		log.Fatal("incorrect input formatting")
	}

	beamIndices := make([]bool, len(scanner.Bytes()))
	for i, b := range scanner.Bytes() {
		if b == 'S' {
			beamIndices[i] = true
			break
		}
	}

	beamIndicesSwap := make([]bool, len(scanner.Bytes()))
	splits := 0
	for scanner.Scan() {
		copy(beamIndicesSwap, beamIndices)
		for i, b := range scanner.Bytes() {
			if b == '^' {
				if beamIndices[i] {
					splits += 1
					beamIndicesSwap[i] = false
					if i > 0 {
						beamIndicesSwap[i-1] = true
					}
					if i < len(beamIndices)-1 {
						beamIndicesSwap[i+1] = true
					}
				}
			}
		}
		beamIndices = beamIndicesSwap
	}
	fmt.Println(splits)
}

func part2() {
	f, err := os.Open("./day7/input")
	if err != nil {
		log.Fatal(err)
	}
	solvePart2(f)
	solvePart2(strings.NewReader(`.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`))
}
func solvePart2(f io.Reader) {
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var beamIndex int = -1
	for i, b := range lines[0] {
		if b == 'S' {
			beamIndex = i
			break
		}
	}
	if beamIndex == -1 {
		log.Fatal("no starting index")
	}

	memo := map[int]map[int]int{}
	for i := range len(lines) {
		memo[i] = map[int]int{}
	}
	timelines := exploreTimeline(lines, 1, beamIndex, memo)
	fmt.Println(timelines)
}

func exploreTimeline(lines []string, depth, beamIndex int, memo map[int]map[int]int) int {
	if depth >= len(lines) {
		return 1
	}
	if tl, ok := memo[depth][beamIndex]; ok {
		return tl
	}
	if lines[depth][beamIndex] != '^' {
		tl := exploreTimeline(lines, depth+1, beamIndex, memo)
		memo[depth][beamIndex] = tl
		return tl
	} else {
		tl := 0
		if beamIndex > 0 {
			tl += exploreTimeline(lines, depth+1, beamIndex-1, memo)
		}
		if beamIndex < len(lines[depth])-1 {
			tl += exploreTimeline(lines, depth+1, beamIndex+1, memo)
		}
		memo[depth][beamIndex] = tl
		return tl
	}
}
