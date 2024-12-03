package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(input)

	sum := 0
	for scanner.Scan() {
		card := scanner.Text()
		parts := strings.Split(card, "|")
		winners := atoi(strings.Fields(strings.TrimSpace(parts[0][strings.Index(parts[0], ":")+1:])))
		nums := atoi(strings.Fields(strings.TrimSpace(parts[1])))
		slices.Sort(winners)
		slices.Sort(nums)

		winCount := 0
		for i, j := 0, 0; i < len(winners) && j < len(nums); {
			if winners[i] == nums[j] {
				winCount++
				i++
				j++
				continue
			}
			if winners[i] > nums[j] {
				j++
			} else {
				i++
			}
		}
		sum += int(math.Pow(2, float64(winCount)-1))
	}
	fmt.Println(sum)
}

func atoi(slice []string) []int {
	intSlice := []int{}
	for _, e := range slice {
		num, _ := strconv.Atoi(e)
		intSlice = append(intSlice, num)
	}
	return intSlice
}
