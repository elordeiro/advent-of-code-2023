package main

import (
	"adventofcode/day10/first"
	"adventofcode/day10/second"
	"fmt"
)

func main() {
	tests1 := []struct {
		fileName string
		want     int
	}{
		{"test1.txt", 4},
		{"test2.txt", 8},
		// {"../test3.txt", 0},
		{"input.txt", 6778},
	}

	for _, test := range tests1 {
		_, path := first.Solve(test.fileName)
		got := len(path) / 2
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}

	tests2 := []struct {
		fileName string
		want     int
	}{
		{"test1.txt", 1},
		{"test2.txt", 4},
		{"test3.txt", 8},
		{"input.txt", 433},
	}

	for _, test := range tests2 {
		got := second.Solve(test.fileName)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
