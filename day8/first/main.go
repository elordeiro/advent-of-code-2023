package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left, right string
}

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	instructions := scanner.Text()
	n := len(instructions)

	nodes := map[string]Node{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		parts := strings.Split(scanner.Text(), "=")
		node := strings.TrimSpace(parts[0])
		parts = strings.Split(parts[1], ",")
		left := strings.Trim(parts[0], "( )")
		right := strings.Trim(parts[1], "( )")

		nodes[node] = Node{left, right}
	}

	var steps int
	for i, cur := 0, "AAA"; cur != "ZZZ"; i++ {
		steps++
		if i == n {
			i = 0
		}
		if instructions[i] == 'L' {
			cur = nodes[cur].left
		} else {
			cur = nodes[cur].right
		}
	}

	fmt.Println(steps)
}
