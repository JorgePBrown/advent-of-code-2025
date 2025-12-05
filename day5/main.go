package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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
	f, err := os.Open("./day5/input")
	if err != nil {
		log.Fatal(err)
	}
	solvePart1(f)
	solvePart1(strings.NewReader(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`))
}
func solvePart1(f io.Reader) {
	scanner := bufio.NewScanner(f)
	ranges := [][2]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		dashIdx := strings.IndexByte(line, '-')
		if dashIdx == -1 {
			log.Fatal("missing dash")
		}

		n1, err := strconv.Atoi(line[:dashIdx])
		if err != nil {
			log.Fatal(err)
		}
		n2, err := strconv.Atoi(line[dashIdx+1:])
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, [2]int{n1, n2})
	}

	sum := 0
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		for _, r := range ranges {
			if n >= r[0] && n <= r[1] {
				sum += 1
				break
			}
		}
	}

	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("./day5/input")
	if err != nil {
		log.Fatal(err)
	}
	solvePart2(f)
	solvePart2(strings.NewReader(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`))
}
func solvePart2(f io.Reader) {
	scanner := bufio.NewScanner(f)
	ranges := []*Range{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		dashIdx := strings.IndexByte(line, '-')
		if dashIdx == -1 {
			log.Fatal("missing dash")
		}

		n1, err := strconv.Atoi(line[:dashIdx])
		if err != nil {
			log.Fatal(err)
		}
		n2, err := strconv.Atoi(line[dashIdx+1:])
		if err != nil {
			log.Fatal(err)
		}

		ranges = append(ranges, &Range{n1, n2})

		for i := range ranges {
			if ranges[i] != nil {
				for j := range ranges {
					if j != i && ranges[j] != nil {
						ranges[i], ranges[j] = merge(ranges[i], ranges[j])
						if ranges[i] == nil || ranges[j] == nil {
							break
						}
					}
				}
			}
		}
	}

	sum := 0
	for _, r := range ranges {
		if r != nil {
			sum += r[1] - r[0] + 1
		}
	}

	fmt.Println(sum)
}

type Range [2]int

func (r Range) String() string {
	return strconv.Itoa(r[0]) + "-" + strconv.Itoa(r[1])
}

func merge(r1, r2 *Range) (*Range, *Range) {
	var gap int
	if r1[0] > r2[0] {
		gap = r1[0] - r2[1]
	} else {
		gap = r2[0] - r1[1]
	}

	if gap > 0 {
		return r1, r2
	} else if gap == 0 {
		if r1[0] > r2[0] {
			r2[1] = r1[1]
			return r2, nil
		} else {
			r1[1] = r2[1]
			return r1, nil
		}
	} else {
		if r1[0] > r2[0] {
			if r2[1] < r1[1] {
				r2[1] = r1[1]
			}
			return r2, nil
		} else {
			if r1[1] < r2[1] {
				r1[1] = r2[1]
			}
			return r1, nil
		}
	}
}
