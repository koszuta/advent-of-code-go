package disc

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	// Width  = 3
	// Height = 3
	Width  = 34
	Height = 30
	dataX  = Width - 1
	dataY  = 0
)

var targetPos = Vec2i{0, 0}

type nodeUsage struct {
	Size, Used, Available int
}

type GridState struct {
	dataPos Vec2i
	Nodes   [Height][Width]nodeUsage
}

func (state *GridState) DataIsAtTarget() bool {
	return state.dataPos == targetPos
}

func (state *GridState) MoveData(currPos, nextPos Vec2i) (newState GridState, moved bool) {
	newState.Nodes = state.Nodes
	if currPos.InBounds() && nextPos.InBounds() && Manhattan(currPos, nextPos) == 1 {
		used := newState.Nodes[currPos.Y][currPos.X].Used
		available := newState.Nodes[nextPos.Y][nextPos.X].Available
		if used <= available {
			if currPos == state.dataPos {
				newState.dataPos = nextPos
			} else {
				newState.dataPos = state.dataPos
			}
			newState.Nodes[currPos.Y][currPos.X].Available += used
			newState.Nodes[currPos.Y][currPos.X].Used -= used
			newState.Nodes[nextPos.Y][nextPos.X].Available -= used
			newState.Nodes[nextPos.Y][nextPos.X].Used += used
			return newState, true
		}
	}
	return newState, false
}

func Parse(filepath string) (state GridState) {
	diskUsageRex := regexp.MustCompile(`/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T\s+\d+%`)
	file, err := os.Open(filepath)
	if err != nil {
		log.Panicln(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := diskUsageRex.FindAllStringSubmatch(line, -1)
		if match != nil {
			x, err := strconv.Atoi(match[0][1])
			if err != nil {
				log.Panicln(err)
			}
			y, err := strconv.Atoi(match[0][2])
			if err != nil {
				log.Panicln(err)
			}
			size, err := strconv.Atoi(match[0][3])
			if err != nil {
				log.Panicln(err)
			}
			used, err := strconv.Atoi(match[0][4])
			if err != nil {
				log.Panicln(err)
			}
			available, err := strconv.Atoi(match[0][5])
			if err != nil {
				log.Panicln(err)
			}
			if x >= Width {
				log.Panicf("x=%d > width=%d - 1\n", x, Width)
			}
			if y >= Height {
				log.Panicf("y=%d > height=%d - 1\n", y, Height)
			}
			state.Nodes[y][x] = nodeUsage{Size: size, Used: used, Available: available}
		}
	}
	state.dataPos = Vec2i{X: dataX, Y: dataY}
	return state
}

func (state *GridState) String() string {
	grid := ""
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			pos := Vec2i{X: x, Y: y}
			node := state.Nodes[y][x]
			if state.dataPos == pos {
				grid += "["
			} else if targetPos == pos {
				grid += "("
			} else {
				// grid += " "
			}
			if node.Used < 10 {
				// grid += " "
			}
			grid += strconv.Itoa(node.Used) + "T/"
			if node.Size < 10 {
				// grid += " "
			}
			grid += strconv.Itoa(node.Available) + "T"
			if state.dataPos == pos {
				grid += "]"
			} else if targetPos == pos {
				grid += ")"
			} else {
				// grid += " "
			}
			if x < Width-1 {
				grid += ","
			}
		}
		if y < Height-1 {
			grid += "\n" //    |            |            |\n"
		}
	}
	return grid
}
