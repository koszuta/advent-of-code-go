package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	totalAvailableSpace = 70000000
	requiredSpace       = 30000000
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

	unusedSpace := totalAvailableSpace - dirSizes["/"]
	log.Println("current unused space:", unusedSpace)

	neededSpace := requiredSpace - unusedSpace
	log.Println("space needed to free up:", neededSpace)

	smallestDirToDelete := totalAvailableSpace
	for _, size := range dirSizes {
		if size > neededSpace && size < smallestDirToDelete {
			smallestDirToDelete = size
		}
	}
	log.Println("size of smallest dir to delete:", smallestDirToDelete)
}
