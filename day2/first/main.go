package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var maxes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(input)

	games := []int{}
	gameNum := 1
	for scanner.Scan() {
		possible := true
		line := scanner.Text()
		sets := strings.Split(line[strings.Index(line, ":")+1:], ";")
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				var count int
				var color string
				fmt.Sscanf(cube, "%d %s", &count, &color)
				if count > maxes[color] {
					possible = false
				}
			}
		}

		if possible {
			games = append(games, gameNum)
		}
		gameNum++
	}

	sum := 0
	for _, game := range games {
		sum += game
	}
	fmt.Println(sum)
}
