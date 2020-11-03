package game

import (
	"Advent-Of-Code-Go/2019/day13/intcode"
	"Advent-Of-Code-Go/2019/day13/screen"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Game ...
type Game struct {
	program []int
	screen  screen.Screen
	score   int
}

// New game
func New(fileName string) Game {
	return Game{
		program: intcode.ParseProgram(fileName),
		screen:  screen.New(),
		score:   0,
	}
}

// Play the game
func (g *Game) Play() int {

	i := 0
	vals := [3]int{}
	scanner := bufio.NewScanner(os.Stdin)
	intcode.Run(
		g.program,
		func() int {

			// Draw the screen now that all tiles have been calculated by the computer
			g.screen.Draw()

			// Get "joystick" input
			input := 414
			{
				for input == 414 {
					if !scanner.Scan() {
						log.Panicln(scanner.Err())
					}
					s := scanner.Text()
					if strings.HasPrefix(s, "r") {
						input, _ = strconv.Atoi(s[1:])
					}
					switch s {
					case "a":
						input = -1
					case "":
						input = 0
					case "d":
						input = 1
					}
				}
			}

			return input
		},
		func(v int) {
			vals[i] = v
			i++
			if i == 3 {
				x, y, q := vals[0], vals[1], vals[2]
				if x == -1 && y == 0 {
					g.score = q
					g.screen.SetScore(q)
				} else {
					g.screen.SetTile(x, y, screen.Tile(q))
				}
				i = 0
			}
		})

	// Draw the screen one last time
	g.screen.Draw()
	fmt.Println("\t---  G A M E  O V E R  ---")
	return g.score
}
