package intcode

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	positionMode int = iota
	immediateMode
	relativeMode
)

const (
	noop int = iota
	addOp
	multiplyOp
	inputOp
	outputOp
	jumpIfTrueOp
	jumpIfFalseOp
	lessThanOp
	equalsOp
	relBaseOffsetOp
	haltOp = 99
)

type ptrRecord struct {
	ptr           int
	offsetRelBase bool
	wrote         bool
}

type memRecord struct {
	addr uint64
	old  int
}

func read(ptr, mode, relBase int, program []int, mem map[uint64]int) int {
	var addr uint64
	switch mode {
	case immediateMode:
		return program[ptr]
	case positionMode:
		addr = uint64(program[ptr])
	case relativeMode:
		addr = uint64(program[ptr] + relBase)
	default:
		log.Println(program)
		log.Panicf("unknown read mode %d at address %d\n", mode, ptr)
	}
	if addr < uint64(len(program)) {
		return program[addr]
	}
	return mem[addr]
}

func write(val, ptr, mode, relBase int, program *[]int, mem *map[uint64]int, writeHist func(memRecord)) {
	var addr uint64
	switch mode {
	case positionMode:
		addr = uint64((*program)[ptr])
	case relativeMode:
		addr = uint64((*program)[ptr] + relBase)
	default:
		log.Println(program)
		log.Panicf("unknown write mode %d at address %d\n", mode, ptr)
	}
	if addr < uint64(len(*program)) {
		writeHist(memRecord{addr: addr, old: (*program)[addr]})
		(*program)[addr] = val
	} else {
		writeHist(memRecord{addr: addr, old: (*mem)[addr]})
		(*mem)[addr] = val
	}
}

func exec3i(ptr, paramModes, relBase int, op func(int, int) int, program *[]int, mem *map[uint64]int, writeHist func(memRecord)) int {
	arg1 := read(ptr, paramModes%10, relBase, *program, *mem)
	paramModes /= 10
	arg2 := read(ptr+1, paramModes%10, relBase, *program, *mem)
	paramModes /= 10
	write(op(arg1, arg2), ptr+2, paramModes%10, relBase, program, mem, writeHist)
	return ptr + 3
}

func addFunc(v1, v2 int) int {
	return v1 + v2
}

func mulFunc(v1, v2 int) int {
	return v1 * v2
}

func ltFunc(v1, v2 int) int {
	if v1 < v2 {
		return 1
	}
	return 0
}

func eqFunc(v1, v2 int) int {
	if v1 == v2 {
		return 1
	}
	return 0
}

func intcode(program []int, in func() int, out func(int)) {
	// Save execution history for rewinding
	writeHist := make([]memRecord, 0, 0)
	doWriteHist := func(record memRecord) {
		writeHist = append(writeHist, record)
	}
	wrote := false
	relBaseHist := make([]int, 0, 0)
	offsetRelBase := false
	ptrHist := make([]ptrRecord, 0, 0)

	// Init CPU memory and pointers
	mem := make(map[uint64]int)
	relBase := 0
	ptr := 0

	// Start execution cycle
	for {
		// Save pointer history
		ptrHist = append(ptrHist, ptrRecord{ptr: ptr, offsetRelBase: offsetRelBase, wrote: wrote})
		offsetRelBase = false
		wrote = false

		paramModes := program[ptr] / 100
		opCode := program[ptr] % 100
		ptr++

		// Handle op code
		switch opCode {

		case addOp:
			ptr = exec3i(ptr, paramModes, relBase, addFunc, &program, &mem, doWriteHist)
			wrote = true

		case multiplyOp:
			ptr = exec3i(ptr, paramModes, relBase, mulFunc, &program, &mem, doWriteHist)
			wrote = true

		case inputOp:
			// Game input will be in {-1, 0, 1}
			input := in()

			// If the input is 2 or more, then that number is the number of previous inputs to undo
			if input > 1 {
				// log.Printf("input=%d\n", input)
				// Rewind until the previous input operation (or we're back to the beginning of the program)
				undos := 0
				for len(ptrHist) > 0 {
					// l := len(writeHist)
					// if l > 0 {
					// 	log.Printf("writes=%d last=%#v\n", l, writeHist[l-1])
					// }
					// log.Printf("ops=%d last=%#v\n", l, ptrHist[len(ptrHist)-1])
					ptrRecord := ptrHist[len(ptrHist)-1]
					if ptrRecord.wrote {
						memRecord := writeHist[len(writeHist)-1]
						addr := memRecord.addr
						if addr < uint64(len(program)) {
							program[addr] = memRecord.old
						} else {
							mem[addr] = memRecord.old
						}
						writeHist = writeHist[:len(writeHist)-1]
					}
					if ptrRecord.offsetRelBase {
						relBase = relBaseHist[len(relBaseHist)-1]
						relBaseHist = relBaseHist[:len(relBaseHist)-1]
					}
					ptr = ptrRecord.ptr
					ptrHist = ptrHist[:len(ptrHist)-1]

					if program[ptr]%100 == inputOp {
						undos++
					}
					if undos > input {
						undos--
						break
					}
				}
			} else {
				write(input, ptr, paramModes%10, relBase, &program, &mem, doWriteHist)
				wrote = true
				ptr++
			}

		case outputOp:
			out(read(ptr, paramModes%10, relBase, program, mem))
			ptr++

		case jumpIfTrueOp:
			if read(ptr, paramModes%10, relBase, program, mem) != 0 {
				ptr = read(ptr+1, paramModes/10%10, relBase, program, mem)
			} else {
				ptr += 2
			}

		case jumpIfFalseOp:
			if read(ptr, paramModes%10, relBase, program, mem) == 0 {
				ptr = read(ptr+1, paramModes/10%10, relBase, program, mem)
			} else {
				ptr += 2
			}

		case lessThanOp:
			ptr = exec3i(ptr, paramModes, relBase, ltFunc, &program, &mem, doWriteHist)
			wrote = true

		case equalsOp:
			ptr = exec3i(ptr, paramModes, relBase, eqFunc, &program, &mem, doWriteHist)
			wrote = true

		case relBaseOffsetOp:
			offsetRelBase = true
			relBaseHist = append(relBaseHist, relBase)
			relBase += read(ptr, paramModes%10, relBase, program, mem)
			ptr++

		case haltOp:
			return

		default:
			log.Println(program)
			log.Panicf("unknown intcode %d at address %d\n", program[ptr], ptr)
		}
	}
}

// Run starts a new intcode computer and executes the program
func Run(program []int, in func() int, out func(int)) {
	intcode(program, in, out)
}

// ParseProgram parses the supplied file as intcode instructions
func ParseProgram(fileName string) []int {
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
