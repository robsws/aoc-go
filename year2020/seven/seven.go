package seven

import (
	"aoc-go/files"
	"aoc-go/utils"
	"fmt"
	"regexp"
)

// PartOne - find number of containers that can contain shiny gold ones
func PartOne(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	containerRe := regexp.MustCompile("^[a-z]+ [a-z]+")
	containedRe := regexp.MustCompile("([0-9]+) ([a-z]+ [a-z]+)")
	containedBy := make(map[string][]string)
	for line := range fileStream {
		container := containerRe.FindString(line)
		contained := containedRe.FindAllStringSubmatch(line, -1)
		for _, submatch := range contained {
			containedBy[submatch[2]] = append(containedBy[submatch[2]], container)
		}
	}
	goldContainers := getContainers(containedBy, "shiny gold", true)
	return fmt.Sprint(goldContainers.Len())
}

// PartTwo - find number of containers within shiny gold one
func PartTwo(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	containerRe := regexp.MustCompile("^[a-z]+ [a-z]+")
	containedRe := regexp.MustCompile("([0-9]+) ([a-z]+ [a-z]+)")
	containerOf := make(map[string][]container)
	for line := range fileStream {
		containerColour := containerRe.FindString(line)
		containedColours := containedRe.FindAllStringSubmatch(line, -1)
		for _, submatch := range containedColours {
			amount := utils.MustAtoi(submatch[1])
			containerOf[containerColour] = append(containerOf[containerColour], container{Colour: submatch[2], Amount: amount})
		}
	}
	goldContains := countContained(containerOf, container{Colour: "shiny gold", Amount: 1}, true)
	return fmt.Sprint(goldContains)
}
