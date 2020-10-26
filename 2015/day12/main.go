package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func containsRed(m map[string]interface{}) bool {
	for _, v := range m {
		switch concrete := v.(type) {
		case string:
			if concrete == "red" {
				return true
			}
		}
	}
	return false
}

func sumNonRed(i interface{}) float64 {
	sum := 0.0

	switch concrete := i.(type) {
	case float64:
		sum += concrete
	case []interface{}:
		for _, v := range concrete {
			sum += sumNonRed(v)
		}
	case map[string]interface{}:
		if !containsRed(concrete) {
			for _, v := range concrete {
				sum += sumNonRed(v)
			}
		}
	}

	return sum
}

func main() {
	bytes, err := ioutil.ReadFile("input.json")
	if err != nil {
		panic(err)
	}

	sum := 0
	for i := 0; i < len(bytes); i++ {
		start := i
		for i < len(bytes) && isDigit(bytes[i]) {
			i++
		}
		if i > 0 && isDigit(bytes[start]) && bytes[start-1] == '-' {
			start--
		}
		if start < i {
			num, err := strconv.Atoi(string(bytes[start:i]))
			if err != nil {
				panic(err)
			}
			sum += num
		}
	}
	fmt.Printf("sum=%d\n", sum)

	var file *os.File
	file, err = os.Open("input.json")
	if err != nil {
		panic(err)
	}
	var jsonMap map[string]interface{}
	err = json.NewDecoder(file).Decode(&jsonMap)
	if err != nil {
		panic(err)
	}

	fmt.Printf("sumNonRed=%d\n", int(sumNonRed(jsonMap)))
}
