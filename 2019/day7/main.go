package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func intcodeCPU(instrs []int, in, out chan int, id string) {
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
			log.Panicf("%s unknown read mode %d at %d\n", id, mode, ptr)
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
			log.Panicf("%s unknown write mode %d at %d\n", id, mode, ptr)
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
		paramModes := instrs[ptr] / 100
		opCode := instrs[ptr] % 100
		ptr++
		switch opCode {

		case 1: // add
			ptr = exec3i(ptr, paramModes, addFunc)

		case 2: // multiply
			ptr = exec3i(ptr, paramModes, mulFunc)

		case 3: // input
			time.Sleep(time.Duration(1000000))
			// log.Printf("%s getting input...", id)
			v, ok := <-in
			if !ok {
				goto END
			}
			time.Sleep(time.Duration(1000000))
			// log.Printf("%s ...input is %d", id, v)
			write(v, ptr, paramModes%10)
			ptr++

		case 4: // output
			time.Sleep(time.Duration(1000000))
			v := read(ptr, paramModes%10)
			// log.Printf("%s sending output...", id)
			out <- v
			time.Sleep(time.Duration(1000000))
			// log.Printf("%s ...output is %d", id, v)
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
			goto END

		default:
			log.Println(instrs)
			log.Panicf("%s encountered unknown intcode %d at index %d\n", id, instrs[ptr], ptr)
		}
	}
END:
	// log.Printf("stopping %s...\n", id)
	close(out)
	for range in {
		// log.Printf("%s flushed input value %d...", id, input)
	}
}

func startIntcodeCPU(instrs []int, in chan int, id string) chan int {
	out := make(chan int)
	go intcodeCPU(instrs, in, out, id)
	return out
}

func parseProgram(fileName string) []int {
	program := make([]int, 0, 0)
	{
		file, err := os.Open(fileName)
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
	return program
}

func amplify(phases, remaining, program []int, feedback bool, maxSignal int) int {
	if len(remaining) == 0 {
		inputs := make([]chan int, 0, len(phases))
		outputs := make([]chan int, 0, len(phases))
		for i := range phases {
			in := make(chan int)
			inputs = append(inputs, in)
			out := startIntcodeCPU(program, in, fmt.Sprintf("cpu%d", i))
			outputs = append(outputs, out)
		}

		// Seed each amplifier with its phase value
		for i := 0; i < len(phases); i++ {
			inputs[i] <- phases[i]
		}

		// Initial input for Amp A
		inputs[0] <- 0

		if feedback {
			i := 1
			for {
				j := (len(phases) + i - 1) % len(phases)
				output, ok := <-outputs[j]
				if !ok {
					break
				}
				maxSignal = output
				inputs[i] <- output
				i = (i + 1) % len(phases)
			}
		} else {
			for i := 1; i < len(phases); i++ {
				inputs[i] <- <-outputs[i-1]
			}
			maxSignal = <-outputs[4]
		}
		for _, in := range inputs {
			close(in)
		}
		log.Printf("%v -> %d\n", phases, maxSignal)
	}
	for i, p := range remaining {
		phases = append(phases, p)
		newRemaining := make([]int, 0, 0)
		newRemaining = append(newRemaining, remaining[:i]...)
		newRemaining = append(newRemaining, remaining[i+1:]...)
		max := amplify(phases, newRemaining, program, feedback, maxSignal)
		if max > maxSignal {
			maxSignal = max
		}
		phases = phases[:len(phases)-1]
	}
	return maxSignal
}

func main() {
	fileName := flag.String("file", "input.txt", "input file")
	flag.Parse()
	program := parseProgram(*fileName)

	phases := make([]int, 0, 0)
	remaining := make([]int, 0, 0)
	for i := 0; i < 5; i++ {
		remaining = append(remaining, i)
	}
	log.Printf("max signal without feedback is %d\n\n", amplify(phases, remaining, program, false, 0))

	phases = make([]int, 0, 0)
	remaining = make([]int, 0, 0)
	for i := 5; i < 10; i++ {
		remaining = append(remaining, i)
	}
	log.Printf("max signal with feedback is %d\n", amplify(phases, remaining, program, true, 0))
}
