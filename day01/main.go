package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	start := 50
	zeroCount := 0
	zeroCount2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var isLeft bool
		switch line[0] {
		case 'L':
			isLeft = true
		case 'R':
			isLeft = false
		default:
			log.Fatalln("Invalid direction")
		}
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalln(err)
		}

		zeroCount2 += num / 100
		num = num % 100
		if num > 0 {
			if isLeft {
				if start == 0 {
					start = 100 - num
				} else {
					start -= num
				}
			} else {
				start += num
			}
			if start == 0 || start == 100 {
				start = 0
				zeroCount += 1
				zeroCount2 += 1
			} else if start < 0 {
				start += 100
				zeroCount2 += 1
			} else if start > 100 {
				start -= 100
				zeroCount2 += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Part 1:", zeroCount)
	fmt.Println("Part 2:", zeroCount2)
}
