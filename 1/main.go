package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var input = parseInput()

	partOne(input)
	partTwo(input)
}

func parseInput() map[int]int {
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return elfSnacks
}

func partOne(elfSnacks map[int]int) {

	var bestPreparedElf int
	for k, v := range elfSnacks {
		if v < elfSnacks[bestPreparedElf] {
			continue
		}

		bestPreparedElf = k
	}

	// fmt.Println(elfSnacks)

	fmt.Printf("best prepared elf: #%d with %d calories\n", bestPreparedElf, elfSnacks[bestPreparedElf])
}

func partTwo(elfSnacks map[int]int) {
	elves := make([][]int, len(elfSnacks))

	i := 0
	for k := range elfSnacks {
		elves[i] = []int{k, elfSnacks[k]}
		i++
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i][1] < elves[j][1]
	})

	fmt.Println("finding top three best prepared elves...")
	var total = 0
	topThree := elves[len(elves)-3:]
	for i, elf := range topThree {
		fmt.Printf("#%d: %d\n", 3-i, elf[1])
		total += elf[1]
	}
	fmt.Printf("total calories: %d\n", total)
}
