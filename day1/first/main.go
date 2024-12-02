package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(input)

	values := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		var l, r rune
		for _, c := range line {
			if c >= '0' && c <= '9' {
				if l == 0 {
					l = c
				}
				r = c
			}
		}
		val, _ := strconv.Atoi(string(l) + string(r))
		values = append(values, val)
	}

	sum := 0
	for _, val := range values {
		sum += val
	}
	fmt.Println(sum)
}
