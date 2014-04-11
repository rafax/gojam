package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var testCases int
	scanner.Scan()
	fmt.Fscanf(strings.NewReader(scanner.Text()), "%d", &testCases)
	for i := 0; i < testCases; i++ {
		fmt.Printf("Case #%v: %v\n", i+1, solve(parseTestCase(scanner)))
	}
}

func solve(cards []string) int {
	cnt := 0
	sortedUntil := 1
	for sortedUntil < len(cards) {

		if cards[sortedUntil-1] > cards[sortedUntil] {
			for i := 0; i < sortedUntil; i++ {
				if cards[i] > cards[sortedUntil] {
					notOrdered := cards[sortedUntil]
					cards = append(cards[:sortedUntil], cards[sortedUntil+1:]...)
					cards = append(cards[:i], append([]string{notOrdered}, cards[i:]...)...)
					cnt++
				}
			}
		}
		sortedUntil++
	}
	return cnt
}

func parseTestCase(scanner *bufio.Scanner) []string {
	var cardCount int
	scanner.Scan()
	pairs := []string{}
	fmt.Fscanf(strings.NewReader(scanner.Text()), "%d", &cardCount)
	for i := 0; i < cardCount; i++ {
		scanner.Scan()
		pairs = append(pairs, scanner.Text())
	}
	return pairs
}
