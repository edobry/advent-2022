package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var priorities map[rune]int = make(map[rune]int)
	for ch, i := 'A', 27; ch <= 'Z'; ch++ {
		priorities[ch] = i
		i++
	}
	for ch, i := 'a', 1; ch <= 'z'; ch++ {
		priorities[ch] = i
		i++
	}

	fmt.Printf("part 1 result: %d\n", partOne(priorities))
	// parseInput(partTwo)
}

// func parseInput() {
func partOne(priorities map[rune]int) int {
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
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		size := len(line)

		if size%2 != 0 {
			fmt.Printf("invalid rucksack size, noneven: %s", line)
			panic("")
		}

		fmt.Printf("processing rucksack size %d: %s\n", size, line)

		var firstCompartment map[byte]bool = make(map[byte]bool)
		var secondCompartment map[byte]bool = make(map[byte]bool)
		for i := 0; i < size/2; i++ {
			// fmt.Printf("setting: %d, %d\n", i, i+(size/2-1))
			firstCompartment[line[i]] = true
			secondCompartment[line[i+(size/2)]] = true
		}

		for itemType, _ := range firstCompartment {
			if secondCompartment[itemType] {
				fmt.Printf("item type '%c' duplicated; priority %d\n\n", itemType, priorities[rune(itemType)])
				sum += priorities[rune(itemType)]
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

// func partOne() {

// }

// func partTwo() {

// }
