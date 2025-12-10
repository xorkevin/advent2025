package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		t := parseTarget(fields[0])
		buttons := make([]uint32, 0, len(fields)-2)
		for _, i := range fields[1 : len(fields)-1] {
			b, err := parseButton(i)
			if err != nil {
				log.Fatalln(err)
			}
			buttons = append(buttons, b)
		}
		v := findToggleSet(t, 0, buttons, map[uint64]int{})
		if v < 0 {
			log.Fatalln("not possible", scanner.Text())
		}
		sum += v
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Part 1:", sum)
}

func findToggleSet(t uint32, o int, buttons []uint32, cache map[uint64]int) int {
	if t == 0 {
		return 0
	}
	if o == len(buttons) {
		return -1
	}

	id := uint64(o)<<32 | uint64(t)
	if v, ok := cache[id]; ok {
		return v
	}

	b := buttons[o]
	c0 := findToggleSet(t, o+1, buttons, cache)
	c1 := findToggleSet(t^b, o+1, buttons, cache)

	var v int
	if c0 < 0 {
		if c1 < 0 {
			v = -1
		} else {
			v = c1 + 1
		}
	} else {
		if c1 < 0 {
			v = c0
		} else {
			v = min(c0, c1+1)
		}
	}
	cache[id] = v
	return v
}

func parseTarget(s string) uint32 {
	var x uint32 = 0
	for _, i := range slices.Backward([]byte(s[1 : len(s)-1])) {
		x <<= 1
		if i == '#' {
			x |= 1
		}
	}
	return x
}

func parseButton(s string) (uint32, error) {
	var x uint32 = 0
	for a := range strings.SplitSeq(s[1:len(s)-1], ",") {
		i, err := strconv.Atoi(a)
		if err != nil {
			return 0, err
		}
		x |= (1 << i)
	}
	return x, nil
}
