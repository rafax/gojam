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
	scanner := bufio.NewScanner(os.Stdin)
	var testCases int
	scanner.Scan()
	fmt.Fscanf(strings.NewReader(scanner.Text()), "%d", &testCases)
	for i := 0; i < testCases; i++ {
		res := solve(parseTestCase(scanner))
		if res <= 0 {
			fmt.Printf("Case #%v: %v\n", i+1, errors[-res])
		} else {
			fmt.Printf("Case #%v: %v\n", i+1, res)
		}
	}
}

func solve(a, b []string) int {
	common := common(a, b)
	if len(common) == 0 {
		return -1
	}
	if len(common) > 1 {
		return 0
	}
	res, _ := strconv.Atoi(common[0])
	return res
}

func common(a, b []string) []string {
	res := []string{}
	amap := map[string]bool{}
	for _, v := range a {
		amap[v] = true
	}
	for _, v := range b {
		if amap[v] {
			res = append(res, v)
		}
	}
	return res
}

func parseTestCase(scanner *bufio.Scanner) ([]string, []string) {
	linea := getSelectedLine(scanner)
	lineb := getSelectedLine(scanner)
	return linea, lineb
}

func getSelectedLine(scanner *bufio.Scanner) []string {
	scanner.Scan()
	lineNo, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < lineNo; i++ {
		scanner.Scan()
	}
	res := strings.Split(scanner.Text(), " ")
	for i := 0; i < 4-lineNo; i++ {
		scanner.Scan()
	}
	return res
}
