package assembunny

import (
	"strconv"
)

func (op *Instruction) mul(i int, registers map[string]int, instructions []Instruction) ([]Instruction, int) {
	if _, valid := registers[op.z]; valid {
		x, valid := registers[op.x]
		if !valid {
			var err error
			x, err = strconv.Atoi(op.x)
			if err != nil {
				return instructions, i + 1 // skip instruction
			}
		}
		y, valid := registers[op.y]
		if !valid {
			var err error
			y, err = strconv.Atoi(op.y)
			if err != nil {
				return instructions, i + 1 // skip instruction
			}
		}
		// log.Println("register", op.z, "+=", x, "*", y, "=", x*y)
		registers[op.z] += x * y
	}
	return instructions, i + 1
}
