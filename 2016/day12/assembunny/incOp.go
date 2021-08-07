package assembunny

type incOp struct {
	reg string
}

func (instr *incOp) Execute(i int, registers map[string]int) int {
	registers[instr.reg]++
	return i + 1
}
