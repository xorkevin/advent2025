package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	puzzleInput = "input.txt"
)

func main() {
	log.SetFlags(log.Lshortfile)

	file, err := os.Open(puzzleInput)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	var nums [][]int
	var ops []byte

	var grid [][]byte
	var opBytes []byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if line[0] == '*' {
			opBytes = []byte(line)
			ops = make([]byte, 0, len(fields))
			for _, s := range fields {
				ops = append(ops, s[0])
			}
			continue
		}
		grid = append(grid, []byte(line))
		row := make([]int, 0, len(fields))
		for _, s := range fields {
			num, err := strconv.Atoi(strings.TrimSpace(s))
			if err != nil {
				log.Fatalln(err)
			}
			row = append(row, num)
		}
		nums = append(nums, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	total := 0
	for n, i := range ops {
		if i == '*' {
			prod := 1
			for _, i := range nums {
				prod *= i[n]
			}
			total += prod
		} else {
			sum := 0
			for _, i := range nums {
				sum += i[n]
			}
			total += sum
		}
	}

	fmt.Println("Part 1:", total)

	splits := make([]Tuple2, 0, len(ops))
	lastSplit := 0
	for n, i := range opBytes {
		if i != ' ' {
			t := Tuple2{
				x: lastSplit,
				y: n - 1,
			}
			if t.x >= t.y {
				continue
			}
			lastSplit = n
			splits = append(splits, t)
		}
	}
	splits = append(splits, Tuple2{
		x: lastSplit,
		y: len(opBytes),
	})

	total = 0
	for n, i := range splits {
		if ops[n] == '*' {
			prod := 1
			for j := i.x; j < i.y; j++ {
				prod *= calcColNum(grid, j)
			}
			total += prod
		} else {
			sum := 0
			for j := i.x; j < i.y; j++ {
				sum += calcColNum(grid, j)
			}
			total += sum
		}
	}
	fmt.Println("Part 2:", total)
}

type (
	Tuple2 struct {
		x, y int
	}
)

func calcColNum(grid [][]byte, c int) int {
	v := 0
	for _, i := range grid {
		digit := int(i[c] - '0')
		if digit <= 9 {
			v *= 10
			v += digit
		}
	}
	return v
}
