package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var cookiesPerMinute = 2.0

func main() {
	var _ = time.Sleep
	scanner := bufio.NewScanner(os.Stdin)
	var testCases int
	scanner.Scan()
	fmt.Fscanf(strings.NewReader(scanner.Text()), "%d", &testCases)
	for i := 0; i < testCases; i++ {
		fmt.Printf("Case #%v: %.7f\n", i+1, solve(parseTestCase(scanner)))
	}
}

func solve(f, p, g float64) float64 {
	cpm := cookiesPerMinute
	t := 0.0
	c := 0.0
	for c < g {
		// time.Sleep(100 * time.Millisecond)
		no_f_t := (g - c) / cpm
		next_f_t := (f - c) / cpm
		//fmt.Printf("\tAt %v we have %v with cpm %v, %v until goal %v until next factory\n", t, c, cpm, no_f_t, next_f_t)
		if no_f_t < 0.000001 {
			return t
		}
		if no_f_t < next_f_t {
			t += no_f_t
			c += cpm * no_f_t
		} else {
			t += next_f_t
			c += cpm * next_f_t
		}
		if next_f_t == 0 {
			t_with_f := (g - c + f) / (cpm + p)
			t_without_f := (g - c) / (cpm)

			// fmt.Printf("\t\tWithout f %v, with f %v", t_without_f, t_with_f)
			if t_with_f < t_without_f {
				// fmt.Print(" Buy\n")
				c -= f
				cpm += p
			} else {
				// fmt.Print(" Not buy\n")
				t += t_without_f
				c += cpm * t_without_f
			}
		}
	}
	return t
}

func parseTestCase(scanner *bufio.Scanner) (float64, float64, float64) {
	scanner.Scan()
	vars := strings.Split(scanner.Text(), " ")

	f, _ := strconv.ParseFloat(vars[0], 64)
	p, _ := strconv.ParseFloat(vars[1], 64)
	g, _ := strconv.ParseFloat(vars[2], 64)

	return f, p, g
}
