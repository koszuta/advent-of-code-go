package assembunny

import (
	"strconv"
)

func (op *Instruction) cpy(i int, registers map[string]int, instructions []Instruction) ([]Instruction, int) {
	if _, valid := registers[op.y]; valid {
		if x, valid := registers[op.x]; valid {
			// log.Println("copied value of register ", op.x, "=", x, "to register", op.y)
			registers[op.y] = x
		} else {
			if x, err := strconv.Atoi(op.x); err == nil {
				// log.Println("copied value", x, "to register", op.y)
				registers[op.y] = x
			}
		}
	}
	return instructions, i + 1
}
