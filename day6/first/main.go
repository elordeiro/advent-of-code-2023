package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	times := ints(strings.Fields(scanner.Text())[1:])
	scanner.Scan()
	dists := ints(strings.Fields(scanner.Text())[1:])

	var counts []int
	n := len(times)

	for race := range n {
		var count int
		duration := times[race]
		record := dists[race]
		for time := 1; time < duration; time++ {
			speed := time
			raceTime := duration - time
			dist := speed * raceTime
			if dist > record {
				count++
			}
		}
		counts = append(counts, count)
	}

	margin := 1
	for _, count := range counts {
		margin *= count
	}

	fmt.Println(margin)
}

func ints(slice []string) []int {
	res := []int{}
	for _, e := range slice {
		val, _ := strconv.Atoi(e)
		res = append(res, val)
	}
	return res
}
