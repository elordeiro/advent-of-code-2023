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
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
}

func main() {
	input, err := os.Open("../test.txt")
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
		jCount := m['J']

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
			if jCount > 0 {
				hands[i].typ = Five
			}
		case 3:
			for _, v := range m {
				if v == 3 {
					hands[i].typ = Three
					if jCount > 0 {
						hands[i].typ = Four
					}
					break
				} else if v == 2 {
					hands[i].typ = Two
					if jCount > 0 {
						hands[i].typ += 1 + jCount
					}
					break
				}
			}
		case 4:
			hands[i].typ = One
			if jCount > 0 {
				hands[i].typ = Three
			}
		case 5:
			hands[i].typ = High + jCount
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
