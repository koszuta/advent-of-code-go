package operation

type Operation interface {
	Execute([]rune) []rune
}

type SwapPositionsOp struct {
	X, Y int
}

func (op *SwapPositionsOp) Execute(arr []rune) []rune {
	r := arr[op.X]
	arr[op.X] = arr[op.Y]
	arr[op.Y] = r
	return arr
}

type SwapLettersOp struct {
	X, Y rune
}

func (op *SwapLettersOp) Execute(arr []rune) []rune {
	xPos, yPos := -1, -1
	for i, r := range arr {
		if r == op.X {
			xPos = i
		}
		if r == op.Y {
			yPos = i
		}
	}
	arr[xPos] = op.Y
	arr[yPos] = op.X
	return arr
}

type RotateLeftOp struct {
	X int
}

func (op *RotateLeftOp) Execute(arr []rune) []rune {
	i := op.X % len(arr)
	arr = append(arr[i:], arr[:i]...)
	return arr
}

type RotateRightOp struct {
	X int
}

func (op *RotateRightOp) Execute(arr []rune) []rune {
	i := len(arr) - (op.X % len(arr))
	arr = append(arr[i:], arr[:i]...)
	return arr
}

type RotateOnLetterOp struct {
	X rune
}

func (op *RotateOnLetterOp) Execute(arr []rune) []rune {
	i := -1
	for j, r := range arr {
		if r == op.X {
			i = j
			break
		}
	}
	if i >= 4 {
		i++
	}
	if i >= 0 {
		i++
	}
	rotRight := RotateRightOp{X: i}
	arr = rotRight.Execute(arr)
	return arr
}

type InverseRotateOnLetterOp struct {
	X rune
}

func (op *InverseRotateOnLetterOp) Execute(arr []rune) []rune {
	for i := range arr {
		rotLeft := RotateLeftOp{X: i}
		try := rotLeft.Execute(arr)
		rotOnLetter := RotateOnLetterOp{X: op.X}
		if string(rotOnLetter.Execute(try)) == string(arr) {
			return try
		}
	}
	return arr
}

type ReverseSpanOp struct {
	X, Y int
}

func (op *ReverseSpanOp) Execute(arr []rune) []rune {
	for i := 0; i <= (op.Y-op.X)/2; i++ {
		arr[op.X+i], arr[op.Y-i] = arr[op.Y-i], arr[op.X+i]
	}
	return arr
}

type MoveOp struct {
	X, Y int
}

func (op *MoveOp) Execute(arr []rune) []rune {
	r := arr[op.X]
	arr = append(arr[:op.X], arr[op.X+1:]...)
	arr = append(arr[:op.Y+1], arr[op.Y:]...)
	arr[op.Y] = r
	return arr
}
