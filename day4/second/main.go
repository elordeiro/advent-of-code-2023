package main

import (
	"bufio"
	"fmt"
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

	cardCounts := []int{}
	sum := 0

	cardIdx := 0
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

		if cardIdx == len(cardCounts) {
			cardCounts = append(cardCounts, 1)
		}

		count := cardCounts[cardIdx]

		for i := cardIdx + 1; i <= cardIdx+winCount; i++ {
			if i == len(cardCounts) {
				cardCounts = append(cardCounts, count+1)
			} else {
				cardCounts[i] += count
			}
		}
		cardIdx++
	}
	for _, val := range cardCounts {
		sum += val
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
