package screen

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
)

// Tile represents a tile that can be drawn to a screen
type Tile int

// Tiles that can be drawn to a screen
const (
	Empty Tile = iota
	Wall
	Block
	Paddle
	Ball
)

// Screen ...
type Screen struct {
	tiles map[coord2D]Tile
	score int
	m     sync.Mutex
	w     *bufio.Writer
}

type coord2D struct {
	x, y int
}

// New screen
func New() Screen {
	return Screen{
		tiles: make(map[coord2D]Tile),
		w:     bufio.NewWriter(os.Stdout),
	}
}

// SetTile on the screen
func (s *Screen) SetTile(x, y int, tile Tile) {
	// s.m.Lock()
	s.tiles[coord2D{x, y}] = tile
	// s.m.Unlock()
}

// SetScore on the screen
func (s *Screen) SetScore(score int) {
	s.score = score
}

// Draw the screen
func (s *Screen) Draw() {
	// s.m.Lock()
	defer s.w.Flush()

	lastY := -1
	for _, c := range sortedCoords(s.tiles) {
		if c.y != lastY {
			s.w.WriteRune('\n')
			lastY = c.y
		}
		s.printTile(c.x, c.y, s.tiles[c])
	}
	s.w.WriteString(fmt.Sprintf("\n\t---  S C O R E : %d  ---\n", s.score))
	// s.m.Unlock()
}

func sortedCoords(tiles map[coord2D]Tile) []coord2D {
	coords := make([]coord2D, 0, len(tiles))
	for c := range tiles {
		coords = append(coords, c)
	}
	sort.Slice(coords, func(i, j int) bool {
		return coords[i].y < coords[j].y || (coords[i].y == coords[j].y && coords[i].x < coords[j].x)
	})
	return coords
}

func (s *Screen) printTile(x, y int, tile Tile) {
	switch tile {
	case Empty:
		s.w.WriteRune(' ')
	case Wall:
		s.w.WriteRune('▒')
	case Block:
		s.w.WriteRune('▄')
	case Paddle:
		s.w.WriteRune('▔')
	case Ball:
		s.w.WriteRune('◉')
	default:
		log.Panicf("unknown tile ID %d\n", tile)
	}
}
