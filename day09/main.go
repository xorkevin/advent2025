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

	var points []Vec2

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.SplitN(scanner.Text(), ",", 2)
		x, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalln(err)
		}
		y, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalln(err)
		}
		points = append(points, Vec2{
			x: x,
			y: y,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	area := 0
	area2 := 0
	for i := range len(points) - 1 {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]
			v := (abs(a.x-b.x) + 1) * (abs(a.y-b.y) + 1)
			area = max(area, v)
			if checkIntersection(a, b, points) {
				area2 = max(area2, v)
			}
		}
	}
	fmt.Println("Part 1:", area)
	fmt.Println("Part 2:", area2)
}

func checkIntersection(a, b Vec2, points []Vec2) bool {
	if a.x > b.x {
		a.x, b.x = b.x, a.x
	}
	if a.y > b.y {
		a.y, b.y = b.y, a.y
	}
	for i := range len(points) - 1 {
		c := points[i]
		d := points[i+1]
		if c.x > d.x {
			c.x, d.x = d.x, c.x
		}
		if c.y > d.y {
			c.y, d.y = d.y, c.y
		}
		if c.x == d.x {
			if c.x > a.x && c.x < b.x && c.y < b.y && d.y > a.y {
				return false
			}
		} else {
			if c.y > a.y && c.y < b.y && c.x < b.x && d.x > a.x {
				return false
			}
		}
	}
	return true
}

type (
	Vec2 struct {
		x, y int
	}
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
