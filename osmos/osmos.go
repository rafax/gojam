package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	testCases, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < testCases; i++ {
		fmt.Printf("Case #%v: %v\n", i+1, solve(parseTestCase(scanner)))
	}
}

func solve(us int, others []int) int {
	cnt := 0
	fmt.Println(us, others)
	probs := make([]int, len(others))
	if us > others[0] {
		probs[0] = us
	}
	for i := 1; i < len(others); i++ {
		if probs[i-1]+others[i-1] > others[i] {
			probs[i] = probs[i-1] + others[i-1]
		}
	}
	//fmt.Println(probs)
	for i := 0; i < len(probs); i++ {
		if probs[i] <= others[i] {
			cnt++
		}
	}

	return cnt
}

func parseTestCase(scanner *bufio.Scanner) (int, []int) {
	scanner.Scan()
	pair := strings.Split(scanner.Text(), " ")

	us, _ := strconv.Atoi(pair[0])
	scanner.Scan()
	others := []int{}
	for _, m := range strings.Split(scanner.Text(), " ") {
		mote, _ := strconv.Atoi(m)
		others = append(others, mote)
	}
	sort.Ints(others)
	return us, others
}
