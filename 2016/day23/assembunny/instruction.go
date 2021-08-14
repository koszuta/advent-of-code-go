package assembunny

import (
	"fmt"
	"log"
	"strings"
)

type op string

const (
	nop op = "nop"
	inc op = "inc"
	dec op = "dec"
	cpy op = "cpy"
	jnz op = "jnz"
	tgl op = "tgl"
	add op = "add"
	mul op = "mul"
)

type Instruction struct {
	name    op
	x, y, z string
}

func (op *Instruction) Execute(i int, registers map[string]int, instructions []Instruction) ([]Instruction, int) {
	switch op.name {
	case nop:
		return op.nop(i, registers, instructions)

	case inc:
		return op.inc(i, registers, instructions)

	case dec:
		return op.dec(i, registers, instructions)

	case cpy:
		return op.cpy(i, registers, instructions)

	case jnz:
		return op.jnz(i, registers, instructions)

	case tgl:
		return op.tgl(i, registers, instructions)

	case add:
		return op.add(i, registers, instructions)

	case mul:
		return op.mul(i, registers, instructions)

	default:
		log.Panicln("unhandled operation", op)
		return instructions, -1
	}
}

func Parse(line string, registers map[string]int) (*Instruction, error) {
	parts := strings.Split(line, " ")
	if len(parts) == 0 || len(parts) > 4 {
		return nil, fmt.Errorf("unhandled number of arguments in instruction %s", line)
	}

	op := op(parts[0])
	var x, y, z string
	if len(parts) > 1 {
		x = parts[1]
	}
	if len(parts) > 2 {
		y = parts[2]
	}
	if len(parts) > 3 {
		z = parts[3]
	}
	return &Instruction{name: op, x: x, y: y, z: z}, nil
}
