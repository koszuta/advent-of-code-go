package assembunny

import "strconv"

func (op *Instruction) tgl(i int, registers map[string]int, instructions []Instruction) ([]Instruction, int) {
	x, valid := registers[op.x]
	if !valid {
		var err error
		x, err = strconv.Atoi(op.x)
		if err != nil {
			return instructions, i + 1 // skip instruction
		}
	}

	j := i + x
	if j >= 0 && j < len(instructions) {
		toToggle := instructions[j]
		switch toToggle.name {
		case inc:
			instructions[j].name = dec
		case dec, tgl:
			instructions[j].name = inc
		case cpy:
			instructions[j].name = jnz
		case jnz:
			instructions[j].name = cpy
		}
	}
	return instructions, i + 1
}
