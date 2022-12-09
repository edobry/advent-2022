package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	parseInput(partOne)
	parseInput(partTwo)
}

type Move int

const (
	Rock Move = iota
	Paper
	Scissors
)

func pickMove(move Move, outcome Outcome) Move {
	if outcome == Draw {
		return move
	}
	if outcome == Win {
		return Move((move + 1) % 3)
	}
	return defeats[move]
}

func (m Move) ordinal() int {
	return int(m)
}

func (m Move) String() string {
	return ([]string{
		"rock",
		"paper",
		"scissors",
	})[m]
}

// tried to be clever here with modulo arithmetic
// didnt work naturally
//
// 0 _ 0 =  1 draw
// 1 _ 1 =  1 draw
// 2 _ 2 =  1 draw
// 0 _ 1 =  2 win
// 0 _ 2 =  0 lose
// 1 _ 2 =  2 win
// 1 _ 0 =  0 lose
// 2 _ 0 =  2 win
// 2 _ 1 =  0 lose

type Outcome int

const (
	Loss Outcome = iota
	Draw
	Win
)

var defeats = map[Move]Move{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

func (o Outcome) String() string {
	return ([]string{
		"loss", "draw", "win",
	})[o]
}

func (o Outcome) score() int {
	return int(o) * 3
}

var moveCodes = map[byte]Move{
	'A': Rock,
	'B': Paper,
	'C': Scissors,
	'X': Rock,
	'Y': Paper,
	'Z': Scissors,
}

var outcome = []string{
	"loss", "draw", "win",
}

func fight(moveA Move, moveB Move) Outcome {
	// fmt.Printf("moveA: %d, moveB: %d\n", moveA.ordinal(), moveB.ordinal())
	if moveB == moveA {
		return Draw
	}
	if defeats[moveB] == moveA {
		return Win
	}
	return Loss
	// fmt.Printf("result: %d, outcome: %s\n", result, outcome[result+1])
}

func parseInput(roundBuilder func(string) []Move) {
	// open file stream
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// make sure we close it eventually
	defer file.Close()

	// create stream scanner
	scanner := bufio.NewScanner(file)

	// iterate over scanner chunks(?)
	var rounds [][]Move = make([][]Move, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) != 3 {
			fmt.Printf("invalid round: %s", line)
			panic("")
		}

		rounds = append(rounds, roundBuilder(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	calculateScore(rounds)
}

func calculateScore(rounds [][]Move) int {
	var score int
	for _, round := range rounds {
		// fmt.Printf("move1: %s, move2: %s\n", round[0], round[1])
		outcome := fight(round[0], round[1])
		fmt.Printf("outcome: %s\n", outcome)
		roundScore := round[1].ordinal() + 1 + outcome.score()
		// fmt.Printf("score: %d\n", round[1].ordinal()+1+outcome.score())
		score += roundScore
	}

	// fmt.Println(elfSnacks)

	fmt.Printf("total score: %d\n", score)
	return score
}

func partOne(line string) []Move {
	return []Move{
		moveCodes[line[0]], moveCodes[line[2]],
	}
}

func partTwo(line string) []Move {
	opponentMove := moveCodes[line[0]]

	outcome := map[byte]Outcome{
		'X': Loss,
		'Y': Draw,
		'Z': Win,
	}[line[2]]

	fmt.Printf("opponent move: %s\n", opponentMove)
	fmt.Printf("outcome needed: %s\n", outcome)
	fmt.Printf("move picked: %s\n\n", pickMove(opponentMove, outcome))

	return []Move{
		opponentMove, pickMove(opponentMove, outcome),
	}
}
