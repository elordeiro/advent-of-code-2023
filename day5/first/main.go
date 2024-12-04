package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type SDR struct {
	src, dst, rng int
}

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	maps := make([][]SDR, 7)

	scanner.Scan()
	seeds := atoi(strings.TrimPrefix(scanner.Text(), "seeds:"))

	mapIdx := -1
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		if !unicode.IsDigit(rune(text[0])) {
			mapIdx++
			continue
		}
		dsr := atoi(text)
		maps[mapIdx] = append(maps[mapIdx], SDR{dsr[1], dsr[0], dsr[2]})
	}

	closest := math.MaxInt

	for _, seed := range seeds {
		for _, m := range maps {
			for _, sdr := range m {
				if seed >= sdr.src && seed <= sdr.src+sdr.rng {
					seed = sdr.dst + (seed - sdr.src)
					break
				}
			}
		}
		closest = min(closest, seed)
	}

	fmt.Println(closest)
}

func atoi(str string) []int {
	strSlice := strings.Split(strings.TrimSpace(str), " ")
	intSlice := []int{}
	for _, e := range strSlice {
		num, _ := strconv.Atoi(e)
		intSlice = append(intSlice, num)
	}
	return intSlice
}
