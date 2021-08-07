package assembunny

type decOp struct {
	reg string
}

func (instr *decOp) Execute(i int, registers map[string]int) int {
	registers[instr.reg]--
	return i + 1
}
