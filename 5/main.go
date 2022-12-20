package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	parseInput(partOne)
	parseInput(partTwo)
	// fmt.Printf("part 1 result: %d\n", parseInput(partOne))
	// fmt.Printf("part 2 result: %d\n", parseInput(partTwo))
}

func prependString(slice []string, val string) []string {
	// 0 is used bc zero value is free(?)
	slice = append(slice, "")
	//shifts all values down one slot
	copy(slice[1:], slice)
	// sets value as new first element
	slice[0] = val
	return slice
}

// func parseInput() {
func parseInput(movePopped func([]string, int, map[int][]string)) {
	// open file stream
	file, err := os.Open("input.txt")
	// file, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	// make sure we close it eventually
	defer file.Close()

	// create stream scanner
	scanner := bufio.NewScanner(file)

	// iterate over scanner chunks(?)
	readingStart := true
	readingMiddle := false
	readingMoves := false
	var stacks map[int][]string = make(map[int][]string)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		if readingMoves {
			r := regexp.MustCompile(`move (?P<quantity>\d+) from (?P<from>\d+) to (?P<to>\d+)`)
			match := r.FindStringSubmatch(line)
			// fmt.Printf("%#v\n", r.FindStringSubmatch(line))
			// fmt.Printf("%#v\n", r.SubexpNames())

			result := make(map[string]string)
			for i, name := range r.SubexpNames() {
				if i != 0 && name != "" {
					result[name] = match[i]
				}
			}
			// fmt.Println(result)

			// execute moves

			quantity, err := strconv.Atoi(result["quantity"])
			if err != nil {
				log.Fatal(err)
			}

			fromStack, err := strconv.Atoi(result["from"])
			if err != nil {
				log.Fatal(err)
			}

			toStack, err := strconv.Atoi(result["to"])
			if err != nil {
				log.Fatal(err)
			}
			// fmt.Println(stacks[fromStack])
			// fmt.Println(stacks[toStack])

			popped := stacks[fromStack][0:quantity]
			// fmt.Println(popped)
			stacks[fromStack] = stacks[fromStack][quantity:]
			// fmt.Println(stacks[fromStack])
			movePopped(popped, toStack, stacks)
		}
		if readingMiddle {
			if len(line) != 0 {
				panic(fmt.Sprintf("unexpected non-empty middle line: %s\n", line))
			}
			readingMiddle = false
			readingMoves = true

			fmt.Println("starting to read moves...")
		}
		if readingStart {
			openBrace := false
			for i, x := range line {
				// fmt.Println(x)
				if x == ' ' {
					continue
				}
				if x == '[' {
					openBrace = true
					continue
				}
				if x == ']' {
					openBrace = false
					continue
				}
				if !openBrace {
					if _, err := strconv.Atoi(string(x)); err == nil {
						fmt.Printf("finished reading start\n\n")
						readingStart = false
						readingMiddle = true
						break
					}
					log.Fatal(fmt.Sprintf("char %s is invalid at this position!", string(x)))
				}

				stack := i/4 + 1
				fmt.Printf("found crate %s in stack %d\n", string(x), stack)

				_, ok := stacks[stack]
				if !ok {
					stacks[stack] = []string{string(x)}
				} else {
					stacks[stack] = append(stacks[stack], string(x))
				}

			}
		}

		// fmt.Println(stacks)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// read tops
	fmt.Printf("stacks length: %d\n", len(stacks))
	stackNames := make([]int, len(stacks))
	stackI := 0
	for name, _ := range stacks {
		stackNames[stackI] = name
		stackI++
	}
	sort.Ints(stackNames)
	// fmt.Println(stackNames)

	tops := make([]string, len(stacks))
	for i := range stackNames {
		// fmt.Println(stackNames[i])
		// fmt.Println(stacks[stackNames[i]])
		tops[i] = stacks[stackNames[i]][0]
	}
	fmt.Printf("result: %s", strings.Join(tops, ""))
}

func partOne(popped []string, stack int, stacks map[int][]string) {
	for _, x := range popped {
		stacks[stack] = prependString(stacks[stack], x)
	}
}

func partTwo(popped []string, stack int, stacks map[int][]string) {
	newPopped := make([]string, len(popped))
	copy(newPopped, popped)
	stacks[stack] = append(newPopped, stacks[stack]...)
}
