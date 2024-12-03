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
	parts := make([][]int, n)
	for row := range n {
		parts[row] = make([]int, m)
	}

	idToVal := map[int]int{}

	partId := 1
	for i := 0; i < n; i++ {
		for r := 0; r < m; r++ {
			if !isNumeric(matrix[i][r]) {
				continue
			}
			l := r
			for r < m && isNumeric(matrix[i][r]) {
				r++
			}
			if isPart(i, l, r-1, matrix) {
				for j := l; j < r; j++ {
					parts[i][j] = partId
				}
				val, _ := strconv.Atoi(string(matrix[i][l:r]))
				idToVal[partId] = val
				partId++
			}
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] != '*' {
				continue
			}
			Ids := isGear(i, j, parts)
			if Ids == nil {
				continue
			}
			sum += idToVal[Ids[0]] * idToVal[Ids[1]]
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

func isGear(i, j int, mat [][]int) []int {
	adjs := map[int]bool{}
	if j > 0 && mat[i][j-1] > 0 {
		adjs[mat[i][j-1]] = true
	}
	if j < len(mat[0])-1 && mat[i][j+1] > 0 {
		adjs[mat[i][j+1]] = true
	}
	if i > 0 {
		if mat[i-1][j] > 0 {
			adjs[mat[i-1][j]] = true
		}
		if j > 0 && mat[i-1][j-1] > 0 {
			adjs[mat[i-1][j-1]] = true
		}
		if j < len(mat[0])-1 && mat[i-1][j+1] > 0 {
			adjs[mat[i-1][j+1]] = true
		}
	}
	if i < len(mat)-1 {
		if mat[i+1][j] > 0 {
			adjs[mat[i+1][j]] = true
		}
		if j > 0 && mat[i+1][j-1] > 0 {
			adjs[mat[i+1][j-1]] = true
		}
		if j < len(mat[0])-1 && mat[i+1][j+1] > 0 {
			adjs[mat[i+1][j+1]] = true
		}
	}

	if len(adjs) != 2 {
		return nil
	}

	res := []int{}
	for adj := range adjs {
		res = append(res, adj)
	}
	return res
}
