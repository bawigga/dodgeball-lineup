package main

import (
	"fmt"
)

type Player struct {
	Name   string
	gender string
	Power  int
}

type Lineup struct {
	players []Player
	power   int
}

func main() {
	men := []Player{
		// men
		Player{"Cruz", "m", 5},
		Player{"Brian", "m", 4},
		Player{"Clayton", "m", 4},
		Player{"Boase", "m", 3},
		Player{"Deathe", "m", 3},
		Player{"Jesse", "m", 2},
		Player{"Swift", "m", 1},
	}

	women := []Player{
		// women
		Player{"Swift Friend", "f", 4},
		Player{"Michelle", "f", 3},
		Player{"Elisa", "f", 3},
		Player{"Heather", "f", 3},
		Player{"Elliot", "f", 2},
		Player{"Kim", "f", 1},
		Player{"Cassie", "f", 1},
	}

	brute(men, women)
}

// brute returns all possible combinations of lineups sorted by power.
func brute(men []Player, women []Player) int {
	menCombinations := combinations(5, len(men))
	womenCombinations := combinations(3, len(women))
	totalCombinations := menCombinations * womenCombinations
	fmt.Printf("Total Combinations: %v\n", totalCombinations)

	getLineups(men, 5, 0)

	return totalCombinations
}

func getLineups(players []Player, lineupSize int) []Lineup {

	lineups := []Lineup{}
	for _, player := range players {
		fmt.Println(player)
	}
	return lineups
}

// ------------------------

// Alex, Brian, Cruz

// Alex, Brian
// Alex, Cruz
// Brian, Cruz

// ------------------------

// Alex, Brian, Cruz, Clayton

// Alex, Brian
// Alex, Cruz
// Alex, Clayton

// Brian, Cruz
// Brian, Clayton

// Cruz, Clayton

func genLineups(players []Player, idx int, lineupSize int) []Player {
	if len(players) == lineupSize { //} || idx == lineupSize-1 {
		return players
	}
	lineups := []Lineup
	for i, player := range players {
		lineup.append(genLineups(players[i+1:], lineupSize))
	}
}

// func getLineups(players []Player, lineupSize int, i int, lineups []Lineup) {
// 	if i > len(players) {
// 		return
// 	}
// }

func combinations(choose int, setLen int) int {
	if choose > setLen {
		panic(fmt.Sprintf("Not enough players. Requested %v but got %v", choose, setLen))
	}
	return factorial(setLen) / (factorial((setLen - choose)) * factorial(choose))
}

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}
