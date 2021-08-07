package assembunny

type cpyValOp struct {
	val  int
	dest string
}

func (instr *cpyValOp) Execute(i int, registers map[string]int) int {
	registers[instr.dest] = instr.val
	return i + 1
}
