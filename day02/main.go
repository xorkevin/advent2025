package main

import (
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

	file, err := os.ReadFile(puzzleInput)
	if err != nil {
		log.Fatalln(err)
	}

	sum := 0
	sum2 := 0
	for i := range strings.SplitSeq(strings.TrimSpace(string(file)), ",") {
		as, bs, ok := strings.Cut(i, "-")
		if !ok {
			log.Fatalln("Invalid input")
		}
		a, err := strconv.Atoi(as)
		if err != nil {
			log.Fatalln(err)
		}
		b, err := strconv.Atoi(bs)
		if err != nil {
			log.Fatalln(err)
		}
		for i := a; i <= b; i++ {
			s := strconv.Itoa(i)
			if isDouble(s) {
				sum += i
				sum2 += i
			} else if isRepeat(s) {
				sum2 += i
			}
		}
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)
}

func isRepeat(s string) bool {
	for i := range len(s) / 3 {
		mod := i + 1
		if len(s)%mod != 0 {
			continue
		}
		if strings.Repeat(s[:mod], len(s)/mod) == s {
			return true
		}
	}
	return false
}

func isDouble(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	return s[:len(s)/2] == s[len(s)/2:]
}
