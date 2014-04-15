package main

import (
	"bufio"
	"fmt"
	"math"
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

func solve(l, u uint64) int {
	cnt := 0
	for i := l; i <= u; i++ {
		if fairAndSquare(i) {
			cnt++
		}
	}
	return cnt
}

func fairAndSquare(n uint64) bool {
	if !isPalindrome(n) {
		return false
	}
	sqrt := math.Sqrt(float64(n))
	usqrt := uint64(sqrt)
	if float64(usqrt) != sqrt {
		return false
	}
	return isPalindrome(usqrt)
}

func isPalindrome(n uint64) bool {
	str := strconv.FormatUint(n, 10)
	strlen := len(str)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[strlen-1-i] {
			return false
		}
	}
	return true
}

func parseTestCase(scanner *bufio.Scanner) (uint64, uint64) {
	scanner.Scan()
	pair := strings.Split(scanner.Text(), " ")

	l, _ := strconv.ParseUint(pair[0], 10, 64)
	u, _ := strconv.ParseUint(pair[1], 10, 64)
	return l, u
}
