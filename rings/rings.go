package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func solve(r, t uint64) uint64 {
	left := uint64(0)
	right := uint64(1)

	for verify(r, t, right) {
		left = right
		right = 2 * right
	}

	for right-left > 1 {
		k := (left + right) / 2
		if verify(r, t, k) {
			left = k
		} else {
			right = k
		}
	}
	return left
}

func verify(r, t, k uint64) bool {
	return 2*k*k+(2*r-1)*k <= t
}

func parseTestCase(scanner *bufio.Scanner) (uint64, uint64) {
	scanner.Scan()
	pair := strings.Split(scanner.Text(), " ")

	l, _ := strconv.ParseUint(pair[0], 10, 64)
	u, _ := strconv.ParseUint(pair[1], 10, 64)
	return l, u
}
