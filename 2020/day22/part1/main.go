package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
 *   --- Day 22: Crab Combat ---
 *        --- Part One ---
 *
 *   https://adventofcode.com/2020/day/22
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

	for roundNo := 1; len(deck1) > 0 && len(deck2) > 0; roundNo++ {
		fmt.Printf("\n-- Round %d --\n", roundNo)
		fmt.Println("Player 1's deck:", toString(deck1))
		fmt.Println("Player 2's deck:", toString(deck2))
		card1, card2 := deck1[0], deck2[0]
		deck1, deck2 = deck1[1:], deck2[1:]
		fmt.Println("Player 1 plays:", card1)
		fmt.Println("Player 2 plays:", card2)
		if card1 > card2 {
			fmt.Println("Player 1 wins the round!")
			deck1 = append(deck1, card1, card2)
		} else {
			fmt.Println("Player 2 wins the round!")
			deck2 = append(deck2, card2, card1)
		}
	}

	fmt.Printf("\n\n== Post-game results ==\n")
	fmt.Println("Player 1's deck:", toString(deck1))
	fmt.Println("Player 2's deck:", toString(deck2))

	winningDeck := deck1
	if len(deck2) > 0 {
		winningDeck = deck2
	}

	winningScore := 0
	for i, card := range winningDeck {
		winningScore += card * (len(winningDeck) - i)
	}
	fmt.Println("\nthe winning player's score is", winningScore)
}
