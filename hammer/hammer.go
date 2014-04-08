package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	g float64 = 9.8
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var testCases int
	scanner.Scan()
	fmt.Fscanf(strings.NewReader(scanner.Text()), "%d", &testCases)
	for i := 0; i < testCases; i++ {
		fmt.Printf("Case #%v: %3.8f\n", i+1, solve(parseTestCase(scanner)))
	}
}

func solve(v, d int) float64 {
	sin := float64(d) * g / (float64(v) * float64(v))
	return float64(28.647889756541160438399077407052585166202736233282) * math.Asin(sin)
}

func parseTestCase(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	pair := strings.Split(scanner.Text(), " ")

	v, _ := strconv.Atoi(pair[0])
	d, _ := strconv.Atoi(pair[1])
	return v, d
}
