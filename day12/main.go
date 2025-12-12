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

	areas := []int{6, 7, 5, 7, 7, 7}

	upperBoundCount := 0
	lowerBoundCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 3 {
			continue
		}
		lhs, rhs, ok := strings.Cut(line, ": ")
		if !ok {
			log.Fatalln("Invalid line")
		}
		ws, hs, ok := strings.Cut(lhs, "x")
		if !ok {
			log.Fatalln("Invalid line")
		}
		width, err := strconv.Atoi(ws)
		if err != nil {
			log.Fatalln(err)
		}
		height, err := strconv.Atoi(hs)
		if err != nil {
			log.Fatalln(err)
		}
		boundingArea := width * height
		lowerBoundArea := 0
		upperBoundArea := 0
		for n, i := range strings.Fields(rhs) {
			num, err := strconv.Atoi(i)
			if err != nil {
				log.Fatalln(err)
			}
			lowerBoundArea += num * areas[n]
			upperBoundArea += num * 9
		}
		if lowerBoundArea <= boundingArea {
			upperBoundCount++
		}
		if upperBoundArea <= boundingArea {
			lowerBoundCount++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Part 1:", lowerBoundCount, upperBoundCount)
}
