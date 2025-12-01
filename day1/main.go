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
	f, err := os.Open("./day1/input1")
	if err != nil {
		log.Fatal(err)
	}
	solvePart1(f)

	solvePart1(strings.NewReader(`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`))
}
func solvePart1(f io.Reader) {
	scanner := bufio.NewScanner(f)

	var start, total int64 = 50, 100

	sum := 0
	var v int64 = start
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.ParseInt(line[1:], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		if line[0] == 'L' {
			v -= n
		} else {
			v += n
		}

		if v%total == 0 {
			sum += 1
		}
	}

	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("./day1/input1")
	if err != nil {
		log.Fatal(err)
	}
	solvePart2(f)

	solvePart2(strings.NewReader(`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`))

	solvePart2(strings.NewReader(`R1000`))
}
func solvePart2(f io.Reader) {
	scanner := bufio.NewScanner(f)

	var start, total int = 50, 100

	sum := 0
	var v int = start
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		q, rem := n/total, n%total

		sum += q

		if line[0] == 'L' {
			if v != 0 && rem >= v {
				sum += 1
			}
			v -= rem
		} else {
			if v != 0 && v+rem >= total {
				sum += 1
			}
			v += rem
		}

		if v < 0 {
			v += total
		}
		v %= total
	}

	fmt.Println(sum)
}
