package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func intcodeCPU(instrs []int) {
	scanner := bufio.NewScanner(os.Stdin)

	relBase := 0
	mem := make(map[uint64]int)

	read := func(ptr, mode int) int {
		var addr uint64
		switch mode {
		case 1:
			return instrs[ptr]
		case 0:
			addr = uint64(instrs[ptr])
		case 2:
			addr = uint64(instrs[ptr] + relBase)
		default:
			log.Panicf("unknown read mode %d at %d\n", mode, ptr)
		}
		if addr < uint64(len(instrs)) {
			return instrs[addr]
		}
		return mem[addr]
	}

	write := func(val, ptr, mode int) {
		var addr uint64
		switch mode {
		case 0:
			addr = uint64(instrs[ptr])
		case 2:
			addr = uint64(instrs[ptr] + relBase)
		default:
			log.Panicf("unknown write mode %d at %d\n", mode, ptr)
		}
		if addr < uint64(len(instrs)) {
			instrs[addr] = val
		} else {
			mem[addr] = val
		}
	}

	exec3i := func(ptr, paramModes int, op func(int, int) int) int {
		vals := [2]int{}
		for i := 0; i < 2; i++ {
			vals[i] = read(ptr+i, paramModes%10)
			paramModes /= 10
		}
		write(op(vals[0], vals[1]), ptr+2, paramModes%10)
		return ptr + 3
	}

	addFunc := func(v1, v2 int) int {
		return v1 + v2
	}
	mulFunc := func(v1, v2 int) int {
		return v1 * v2
	}
	ltFunc := func(v1, v2 int) int {
		if v1 < v2 {
			return 1
		}
		return 0
	}
	eqFunc := func(v1, v2 int) int {
		if v1 == v2 {
			return 1
		}
		return 0
	}

	ptr := 0
	for {
		// log.Println(instrs)
		paramModes := instrs[ptr] / 100
		opCode := instrs[ptr] % 100
		ptr++
		switch opCode {

		case 1: // add
			ptr = exec3i(ptr, paramModes, addFunc)

		case 2: // multiply
			ptr = exec3i(ptr, paramModes, mulFunc)

		case 3: // input
			var input int
			{
				if !scanner.Scan() {
					log.Panicln(scanner.Err())
				}
				var err error
				input, err = strconv.Atoi(scanner.Text())
				if err != nil {
					log.Panicln(err)
				}
				log.Printf("<<< %d", input)
			}
			write(input, ptr, paramModes%10)
			ptr++

		case 4: // output
			log.Printf(">>> %d\n", read(ptr, paramModes%10))
			ptr++

		case 5: // jump-if-true
			mode := paramModes % 10
			if read(ptr, mode) != 0 {
				ptr = read(ptr+1, paramModes/10%10)
			} else {
				ptr += 2
			}

		case 6: // jump-if-false
			mode := paramModes % 10
			if read(ptr, mode) == 0 {
				ptr = read(ptr+1, paramModes/10%10)
			} else {
				ptr += 2
			}

		case 7: // less than
			ptr = exec3i(ptr, paramModes, ltFunc)

		case 8: // equals
			ptr = exec3i(ptr, paramModes, eqFunc)

		case 9: // relative base offset
			relBase += read(ptr, paramModes%10)
			ptr++

		case 99: // halt
			return

		default:
			log.Println(instrs)
			log.Panicf("encountered unknown intcode %d at index %d\n", instrs[ptr], ptr)
		}
	}
}

func main() {
	fileName := flag.String("input", "./input.txt", "path to input file")
	flag.Parse()
	log.Printf("fileName=%s\n", *fileName)
	program := make([]int, 0, 0)
	{
		file, err := os.Open(*fileName)
		if err != nil {
			log.Panicln(err)
		}
		programStr, err := bufio.NewReader(file).ReadString(0)
		if err != nil && err != io.EOF {
			log.Panicln(err)
		}
		for _, c := range strings.Split(programStr, ",") {
			intcode, err := strconv.Atoi(c)
			if err != nil {
				log.Panicln(err)
			}
			program = append(program, intcode)
		}
	}

	// For part 1, input 1
	// For part 2, input 2
	instrs := make([]int, len(program))
	copy(instrs, program)
	log.Println(instrs)
	intcodeCPU(instrs)
	// log.Println(instrs)
}
