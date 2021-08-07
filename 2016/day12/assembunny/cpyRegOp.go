package assembunny

type cpyRegOp struct {
	src, dest string
}

func (instr *cpyRegOp) Execute(i int, registers map[string]int) int {
	registers[instr.dest] = registers[instr.src]
	return i + 1
}
