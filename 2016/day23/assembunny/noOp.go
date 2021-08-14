package assembunny

func (op *Instruction) nop(i int, registers map[string]int, instructions []Instruction) ([]Instruction, int) {
	return instructions, i + 1
}
