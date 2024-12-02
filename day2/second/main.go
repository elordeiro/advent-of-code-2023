package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(input)

	games := []int{}
	for scanner.Scan() {
		line := scanner.Text()

		mins := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		sets := strings.Split(line[strings.Index(line, ":")+1:], ";")
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				var count int
				var color string
				fmt.Sscanf(cube, "%d %s", &count, &color)
				mins[color] = max(mins[color], count)
			}
		}

		games = append(games, mins["red"]*mins["green"]*mins["blue"])
	}

	sum := 0
	for _, game := range games {
		sum += game
	}
	fmt.Println(sum)
}
