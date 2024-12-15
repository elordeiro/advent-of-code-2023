package main

import (
	"adventofcode/utils"
	"fmt"
)

func solve(fileName string) int {
	mat := utils.ReadIntMatrix(fileName)

	var res int

	var predict func([]int) int
	predict = func(row []int) int {
		var diffs []int
		allZeros := true
		for i := 1; i < len(row); i++ {
			diff := row[i] - row[i-1]
			diffs = append(diffs, diff)
			if diff != 0 {
				allZeros = false
			}
		}
		if allZeros {
			return row[len(row)-1]
		}
		return row[len(row)-1] + predict(diffs)
	}

	for _, row := range mat {
		res += predict(row)
	}

	return res
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		{"../test1.txt", 114},
		{"../input.txt", 0},
	}

	for _, test := range tests {
		got := solve(test.fileName)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
