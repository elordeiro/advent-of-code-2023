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

	matrix := [][]byte{}
	for scanner.Scan() {
		matrix = append(matrix, append([]byte{}, scanner.Text()...))
	}

	n, m := len(matrix), len(matrix[0])
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !isNumeric(matrix[i][j]) {
				continue
			}
			lj := j
			for j < m && isNumeric(matrix[i][j]) {
				j++
			}
			if isPart(i, lj, j-1, matrix) {
				partNum, _ := strconv.Atoi(string(matrix[i][lj:j]))
				sum += partNum
			}
		}
	}
	fmt.Println(sum)
}

func isNumeric(b byte) bool {
	return b >= '0' && b <= '9'
}

func isPart(i, lj, rj int, mat [][]byte) bool {
	if lj > 0 && mat[i][lj-1] != '.' {
		return true
	}
	if rj < len(mat)-1 && mat[i][rj+1] != '.' {
		return true
	}
	for j := lj - 1; j < rj+2; j++ {
		if j < 0 || j >= len(mat[0]) {
			continue
		}
		if i > 0 {
			if b := mat[i-1][j]; b != '.' && !isNumeric(b) {
				return true
			}
		}
		if i < len(mat)-1 {
			if b := mat[i+1][j]; b != '.' && !isNumeric(b) {
				return true
			}
		}
	}
	return false
}
