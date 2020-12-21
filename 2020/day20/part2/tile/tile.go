package tile

import (
	"fmt"
)

// ImageTile ...
type ImageTile struct {
	ID                                       int
	TopEdge, RightEdge, BottomEdge, LeftEdge string
	Image                                    []string
}

// Rot ...
func (t *ImageTile) Rot() {
	t.TopEdge, t.RightEdge, t.BottomEdge, t.LeftEdge =
		reverse(t.LeftEdge), t.TopEdge, reverse(t.RightEdge), t.BottomEdge
	t.Image = RotStrings(t.Image)
}

// Flip ...
func (t *ImageTile) Flip() {
	t.TopEdge, t.RightEdge, t.BottomEdge, t.LeftEdge =
		t.BottomEdge, reverse(t.RightEdge), t.TopEdge, reverse(t.LeftEdge)
	t.Image = FlipStrings(t.Image)
}

// Print ...
func (t *ImageTile) Print() {
	PrintStrings(t.Image)
}

// PrintStrings ...
func PrintStrings(s []string) {
	for _, line := range s {
		fmt.Println(line)
	}
}

// RotStrings ...
func RotStrings(s []string) (rotated []string) {
	l := len(s)
	for j := 0; j < l; j++ {
		line := make([]rune, l, l)
		for i := l - 1; i >= 0; i-- {
			line[l-i-1] = []rune(s[i])[j]
		}
		rotated = append(rotated, string(line))
	}
	return
}

// FlipStrings ...
func FlipStrings(s []string) (flipped []string) {
	for i := len(s) - 1; i >= 0; i-- {
		flipped = append(flipped, s[i])
	}
	return
}

func reverse(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}
	return string(r)
}
