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
	f, err := os.Open("./day3/input")
	if err != nil {
		log.Fatal(err)
	}
	solvePart1(f)
	solvePart1(strings.NewReader(`987654321111111
811111111111119
234234234234278
818181911112111`))
}

func solvePart1(f io.Reader) {
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		var largest byte = 0
		var largestIdx int = -1

		for i := 0; i < len(line)-1; i += 1 {
			ch := line[i]
			if ch-'0' > largest {
				largest = ch - '0'
				largestIdx = i
			}
			if largest == 9 {
				break
			}
		}

		var secondLargest byte = 0
		for i := len(line) - 1; i > largestIdx; i -= 1 {
			ch := line[i]
			if ch-'0' > secondLargest {
				secondLargest = ch - '0'
			}
			if secondLargest == 9 {
				break
			}
		}

		sum += int(largest)*10 + int(secondLargest)
	}
	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("./day3/input")
	if err != nil {
		log.Fatal(err)
	}
	solvePart2(f)
	solvePart2(strings.NewReader(`987654321111111
811111111111119
234234234234278
818181911112111`))
}
func solvePart2(f io.Reader) {
	scanner := bufio.NewScanner(f)
	nBatteries := 12
	sum := 0
	for scanner.Scan() {
		line := scanner.Bytes()

		batteriesPicked := 0
		n := 0
		var largestIdx int = -1
		for batteriesPicked < nBatteries {
			var largest byte = 0
			for i := largestIdx + 1; i < len(line)-(nBatteries-batteriesPicked)+1; i += 1 {
				ch := line[i]
				if ch-'0' > largest {
					largest = ch - '0'
					largestIdx = i
				}
				if largest == 9 {
					break
				}
			}
			batteriesPicked += 1
			n = n*10 + int(largest)
		}
		sum += n
	}
	fmt.Println(sum)
}
