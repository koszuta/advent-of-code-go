package assembunny

import (
	"strconv"
)

func (op *Instruction) jnz(i int, registers map[string]int, instructions []Instruction) ([]Instruction, int) {
	x, valid := registers[op.x]
	if !valid {
		var err error
		x, err = strconv.Atoi(op.x)
		if err != nil {
			return instructions, i + 1 // skip instruction
		}
	}
	if x != 0 {
		if y, valid := registers[op.y]; valid {
			// log.Println("jnz op val from reg", y)
			return instructions, i + y
		}
		if y, err := strconv.Atoi(op.y); err == nil {
			// log.Println("jnz op val", y)
			return instructions, i + y
		}
	}
	return instructions, i + 1
}
