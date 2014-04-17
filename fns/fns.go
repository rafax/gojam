package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var palindromesOfLength map[uint64][]string = map[uint64][]string{}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Text()
	var testCases int
	scanner.Scan()
	fmt.Fscanf(strings.NewReader(scanner.Text()), "%d", &testCases)
	for i := 0; i < testCases; i++ {
		fmt.Printf("Case #%v: %v\n", i+1, solve(parseTestCase(scanner)))
	}
}

func solve(l, u uint64) int {
	cnt := 0
	llen := len(strconv.FormatUint(l, 10))
	ulen := len(strconv.FormatUint(u, 10))

	for i := uint64(llen); i <= uint64(ulen); i++ {
		pals := genPalindromesOfLength(i, func(str string) bool {
			return str[0] != '2' && str[0] != '3' && str[0] != '7' && str[0] != '8'
		})
		for _, v := range pals {
			cand, _ := strconv.ParseUint(v, 10, 64)
			if l <= cand && cand <= u && fairAndSquare(cand) {
				cnt++
			}

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
	str_bound := len(str) - 1
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[str_bound-i] {
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

func genPalindromesOfLength(i uint64, filter func(string) bool) []string {
	pals, present := palindromesOfLength[i]
	if present {
		return pals
	}
	res := []string{}
	if i == 1 {
		for i := int64(0); i < 10; i++ {
			nu := strconv.FormatInt(i, 10)
			if filter == nil || filter(nu) {
				res = append(res, nu)
			}
		}
	} else if i == 2 {
		for i := int64(1); i < 10; i++ {
			str := strconv.FormatInt(i, 10)
			nu := str + str
			if filter == nil || filter(nu) {
				res = append(res, nu)
			}
		}
	} else {
		smaller := genPalindromesOfLength(i-1, filter)
		if i%2 == 0 {
			for _, v := range smaller {
				half := v[:(len(v)/2)+1]
				nu := half + reverse(half)
				if filter == nil || filter(nu) {
					res = append(res, nu)
				}
			}
		} else {
			for i := int64(0); i < 10; i++ {
				for _, v := range smaller {
					half := v[:(len(v) / 2)]
					nu := half + strconv.FormatInt(i, 10) + reverse(half)
					if filter == nil || filter(nu) {
						res = append(res, nu)
					}
				}
			}
		}

	}
	palindromesOfLength[i] = res
	return palindromesOfLength[i]
}

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
