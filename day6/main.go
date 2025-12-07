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
	f, err := os.Open("./day6/input")
	if err != nil {
		log.Fatal(err)
	}
	solvePart1(f)
	solvePart1(strings.NewReader(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + `))
}
func solvePart1(f io.Reader) {
	operands := [][]int{}
	operators := []byte{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineOperands := []int{}
		n := 0
		for _, b := range scanner.Bytes() {
			switch b {
			case ' ', '\n':
				if n != 0 {
					lineOperands = append(lineOperands, n)
					n = 0
				}
			case '+', '*':
				operators = append(operators, b)
			default:
				n *= 10
				n += int(b - '0')
			}
		}
		if n != 0 {
			lineOperands = append(lineOperands, n)
		}

		if len(lineOperands) > 0 {
			operands = append(operands, lineOperands)
		}
	}

	sum := 0
	for i := range len(operands[0]) {
		operator := operators[i]
		var v int
		switch operator {
		case '*':
			v = 1
		case '+':
			v = 0
		}

		for j := range operands {
			switch operator {
			case '*':
				v *= operands[j][i]
			case '+':
				v += operands[j][i]
			}
		}
		sum += v
	}
	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("./day6/input")
	if err != nil {
		log.Fatal(err)
	}
	solvePart2(f)
	solvePart2(strings.NewReader(
		`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `))
}
func solvePart2(f io.Reader) {
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	problems := []*Problem{}
	for _, ch := range lines[len(lines)-1] {
		if ch != ' ' {
			problems = append(problems, newProblem(byte(ch)))
		}
	}

	columnsChecked := make([]bool, len(lines[0]))
	for i := 0; i < len(lines)-1; i += 1 {
		problemIdx := len(problems) - 1
		inNumber := false
		for j := len(lines[i]) - 1; j >= 0; j -= 1 {
			if lines[i][j] == ' ' {
				if inNumber {
					problemIdx -= 1
					inNumber = false
				}
				continue
			}
			inNumber = true
			if columnsChecked[j] {
				continue
			}
			n := 0
			for k := i; k < len(lines)-1; k += 1 {
				if lines[k][j] != ' ' {
					n *= 10
					n += int(lines[k][j] - '0')
				}
			}
			problems[problemIdx].Update(n)

			columnsChecked[j] = true
		}
	}

	sum := 0
	for _, p := range problems {
		sum += p.v
	}
	fmt.Println(sum)
}

type Problem struct {
	v        int
	operator byte
}

func newProblem(operator byte) *Problem {
	switch operator {
	case '+':
		return &Problem{
			v:        0,
			operator: operator,
		}
	case '*':
		return &Problem{
			v:        1,
			operator: operator,
		}
	}

	log.Fatal("unknown operator", operator)
	return nil
}

func (p *Problem) Update(v int) {
	switch p.operator {
	case '+':
		p.v += v
	case '*':
		p.v *= v
	}
}
