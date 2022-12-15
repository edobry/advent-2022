package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("part 1 result: %d\n", parseInput(partOne))
	fmt.Printf("part 2 result: %d\n", parseInput(partTwo))
}

// func parseInput() {
func parseInput(predicate func([]int, []int) bool) int {
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

		pair := strings.Split(line, ",")

		fmt.Println(pair)

		elf1Raw := strings.Split(pair[0], "-")
		var elf1 []int
		for _, element := range elf1Raw {
			element, _ := strconv.Atoi(element)
			elf1 = append(elf1, element)
		}

		elf2Raw := strings.Split(pair[1], "-")
		var elf2 []int
		for _, element := range elf2Raw {
			element, _ := strconv.Atoi(element)
			elf2 = append(elf2, element)
		}

		if predicate(elf1, elf2) {
			sum++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func partOne(elf1 []int, elf2 []int) bool {
	if elf1[0] <= elf2[0] {
		fmt.Println("elf1 start smaller")
		if elf1[1] >= elf2[1] {
			fmt.Println("elf1 end larger, +1")
			return true
		} else {
			if elf1[0] == elf2[0] {
				fmt.Println("elf2 end larger. +1")
				return true
			} else {
				fmt.Println("elf2 end larger")
			}

		}
	} else {
		fmt.Println("elf2 start smaller")
		if elf1[1] <= elf2[1] {
			fmt.Println("elf2 end larger, +1")
			return true
		} else {
			fmt.Println("elf1 end smaller")
		}
	}

	return false
}

func partTwo(elf1 []int, elf2 []int) bool {
	if elf1[0] <= elf2[0] {
		fmt.Println("elf1 start smaller")
		if elf1[1] >= elf2[0] {
			fmt.Println("elf1 end larger, +1")
			return true
		}
	} else {
		fmt.Println("elf2 start smaller")
		if elf1[0] <= elf2[1] {
			fmt.Println("elf2 end larger, +1")
			return true
		}
	}

	return false
}
