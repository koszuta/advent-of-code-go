package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coord2D struct {
	x, y int
}

type blackWhite int

const (
	black blackWhite = iota
	white
)

type cardinal uint8

const (
	north cardinal = 0
	east  cardinal = 1
	south cardinal = 2
	west  cardinal = 3
)

func turnRight(dir cardinal) cardinal {
	return (dir + 1) % 4
}

func turnLeft(dir cardinal) cardinal {
	return (dir + 3) % 4
}

func move(x, y int, dir cardinal) (int, int) {
	switch dir {
	case north:
		return x, y + 1
	case east:
		return x + 1, y
	case south:
		return x, y - 1
	case west:
		return x - 1, y
	}
	return x, y
}

func intcodeCPU(instrs []int, in, out chan int) {
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
		paramModes := instrs[ptr] / 100
		opCode := instrs[ptr] % 100
		ptr++
		switch opCode {

		case 1: // add
			ptr = exec3i(ptr, paramModes, addFunc)

		case 2: // multiply
			ptr = exec3i(ptr, paramModes, mulFunc)

		case 3: // input
			v, ok := <-in
			if !ok {
				close(out)
				return
			}
			write(v, ptr, paramModes%10)
			ptr++

		case 4: // output
			out <- read(ptr, paramModes%10)
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
			close(out)
			return

		default:
			log.Println(instrs)
			log.Panicf("encountered unknown intcode %d at index %d\n", instrs[ptr], ptr)
		}
	}
}

func startIntcodeCPU(instrs []int, in chan int) chan int {
	out := make(chan int)
	go intcodeCPU(instrs, in, out)
	return out
}

func main() {
	fileName := flag.String("file", "./input.txt", "path to input file")
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

	x, y, dir := 0, 0, north
	panels := make(map[coord2D]blackWhite)
	panels[coord2D{x, y}] = white
	{
		in := make(chan int, 2)
		out := startIntcodeCPU(program, in)
		in <- int(panels[coord2D{x, y}])

		v1 := -1
		for v2 := range out {
			// log.Printf(">>> %d\n", v2)
			if v1 != -1 {
				color := blackWhite(v1)
				panels[coord2D{x, y}] = color

				if v2 == 0 {
					dir = turnLeft(dir)
				} else {
					dir = turnRight(dir)
				}
				x, y = move(x, y, dir)

				val := int(panels[coord2D{x, y}])
				// log.Printf("<<< %d\n", val)
				in <- val

				v1 = -1
			} else {
				v1 = v2
			}
		}
	}
	log.Printf("%d panels painted\n", len(panels))

	coords := make([]coord2D, 0, len(panels))
	for coord := range panels {
		coords = append(coords, coord)
	}
	sort.Slice(coords, func(i, j int) bool {
		return coords[i].y > coords[j].y || (coords[i].y == coords[j].y && coords[i].x < coords[j].x)
	})
	log.Println(coords)

	lastX := 0
	lastY := 0
	for _, coord := range coords {
		if coord.y != lastY {
			fmt.Println()
			lastY = coord.y
			lastX = 0
		}
		for i := 0; i < coord.x-lastX; i++ {
			fmt.Printf("░░")
		}
		switch panels[coord] {
		case black:
			fmt.Printf("░░")
		case white:
			fmt.Printf("██")
		}
		lastX = coord.x
	}
}
