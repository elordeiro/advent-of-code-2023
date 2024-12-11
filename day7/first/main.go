package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	High = iota
	One
	Two
	Three
	Full
	Four
	Five
)

type Hand struct {
	bid, typ int
	cards    []byte
}

var values = map[byte]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	hands := []Hand{}

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, Hand{bid: bid, cards: []byte(parts[0])})
	}

	for i, hand := range hands {
		m := map[byte]int{}
		for _, card := range hand.cards {
			m[card] += 1
		}
		switch len(m) {
		case 1:
			hands[i].typ = Five
		case 2:
			for _, v := range m {
				if v == 1 || v == 4 {
					hands[i].typ = Four
				} else {
					hands[i].typ = Full
				}
				break
			}
		case 3:
			for _, v := range m {
				if v == 3 {
					hands[i].typ = Three
					break
				} else if v == 2 {
					hands[i].typ = Two
					break
				}
			}
		case 4:
			hands[i].typ = One
		case 5:
			hands[i].typ = High

		}
	}

	slices.SortFunc(hands, func(h1, h2 Hand) int {
		if h1.typ != h2.typ {
			return h1.typ - h2.typ
		}
		for i := range 5 {
			if h1.cards[i] != h2.cards[i] {
				return values[h1.cards[i]] - values[h2.cards[i]]
			}
		}
		return 0
	})

	var winnings int
	for i, hand := range hands {
		winnings += hand.bid * (i + 1)
	}

	fmt.Println(winnings)
}
