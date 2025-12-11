package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

	graph := map[string][]string{}
	inDegree := map[string]int{}
	nodeSet := map[string]struct{}{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lhs, rhs, ok := strings.Cut(scanner.Text(), ": ")
		if !ok {
			log.Fatalln("Invalid line")
		}
		fields := strings.Fields(rhs)
		graph[lhs] = fields
		nodeSet[lhs] = struct{}{}
		for _, i := range fields {
			inDegree[i]++
			nodeSet[i] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	dag := make([]string, 0, len(nodeSet))

	var openSet []string
	for k := range nodeSet {
		if inDegree[k] == 0 {
			openSet = append(openSet, k)
		}
	}
	for len(openSet) != 0 {
		v := openSet[len(openSet)-1]
		openSet = openSet[:len(openSet)-1]
		dag = append(dag, v)
		for _, i := range graph[v] {
			inDegree[i]--
			if inDegree[i] == 0 {
				openSet = append(openSet, i)
			}
		}
	}

	fmt.Println("Part 1:", calcNumPaths("you", "out", dag, graph))

	fftIdx := slices.Index(dag, "fft")
	if fftIdx < 0 {
		log.Fatalln("fft not in graph")
	}
	dacIdx := slices.Index(dag, "dac")
	if dacIdx < 0 {
		log.Fatalln("dac not in graph")
	}
	first := "fft"
	second := "dac"
	if fftIdx > dacIdx {
		first, second = second, first
	}

	fmt.Println("Part 2:", calcNumPaths("svr", first, dag, graph)*calcNumPaths(first, second, dag, graph)*calcNumPaths(second, "out", dag, graph))
}

func calcNumPaths(start, end string, dag []string, graph map[string][]string) int {
	numPaths := map[string]int{}
	numPaths[start] = 1
	found := false
	for _, i := range dag {
		if i == end {
			return numPaths[end]
		}
		if found || i == start {
			found = true
			for _, j := range graph[i] {
				numPaths[j] += numPaths[i]
			}
		}
	}
	return -1
}
