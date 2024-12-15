package second

import (
	"adventofcode/day10/first"
)

func Solve(fileName string) int {
	mat, path := first.Solve(fileName)
	inPath := map[first.Cell]bool{}
	n, m := len(mat), len(mat[0])
	mat[path[0].I][path[0].J] = sShape(path[len(path)-1], path[0], path[1])

	for _, cell := range path {
		inPath[cell] = true
	}

	var res int

	for row := 0; row < n; row++ {
		var inLoop bool
		var leftPiece byte
		for col := 0; col < m; col++ {
			cell := first.Cell{I: row, J: col}
			if inPath[cell] {
				shape := mat[row][col]
				switch shape {
				case '|':
					inLoop = !inLoop
				case 'F', 'L':
					leftPiece = shape
				case 'J':
					if leftPiece == 'F' {
						inLoop = !inLoop
					}
				case '7':
					if leftPiece == 'L' {
						inLoop = !inLoop
					}
				}
			} else if inLoop {
				mat[row][col] = 'I'
				res++
			} else {
				mat[row][col] = 'O'
			}
		}
	}

	return res
}

func sShape(prev, start, next first.Cell) byte {
	if prev.I == next.I {
		return '-'
	}
	if prev.J == next.J {
		return '|'
	}
	if prev.I < next.I {
		if prev.J < next.J {
			if prev.I == start.I {
				return 'L'
			}
			return '7'
		}
		if prev.I == start.I {
			return 'F'
		}
		return 'J'
	}
	if prev.J < next.J {
		if prev.I == start.I {
			return 'L'
		}
		return '7'
	}
	if prev.I == start.I {
		return 'J'
	}
	return 'F'
}
