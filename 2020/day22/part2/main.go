package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
 *   --- Day 22: Crab Combat ---
 *        --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/22#part2
 */

func toString(deck []int) string {
	s := ""
	if len(deck) == 0 {
		return s
	}
	for _, card := range deck {
		s += strconv.Itoa(card) + ", "
	}
	return s[:len(s)-2]
}

func hashDecks(deck1, deck2 []int) string {
	hash := make([]rune, 0, 128)
	for _, card := range deck1 {
		hash = append(hash, rune(card), ',')
	}
	hash = append(hash, '|')
	for _, card := range deck2 {
		hash = append(hash, rune(card), ',')
	}
	return string(hash)
}

func recursiveCombat(deck1, deck2 []int, roundNo, gameNo int, deckConfigs map[string]struct{}) (playerOneWon bool, winningDeck []int) {
	if len(deck2) == 0 {
		// fmt.Println("The winner of game", gameNo, "is player 1!")
		// fmt.Printf("\n...anyway, back to game %d.\n", gameNo-1)
		return true, deck1
	}
	if len(deck1) == 0 {
		// fmt.Println("The winner of game", gameNo, "is player 2!")
		// fmt.Printf("\n...anyway, back to game %d.\n", gameNo-1)
		return false, deck2
	}

	hash := hashDecks(deck1, deck2)
	_, seen := deckConfigs[hash]
	if seen {
		// fmt.Println("In order to stop infinite games...")
		// fmt.Println("The winner of game", gameNo, "is player 2!")
		// fmt.Printf("\n...anyway, back to game %d.\n", gameNo-1)
		return true, deck1
	}
	deckConfigs[hash] = struct{}{}

	// fmt.Printf("\n-- Round %d (Game %d) --\n", roundNo, gameNo)
	// fmt.Println("Player 1's deck:", toString(deck1))
	// fmt.Println("Player 2's deck:", toString(deck2))

	card1, card2 := deck1[0], deck2[0]
	deck1, deck2 = deck1[1:], deck2[1:]

	// fmt.Println("Player 1 plays:", card1)
	// fmt.Println("Player 2 plays:", card2)
	if len(deck1) >= card1 && len(deck2) >= card2 {
		subDeck1 := make([]int, card1)
		subDeck2 := make([]int, card2)
		copy(subDeck1, deck1[:card1])
		copy(subDeck2, deck2[:card2])
		// fmt.Printf("Playing a sub-game to determine the winner...\n\n")
		// fmt.Printf("=== Game %d ===\n", gameNo+1)
		playerOneWon, _ = recursiveCombat(subDeck1, subDeck2, 1, gameNo+1, make(map[string]struct{}))

	} else {
		playerOneWon = card1 > card2
	}

	if playerOneWon {
		// fmt.Printf("Player 1 wins round %d of game %d!\n", roundNo, gameNo)
		deck1 = append(deck1, card1, card2)

	} else {
		// fmt.Printf("Player 2 wins round %d of game %d!\n", roundNo, gameNo)
		deck2 = append(deck2, card2, card1)
	}

	return recursiveCombat(deck1, deck2, roundNo+1, gameNo, deckConfigs)
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	var deck1, deck2 []int
	{
		deck := make([][]int, 2, 2)
		deck[0] = make([]int, 0, 0)
		deck[1] = make([]int, 0, 0)

		var idx int
		for scanner.Scan() {
			switch scanner.Text() {
			case "Player 1:":
				idx = 0
				continue
			case "Player 2:":
				idx = 1
				continue
			case "":
				continue
			}

			card, _ := strconv.Atoi(scanner.Text())
			deck[idx] = append(deck[idx], card)
		}
		deck1, deck2 = deck[0], deck[1]
	}

	start := time.Now()
	playerOneWon, winningDeck := recursiveCombat(deck1, deck2, 1, 1, make(map[string]struct{}))
	fmt.Printf("\nthe game of recursive combat took %v\n", time.Since(start))
	if playerOneWon {
		deck1, deck2 = winningDeck, make([]int, 0, 0)
	} else {
		deck2, deck1 = winningDeck, make([]int, 0, 0)
	}

	fmt.Printf("\n\n== Post-game results ==\n")
	fmt.Println("Player 1's deck:", toString(deck1))
	fmt.Println("Player 2's deck:", toString(deck2))

	winningScore := 0
	for i, card := range winningDeck {
		winningScore += card * (len(winningDeck) - i)
	}
	fmt.Println("\nthe winning player's score is", winningScore)
}
