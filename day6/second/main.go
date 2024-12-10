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
	time := atoi(strings.Fields(scanner.Text())[1:])
	scanner.Scan()
	record := atoi(strings.Fields(scanner.Text())[1:])

	var count int

	for t := 1; t < time; t++ {
		speed := t
		raceTime := time - t
		dist := speed * raceTime
		if dist > record {
			count++
		}
	}

	fmt.Println(count)
}

func atoi(slice []string) int {
	str := []byte{}
	for _, e := range slice {
		str = append(str, e...)
	}
	res, _ := strconv.Atoi(string(str))
	return res
}
