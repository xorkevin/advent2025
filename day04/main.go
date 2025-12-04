package main

import (
	"bufio"
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

	h := len(grid)
	w := len(grid[0])

	count := 0
	for r := range h {
		for c := range w {
			if grid[r][c] == '@' && getNeighbors(grid, w, h, r, c) < 4 {
				count++
			}
		}
	}
	fmt.Println("Part 1:", count)
	count = 0
	for {
		progress := false
		for r := range h {
			for c := range w {
				if grid[r][c] == '@' && getNeighbors(grid, w, h, r, c) < 4 {
					grid[r][c] = '.'
					count++
					progress = true
				}
			}
		}
		if !progress {
			break
		}
	}
	fmt.Println("Part 2:", count)
}

type (
	Pos struct {
		x, y int
	}
)

func getNeighbors(grid [][]byte, w, h, r, c int) int {
	count := 0
	for _, i := range []Pos{
		{x: -1, y: -1},
		{x: 0, y: -1},
		{x: 1, y: -1},
		{x: 1, y: 0},
		{x: 1, y: 1},
		{x: 0, y: 1},
		{x: -1, y: 1},
		{x: -1, y: 0},
	} {
		y := r + i.y
		x := c + i.x
		if inBounds(w, h, y, x) && grid[y][x] == '@' {
			count++
		}
	}
	return count
}

func inBounds(w, h, r, c int) bool {
	return r >= 0 && c >= 0 && r < h && c < w
}
