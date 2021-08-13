package disc

type Vec2i struct {
	X, Y int
}

func Manhattan(p1, p2 Vec2i) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func (v *Vec2i) InBounds() bool {
	return v.X >= 0 && v.Y >= 0 && v.X < Width && v.Y < Height
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
