package tile

// ImageTile ...
type ImageTile struct {
	ID                                       int
	TopEdge, RightEdge, BottomEdge, LeftEdge string
}

// Rot ...
func (t *ImageTile) Rot() {
	t.TopEdge, t.RightEdge, t.BottomEdge, t.LeftEdge =
		reverse(t.LeftEdge), t.TopEdge, reverse(t.RightEdge), t.BottomEdge
}

// Flip ...
func (t *ImageTile) Flip() {
	t.TopEdge, t.RightEdge, t.BottomEdge, t.LeftEdge =
		t.BottomEdge, reverse(t.RightEdge), t.TopEdge, reverse(t.LeftEdge)
}

func reverse(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}
	return string(r)
}
