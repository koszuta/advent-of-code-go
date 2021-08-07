package assembunny

type jnzRegOp struct {
	reg   string
	jumps int
}

func (instr *jnzRegOp) Execute(i int, registers map[string]int) int {
	if registers[instr.reg] != 0 {
		return i + instr.jumps
	}
	return i + 1
}
