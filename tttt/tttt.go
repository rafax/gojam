package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var results = []string{"X won", "O won", "Draw", "Game has not completed"}

var xwon = [][]string{[]string{"X", "X", "X", "X"}, []string{"T", "X", "X", "X"},
	[]string{"X", "T", "X", "X"}, []string{"X", "X", "T", "X"}, []string{"X", "X", "X", "T"}}

var owon = [][]string{[]string{"O", "O", "O", "O"}, []string{"T", "O", "O", "O"},
	[]string{"O", "T", "O", "O"}, []string{"O", "O", "T", "O"}, []string{"O", "O", "O", "T"}}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var testCases int
	scanner.Scan()
	fmt.Fscanf(strings.NewReader(scanner.Text()), "%d", &testCases)
	for i := 0; i < testCases; i++ {
		fmt.Printf("Case #%v: %v\n", i+1, results[solve(parseTestCase(scanner))])
	}
}

func solve(grid [][]string) int {
	dotFound := false
	for i := 0; i < 4; i++ {
		res := lineResult(grid[i])
		if res == 0 || res == 1 {
			return res
		}
		if res == 3 {
			dotFound = true
		}
	}
	for i := 0; i < 4; i++ {
		res := lineResult([]string{grid[0][i], grid[1][i], grid[2][i], grid[3][i]})
		if res == 0 || res == 1 {
			return res
		}
		if res == 3 {
			dotFound = true
		}
	}
	res := lineResult([]string{grid[0][0], grid[1][1], grid[2][2], grid[3][3]})
	if res == 0 || res == 1 {
		return res
	}
	if res == 3 {
		dotFound = true
	}
	res = lineResult([]string{grid[0][3], grid[1][2], grid[2][1], grid[3][0]})
	if res == 0 || res == 1 {
		return res
	}
	if res == 3 {
		dotFound = true
	}

	if dotFound {
		return 3
	} else {
		return 2
	}
}

func lineResult(line []string) int {
	if contains(line, ".") {
		return 3
	}
	for _, x := range xwon {
		if equal(line, x) {
			return 0
		}
	}
	for _, y := range owon {
		if equal(line, y) {
			return 1
		}
	}
	return 2
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func parseTestCase(scanner *bufio.Scanner) [][]string {
	grid := [][]string{[]string{}, []string{}, []string{}, []string{}}
	for i := 0; i < 4; i++ {
		scanner.Scan()
		for _, symbol := range scanner.Text() {
			grid[i] = append(grid[i], string(symbol))
		}
	}
	scanner.Scan()
	return grid
}
