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

	var intervals []Interval
	var nums []int

	first := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if first {
			line := scanner.Text()
			if line == "" {
				first = false
				continue
			}
			as, bs, ok := strings.Cut(scanner.Text(), "-")
			if !ok {
				log.Fatalln("Invalid line")
			}
			a, err := strconv.Atoi(as)
			if err != nil {
				log.Fatalln(err)
			}
			b, err := strconv.Atoi(bs)
			if err != nil {
				log.Fatalln(err)
			}
			intervals = append(intervals, Interval{
				start: a,
				end:   b,
			})
			continue
		}
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	count := 0
	intervals = produceIntervals(processIntervals(intervals))
	for _, i := range nums {
		if binarySearchIntervals(intervals, i) {
			count++
		}
	}
	count2 := 0
	for _, i := range intervals {
		count2 += i.end - i.start
	}
	fmt.Println("Part 1:", count)
	fmt.Println("Part 2:", count2)
}

func binarySearchIntervals(intervals []Interval, num int) bool {
	_, ok := slices.BinarySearchFunc(intervals, num, func(i Interval, t int) int {
		if i.end <= t {
			return -1
		}
		if i.start > t {
			return 1
		}
		return 0
	})
	return ok
}

type (
	Interval struct {
		start int
		end   int
	}

	Divider struct {
		idx   int
		start bool
	}
)

func processIntervals(intervals []Interval) []Divider {
	dividers := make([]Divider, 0, len(intervals)*2)
	for _, i := range intervals {
		dividers = append(dividers, Divider{
			idx:   i.start,
			start: true,
		}, Divider{
			idx:   i.end + 1,
			start: false,
		})
	}
	slices.SortFunc(dividers, func(a, b Divider) int {
		if a.idx == b.idx {
			if a.start == b.start {
				return 0
			}
			if a.start {
				return 1
			}
			return -1
		}
		return a.idx - b.idx
	})
	return dividers
}

func produceIntervals(dividers []Divider) []Interval {
	intervals := make([]Interval, 0, len(dividers))
	lastIdx := 0
	height := 0
	for _, i := range dividers {
		var interval Interval
		intervalHeight := height
		if i.start {
			interval = Interval{
				start: lastIdx,
				end:   i.idx,
			}
			lastIdx = i.idx
			height++
		} else {
			interval = Interval{
				start: lastIdx,
				end:   i.idx,
			}
			lastIdx = i.idx
			height--
		}
		if intervalHeight > 0 && interval.start < interval.end {
			if len(intervals) > 0 {
				lastInterval := intervals[len(intervals)-1]
				if interval.start == lastInterval.end {
					lastInterval.end = interval.end
					intervals[len(intervals)-1] = lastInterval
					continue
				}
			}
			intervals = append(intervals, interval)
		}
	}
	return intervals
}
