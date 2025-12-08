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

	var points []Vec3

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalln(err)
		}
		y, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalln(err)
		}
		z, err := strconv.Atoi(fields[2])
		if err != nil {
			log.Fatalln(err)
		}
		points = append(points, Vec3{
			x: x,
			y: y,
			z: z,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	var edges []Vec3
	for i := range len(points) - 1 {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Vec3{
				x: i,
				y: j,
				z: dist(points[i], points[j]),
			})
		}
	}
	slices.SortFunc(edges, func(a, b Vec3) int {
		return a.z - b.z
	})

	edgesByNode := make([][]int, len(points))
	for _, i := range edges[:1000] {
		edgesByNode[i.x] = append(edgesByNode[i.x], i.y)
		edgesByNode[i.y] = append(edgesByNode[i.y], i.x)
	}

	var components []int

	closedSet := make([]bool, len(points))
	openSet := make([]int, len(points))
	for i := range len(points) {
		if closedSet[i] {
			continue
		}
		components = append(components, dfs(i, edgesByNode, closedSet, openSet[:0]))
	}
	slices.Sort(components)
	fmt.Println("Part 1:", components[len(components)-1]*components[len(components)-2]*components[len(components)-3])

	for _, i := range edges[1000:] {
		edgesByNode[i.x] = append(edgesByNode[i.x], i.y)
		edgesByNode[i.y] = append(edgesByNode[i.y], i.x)
		clear(closedSet)
		if dfs(0, edgesByNode, closedSet, openSet[:0]) == len(points) {
			fmt.Println("Part 2:", points[i.x].x*points[i.y].x)
			break
		}
	}
}

func dfs(node int, edgesByNode [][]int, closedSet []bool, openSet []int) int {
	openSet = openSet[:0]
	closedSet[node] = true
	count := 1
	openSet = append(openSet, node)
	for len(openSet) > 0 {
		n := openSet[len(openSet)-1]
		openSet = openSet[:len(openSet)-1]
		for _, i := range edgesByNode[n] {
			if closedSet[i] {
				continue
			}
			closedSet[i] = true
			count++
			openSet = append(openSet, i)
		}
	}
	return count
}

type (
	Vec3 struct {
		x, y, z int
	}
	Vec2 struct {
		x, y int
	}
)

func dist(a, b Vec3) int {
	dx := (a.x - b.x)
	dy := (a.y - b.y)
	dz := (a.z - b.z)
	return dx*dx + dy*dy + dz*dz
}
