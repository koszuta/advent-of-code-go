package assembunny

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction interface {
	Execute(int, map[string]int) int
}

func Parse(line string, registers map[string]struct{}) (Instruction, error) {
	parts := strings.Split(line, " ")
	op := parts[0]
	x := parts[1]

	switch op {
	case "cpy":
		y := parts[2]
		_, found := registers[y]
		if !found {
			return nil, fmt.Errorf("invalid register %s in instruction %s", x, line)
		}
		val, err := strconv.Atoi(x)
		if err != nil {
			_, found := registers[x]
			if !found {
				return nil, fmt.Errorf("invalid register %s in instruction %s", x, line)
			}
			return &cpyRegOp{src: x, dest: y}, nil
		} else {
			return &cpyValOp{val: val, dest: y}, nil
		}

	case "inc":
		_, found := registers[x]
		if !found {
			return nil, fmt.Errorf("invalid register %s in instruction %s", x, line)
		}
		return &incOp{reg: x}, nil

	case "dec":
		_, found := registers[x]
		if !found {
			return nil, fmt.Errorf("invalid register %s in instruction %s", x, line)
		}
		return &decOp{reg: x}, nil

	case "jnz":
		y := parts[2]
		jumps, err := strconv.Atoi(y)
		if err != nil {
			return nil, err
		}
		val, err := strconv.Atoi(x)
		if err != nil {
			_, found := registers[x]
			if !found {
				return nil, fmt.Errorf("invalid register %s in instruction %s", x, line)
			}
			return &jnzRegOp{reg: x, jumps: jumps}, nil
		} else {
			return &jnzValOp{val: val, jumps: jumps}, nil
		}
	default:
		return nil, fmt.Errorf("couldn't parse instruction %s", line)
	}
}
