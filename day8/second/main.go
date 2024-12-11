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

	var starters []string
	for node := range nodes {
		if node[2] == 'A' {
			starters = append(starters, node)
		}
	}

	var stepList []int
	for _, starter := range starters {
		var steps int
		for i := 0; starter[2] != 'Z'; i++ {
			steps++
			if i == n {
				i = 0
			}
			if instructions[i] == 'L' {
				starter = nodes[starter].left
			} else {
				starter = nodes[starter].right
			}
		}
		stepList = append(stepList, steps)
	}

	fmt.Println(lcm(stepList))
}

func lcm(list []int) int {
	result := list[0]
	for _, num := range list[1:] {
		result = result * num / gcd(result, num)
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
