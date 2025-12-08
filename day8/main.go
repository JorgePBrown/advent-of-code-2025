package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
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
	f, err := os.Open("./day8/input")
	if err != nil {
		log.Fatal(err)
	}
	solvePart1(f, 1000)
	solvePart1(strings.NewReader(`162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`), 10)
}
func solvePart1(f io.Reader, connections int) {
	scanner := bufio.NewScanner(f)
	jbs := []JunctionBox{}
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		jb := JunctionBox{}
		var err error
		jb.x, err = strconv.Atoi(s[0])
		if err != nil {
			log.Fatal(err)
		}
		jb.y, err = strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		jb.z, err = strconv.Atoi(s[2])
		if err != nil {
			log.Fatal(err)
		}
		jbs = append(jbs, jb)
	}

	circuits := make([]int, len(jbs))
	for i := range circuits {
		circuits[i] = i
	}

	pAmount := 0
	type DistanceMarker struct {
		distance float64
		i, j     int
	}
	closestPairs := make([]DistanceMarker, connections)

	for i := range jbs {
		for j := i + 1; j < len(jbs); j += 1 {
			d := jbs[i].Distance(jbs[j])
			if pAmount < connections {
				closestPairs[pAmount] = DistanceMarker{d, i, j}
				pAmount += 1
				if pAmount == connections {
					slices.SortFunc(closestPairs, func(o1, o2 DistanceMarker) int {
						if o1.distance == 0 {
							return 1
						} else if o2.distance == 0 {
							return -1
						}
						return int(math.Ceil(o1.distance - o2.distance))
					})
				}
			} else if d < closestPairs[connections-1].distance {
				closestPairs[connections-1] = DistanceMarker{d, i, j}
				slices.SortFunc(closestPairs, func(o1, o2 DistanceMarker) int {
					if o1.distance == 0 {
						return 1
					} else if o2.distance == 0 {
						return -1
					}
					return int(math.Ceil(o1.distance - o2.distance))
				})
			}
		}
	}

	for i := 0; i < pAmount; i += 1 {
		pair := closestPairs[i]
		dsuUnion(circuits, pair.i, pair.j)
	}
	cCount := make([]int, len(circuits))

	for i := range circuits {
		cCount[dsuFind(circuits, i)] += 1
	}

	maxmax, max2, max3 := 0, 0, 0
	for _, count := range cCount {
		if count > maxmax {
			max3, max2, maxmax = max2, maxmax, count
		} else if count > max2 {
			max3, max2 = max2, count
		} else if count > max3 {
			max3 = count
		}
	}
	fmt.Println(maxmax * max2 * max3)
}

func dsuFind(dsu []int, x int) int {
	if dsu[x] != x {
		dsu[x] = dsuFind(dsu, dsu[x])
	}
	return dsu[x]
}
func dsuUnion(dsu []int, x int, y int) {
	rootX, rootY := dsuFind(dsu, x), dsuFind(dsu, y)
	if rootX == rootY {
		return
	}
	dsu[rootX] = rootY
}

type JunctionBox struct {
	x, y, z int
}

func (jb JunctionBox) Distance(jb2 JunctionBox) float64 {
	x, y, z := jb2.x-jb.x, jb2.y-jb.y, jb2.z-jb.z
	d2 := math.Sqrt(float64(x*x + y*y))

	d3 := math.Sqrt(d2*d2 + float64(z*z))

	return d3
}

func part2() {
	f, err := os.Open("./day8/input")
	if err != nil {
		log.Fatal(err)
	}
	solvePart2(f)
}
func solvePart2(f io.Reader) {
	scanner := bufio.NewScanner(f)
	jbs := []JunctionBox{}
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		jb := JunctionBox{}
		var err error
		jb.x, err = strconv.Atoi(s[0])
		if err != nil {
			log.Fatal(err)
		}
		jb.y, err = strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		jb.z, err = strconv.Atoi(s[2])
		if err != nil {
			log.Fatal(err)
		}
		jbs = append(jbs, jb)
	}

	circuits := make([]int, len(jbs))
	for i := range circuits {
		circuits[i] = i
	}

	pAmount := 0
	type DistanceMarker struct {
		distance float64
		i, j     int
	}
	connections := len(jbs) * (len(jbs) - 1) / 2
	closestPairs := make([]DistanceMarker, connections)

	for i := range jbs {
		for j := i + 1; j < len(jbs); j += 1 {
			d := jbs[i].Distance(jbs[j])
			if pAmount < connections {
				closestPairs[pAmount] = DistanceMarker{d, i, j}
				pAmount += 1
				if pAmount == connections {
					slices.SortFunc(closestPairs, func(o1, o2 DistanceMarker) int {
						if o1.distance == 0 {
							return 1
						} else if o2.distance == 0 {
							return -1
						}
						return int(math.Ceil(o1.distance - o2.distance))
					})
				}
			} else if d < closestPairs[connections-1].distance {
				closestPairs[connections-1] = DistanceMarker{d, i, j}
				slices.SortFunc(closestPairs, func(o1, o2 DistanceMarker) int {
					if o1.distance == 0 {
						return 1
					} else if o2.distance == 0 {
						return -1
					}
					return int(math.Ceil(o1.distance - o2.distance))
				})
			}
		}
	}

	for i := 0; i < pAmount; i += 1 {
		pair := closestPairs[i]
		dsuUnion(circuits, pair.i, pair.j)

		oneCircuit := true
		root := -1
		for i := range circuits {
			if root == -1 {
				root = dsuFind(circuits, i)
			} else if root != dsuFind(circuits, i) {
				oneCircuit = false
				break
			}
		}
		if oneCircuit {
			fmt.Println(jbs[pair.i].x * jbs[pair.j].x)
			return
		}
	}
}
