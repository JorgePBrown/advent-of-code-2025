package main

import (
	"aoc2025/util"
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
	f, err := os.Open("./day2/input1")
	if err != nil {
		log.Fatal(err)
	}
	solvePart1(f)
	solvePart1(strings.NewReader(`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`))
}
func solvePart1(f io.Reader) {
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	ranges := [][2]struct {
		s string
		n int64
	}{}
	for b := range strings.SplitSeq(strings.TrimSpace(string(data)), ",") {
		n := strings.Split(b, "-")
		n0, err := strconv.ParseInt(n[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		n1, err := strconv.ParseInt(n[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, [2]struct {
			s string
			n int64
		}{
			{
				s: n[0],
				n: n0,
			},
			{
				s: n[1],
				n: n1,
			},
		})
	}

	var sum int64 = 0
	for _, r := range ranges {
		for chCount := len(r[0].s); chCount <= len(r[1].s); chCount += 1 {
			if chCount&1 == 1 {
				continue
			}

			inc, err := strconv.ParseInt("1"+strings.Repeat("0", chCount/2-1)+"1", 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			start, err := strconv.ParseInt(strings.Repeat("1"+strings.Repeat("0", chCount/2-1), 2), 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			end, err := strconv.ParseInt(strings.Repeat("9", chCount), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			if chCount == len(r[1].s) {
				end = r[1].n
			}

			var n int64
			if chCount == len(r[0].s) {
				n = r[0].n

				for ; n <= end; n += 1 {
					if (n-start)%inc == 0 {
						break
					}
				}
			} else {
				n = start
			}

			for n <= end {
				sum += n
				n += inc
			}
		}
	}
	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("./day2/input1")
	if err != nil {
		log.Fatal(err)
	}
	solvePart2(f)
	solvePart2(strings.NewReader(`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`))
}
func solvePart2(f io.Reader) {
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	ranges := [][2]struct {
		s string
		n int64
	}{}
	for b := range strings.SplitSeq(strings.TrimSpace(string(data)), ",") {
		n := strings.Split(b, "-")
		n0, err := strconv.ParseInt(n[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		n1, err := strconv.ParseInt(n[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, [2]struct {
			s string
			n int64
		}{
			{
				s: n[0],
				n: n0,
			},
			{
				s: n[1],
				n: n1,
			},
		})
	}

	hs := &util.HashSet[int64]{}
	var sum int64 = 0
	for _, r := range ranges {
		for chCount := len(r[0].s); chCount <= len(r[1].s); chCount += 1 {
			if chCount == 1 {
				continue
			}

			div := divisibleBy(chCount)
			for _, d := range div {
				inc, err := strconv.ParseInt(strings.Repeat(strings.Repeat("0", chCount/d-1)+"1", d), 10, 64)
				if err != nil {
					log.Fatal(err)
				}

				start, err := strconv.ParseInt(strings.Repeat("1"+strings.Repeat("0", chCount/d-1), d), 10, 64)
				if err != nil {
					log.Fatal(err)
				}

				end, err := strconv.ParseInt(strings.Repeat("9", chCount), 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				if chCount == len(r[1].s) {
					end = r[1].n
				}

				var n int64
				if chCount == len(r[0].s) {
					n = r[0].n

					for ; n <= end; n += 1 {
						if (n-start)%inc == 0 {
							break
						}
					}
				} else {
					n = start
				}

				for n <= end {
					if !hs.Has(n) {
						sum += n
						hs.Add(n)
					}
					n += inc
				}
			}
		}
	}
	fmt.Println(sum)
}

func divisibleBy(n int) []int {
	div := []int{}
	for i := range n {
		if i > 1 && n%i == 0 {
			div = append(div, i)
		}
	}
	return append(div, n)
}
