package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(b), "\n")

	dirSizes := make(map[string]int)

	var path string
	for _, line := range lines {
		line = strings.TrimSpace(line) // sanitize CRLF

		switch {
		default: // listed file with size
			size, _ := strconv.Atoi(strings.Split(line, " ")[0])
			dirSizes["/"] += size
			if path != "/" {
				for p := path; p != ""; p = path[:strings.LastIndex(p, "/")] {
					dirSizes[p] += size
				}
			}
		case line == "$ cd /":
			path = "/"
		case line == "$ cd ..":
			path = path[:strings.LastIndex(path, "/")]
		case strings.HasPrefix(line, "$ cd "):
			if path != "/" {
				path += "/"
			}
			path += strings.TrimPrefix(line, "$ cd ")

		case line == "$ ls": // unused
		case strings.HasPrefix(line, "dir"): // unused
		}
	}

	var sum int
	for _, size := range dirSizes {
		if size <= 100000 {
			sum += size
		}
	}
	log.Println("sum of dirs of at most 100000:", sum)
}
