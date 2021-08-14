package assembunny

import (
	"strconv"
)

func (op *Instruction) add(i int, registers map[string]int, instructions []Instruction) ([]Instruction, int) {
	if _, valid := registers[op.y]; valid {
		if x, valid := registers[op.x]; valid {
			// log.Println("added value of register", op.x, "=", x, "to register", op.y)
			registers[op.y] += x
		} else {
			if x, err := strconv.Atoi(op.x); err == nil {
				// log.Println("added value", x, "to register", op.y)
				registers[op.y] += x
			}
		}
	}
	return instructions, i + 1
}
