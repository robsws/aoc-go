package six

import (
	"aoc-go/files"
	"fmt"
)

// PartOne - count yes answers
func PartOne(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	total := 0
	group := make(map[rune]int)
	for line := range fileStream {
		if line == "" {
			for _, n := range group {
				total += n
			}
			group = make(map[rune]int)
			continue
		}
		for i := 0; i < len(line); i++ {
			group[rune(line[i])] = 1
		}
	}
	for _, n := range group {
		total += n
	}
	return fmt.Sprint(total)
}

// PartTwo - count answers which everyone in group answered yes
func PartTwo(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	total := 0
	group := make(map[rune]int)
	groupsize := 0
	for line := range fileStream {
		if line == "" {
			for _, n := range group {
				if n == groupsize {
					total++
				}
			}
			group = make(map[rune]int)
			groupsize = 0
			continue
		}
		groupsize++
		for i := 0; i < len(line); i++ {
			group[rune(line[i])]++
		}
	}
	for _, n := range group {
		if n == groupsize {
			total++
		}
	}
	return fmt.Sprint(total)
}
