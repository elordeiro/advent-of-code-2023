package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(input)

	atoi := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	values := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		var l, r rune

		assign := func(c rune) {
			if l == 0 {
				l = c
			}
			r = c
		}

		for i, c := range line {
			if c >= '0' && c <= '9' {
				assign(c)
			} else {
				for k, v := range atoi {
					if strings.HasPrefix(line[i:], k) {
						assign(v)
					}
				}
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
