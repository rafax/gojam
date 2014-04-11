package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type hatePair struct {
	a string
	b string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var testCases int
	scanner.Scan()
	fmt.Fscanf(strings.NewReader(scanner.Text()), "%d", &testCases)
	for i := 0; i < testCases; i++ {
		res := solve(parseTestCase(scanner))
		if res {
			fmt.Fprintf(os.Stdout, "Case #%v: Yes\n", i+1)
		} else {
			fmt.Fprintf(os.Stdout, "Case #%v: No\n", i+1)

		}
	}
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text()) // Println will add back the final '\n'
	// }
	// fmt.Scanf("%d", &testCases)
	// for i := 0; i < testCases; i++ {
	// 	var pairCount int
	// 	fmt.Scanf("%d", &pairCount)
	// 	pairs := []hatePair{}
	// 	for i := 0; i < pairCount; i++ {

	// 	}
	// }
}

func solve(pairs []hatePair) bool {
	hates := map[string][]string{}
	for _, pair := range pairs {
		hates[pair.a] = append(hates[pair.a], pair.b)
		hates[pair.b] = append(hates[pair.b], pair.a)
	}
	return can2color(2, pairs[0].a, initColor(hates), hates)
}

func can2color(currColor int, person string, color map[string]int, hates map[string][]string) bool {
	if color[person] == 0 {
		color[person] = currColor
	}
	for _, h := range hates[person] {
		if color[h] == 0 {
			can2color(3-currColor, h, color, hates)
		}
		if color[h] == color[person] {
			return false
		}
	}
	for k, v := range hates {
		kcolor := color[k]
		for _, h := range v {
			if kcolor == color[h] {
				return false
			}
		}
	}
	return true

}

func initColor(hates map[string][]string) map[string]int {
	color := map[string]int{}
	for k, _ := range hates {
		color[k] = 0
	}
	return color
}

func parseTestCase(scanner *bufio.Scanner) []hatePair {
	var pairCount int
	scanner.Scan()
	pairs := []hatePair{}
	fmt.Fscanf(strings.NewReader(scanner.Text()), "%d", &pairCount)
	for i := 0; i < pairCount; i++ {
		scanner.Scan()
		pairs = append(pairs, *newHatePair(scanner.Text()))
	}
	return pairs
}

func newHatePair(line string) *hatePair {
	pair := strings.Split(line, " ")
	return &hatePair{a: pair[0], b: pair[1]}
}
