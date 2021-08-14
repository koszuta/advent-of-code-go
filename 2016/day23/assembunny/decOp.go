package assembunny

func (op *Instruction) dec(i int, registers map[string]int, instructions []Instruction) ([]Instruction, int) {
	if _, valid := registers[op.x]; valid {
		registers[op.x]--
	}
	return instructions, i + 1
}
