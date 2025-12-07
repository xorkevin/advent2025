package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
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

	var grid [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	start := bytes.IndexByte(grid[0], 'S')
	if start < 0 {
		log.Fatalln("No start found")
	}

	splitCount := 0

	row := make([]int, len(grid[0]))
	next := make([]int, len(grid[0]))
	row[start] = 1

	for _, i := range grid {
		for c, j := range i {
			if k := row[c]; k > 0 {
				if j == '^' {
					if c > 0 {
						next[c-1] += k
					}
					if c < len(grid[0])-1 {
						next[c+1] += k
					}
					splitCount++
				} else {
					next[c] += k
				}
			}
		}
		row, next = next, row
		clear(next)
	}

	fmt.Println("Part 1:", splitCount)

	sum := 0
	for _, i := range row {
		sum += i
	}
	fmt.Println("Part 2:", sum)
}
