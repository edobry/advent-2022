package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// open file stream
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// make sure we close it eventually
	defer file.Close()

	// create stream scanner
	scanner := bufio.NewScanner(file)

	var elfSnacks map[int]int = make(map[int]int)

	var currentElfNum = 0
	var currentElfSnacks []int = make([]int, 0)
	// iterate over scanner chunks(?)
	for scanner.Scan() {
		line := scanner.Text()

		// elf over
		if line == "" {
			// fmt.Println("elf block end")
			result := 0
			for _, v := range currentElfSnacks {
				result += v
			}
			elfSnacks[currentElfNum] = result
			currentElfNum++
			currentElfSnacks = make([]int, 0)
			continue
		}

		// same elf
		// fmt.Println(line)
		i, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("provided value is invalid calorie count: %s\n", line)
			panic(err)
		}
		// fmt.Printf("appending %d to currentElfSnacks\n", i)

		currentElfSnacks = append(currentElfSnacks, i)
	}

	var bestPreparedElf int
	for k, v := range elfSnacks {
		if v < elfSnacks[bestPreparedElf] {
			continue
		}

		bestPreparedElf = k
	}

	// fmt.Println(elfSnacks)

	fmt.Printf("best prepared elf: #%d with %d calories", bestPreparedElf, elfSnacks[bestPreparedElf])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
