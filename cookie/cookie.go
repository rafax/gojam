package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errors []string = []string{"Bad magician!", "Volunteer cheated!"}

func main() {
	var _ = time.Sleep
	scanner := bufio.NewScanner(os.Stdin)
	var testCases int
	scanner.Scan()
	fmt.Fscanf(strings.NewReader(scanner.Text()), "%d", &testCases)
	for i := 0; i < testCases; i++ {
		res := solve(parseTestCase(scanner))
		if res < 0 {
			fmt.Printf("Case #%v: %v\n", i+1, errors[-res])
		} else {
			fmt.Printf("Case #%v: %v\n", i+1, res)
		}
	}
}

func solve(a, b []string) int {
	fmt.Println(a, b)
}

func parseTestCase(scanner *bufio.Scanner) ([]string, []string) {
	linea := getSelectedLine(scanner)
	lineb := getSelectedLine(scanner)
	return linea, lineb
}

func getSelectedLine(scanner) []string {
	scanner.Scan()
	lineNo := strconv.Atoi(scanner.Text())
	for i := 0; i < lineNo; i++ {
		scanner.Scan()
	}
	scanner.Scan()
	for _, symbol := range scanner.Text() {
		linea = append(linea, string(symbol))
	}
	for i := 0; i < 4-lineNo; i++ {
		scanner.Scan()
	}
}
