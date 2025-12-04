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

	sum := 0
	sum2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := make([]byte, 0, len(scanner.Bytes()))
		for _, i := range scanner.Bytes() {
			line = append(line, i-'0')
		}
		sum += findMaxNum(line)
		sum2 += findMaxNumRec(line, 0, 11, map[int]int{})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)
}

func findMaxNum(line []byte) int {
	num := 0
	for i := range len(line) - 1 {
		for j := i + 1; j < len(line); j++ {
			num = max(num, int(line[i])*10+int(line[j]))
		}
	}
	return num
}

func findMaxNumRec(line []byte, start, size int, cache map[int]int) int {
	id := start*20 + size
	if v, ok := cache[id]; ok {
		return v
	}
	// start is start pos, and size is 1 less than digits
	num := 0
	mul := 1
	for range size {
		mul *= 10
	}
	for i := start; i < len(line)-size; i++ {
		x := int(line[i])
		if size == 0 {
			num = max(num, x)
		} else {
			num = max(num, x*mul+findMaxNumRec(line, i+1, size-1, cache))
		}
	}
	cache[id] = num
	return num
}
