package first

import (
	"adventofcode/utils"
	"slices"
)

const (
	North = iota
	South
	West
	East
)

type Cell struct {
	I, J int
}

var dirs = map[int][]byte{
	North: {'|', 'L', 'J', 'S'},
	South: {'|', '7', 'F', 'S'},
	West:  {'-', 'J', '7', 'S'},
	East:  {'-', 'L', 'F', 'S'},
}

func neighbors(mat [][]byte, shape byte, cell Cell) []Cell {
	if shape == '.' {
		return nil
	}
	var cells []Cell
	if cell.I > 0 {
		nn := mat[cell.I-1][cell.J]
		if (shape == 'S' || slices.Contains(dirs[North], shape)) &&
			slices.Contains(dirs[South], nn) {
			cells = append(cells, Cell{cell.I - 1, cell.J})
		}
	}
	if cell.I < len(mat)-1 {
		sn := mat[cell.I+1][cell.J]
		if (shape == 'S' || slices.Contains(dirs[South], shape)) &&
			slices.Contains(dirs[North], sn) {
			cells = append(cells, Cell{cell.I + 1, cell.J})
		}
	}
	if cell.J > 0 {
		wn := mat[cell.I][cell.J-1]
		if (shape == 'S' || slices.Contains(dirs[West], shape)) &&
			slices.Contains(dirs[East], wn) {
			cells = append(cells, Cell{cell.I, cell.J - 1})
		}
	}
	if cell.J < len(mat[0])-1 {
		en := mat[cell.I][cell.J+1]
		if (shape == 'S' || slices.Contains(dirs[East], shape)) &&
			slices.Contains(dirs[West], en) {
			cells = append(cells, Cell{cell.I, cell.J + 1})
		}
	}
	return cells
}

func Solve(fileName string) ([][]byte, []Cell) {
	mat := utils.ReadMatrix(fileName)
	n, m := len(mat), len(mat[0])
	adj := map[Cell][]Cell{}
	var start Cell
	for i, row := range mat {
		for j, shape := range row {
			cell := Cell{i, j}
			adj[cell] = append(adj[cell], neighbors(mat, shape, cell)...)
			if shape == 'S' {
				start = cell
			}
		}
	}

	path := []Cell{start}
	visited := map[Cell]bool{start: true}
	prev := map[Cell]Cell{}

	var dfs func(Cell) bool
	dfs = func(cell Cell) bool {
		if cell.I < 0 || cell.I >= n || cell.J < 0 || cell.J >= m {
			return false
		}

		for _, neighbor := range adj[cell] {
			if !visited[neighbor] {
				visited[neighbor] = true
				prev[neighbor] = cell
				path = append(path, neighbor)
				if dfs(neighbor) {
					return true
				}
				path = path[:len(path)-1]
			} else if prev[cell] != neighbor {
				return true
			}
		}
		return false
	}

	dfs(start)
	return mat, path
}
