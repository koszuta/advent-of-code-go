package assembunny

type jnzValOp struct {
	val, jumps int
}

func (instr *jnzValOp) Execute(i int, registers map[string]int) int {
	if instr.val != 0 {
		return i + instr.jumps
	}
	return i + 1
}
