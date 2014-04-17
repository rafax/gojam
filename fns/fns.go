package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var palindromesOfLength map[int64][]string = map[int64][]string{}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Text()
	// 	var testCases int
	// 	scanner.Scan()
	// 	fmt.Fscanf(strings.NewReader(scanner.Text()), "%d", &testCases)
	// 	for i := 0; i < testCases; i++ {
	// 		fmt.Printf("Case #%v: %v\n", i+1, solve(parseTestCase(scanner)))
	// 	}
	for t := 0; t < 5; t++ {
		start := time.Now()
		for i := 1; i <= 14; i++ {
			ln := int64(i)
			pals := genPalindromesOfLength(ln)
			fmt.Println(ln, len(pals))
		}
		fmt.Println(time.Since(start))
	}
}

func solve(l, u uint64) int {
	cnt := 0
	for i := l; i <= u; {
		if fairAndSquare(i) {
			cnt++
		}
		// i = nextPalindrome(i)
	}
	return cnt
}

func fairAndSquare(n uint64) bool {
	lsd := n % 10
	if lsd == 2 || lsd == 3 || lsd == 7 || lsd == 8 {
		return false
	}
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

func genPalindromesOfLength(i int64) []string {
	pals, present := palindromesOfLength[i]
	if present {
		return pals
	}
	res := []string{}
	if i == 1 {
		for i := int64(0); i < 10; i++ {
			res = append(res, strconv.FormatInt(i, 10))
		}

	} else if i == 2 {
		for i := int64(1); i < 10; i++ {
			str := strconv.FormatInt(i, 10)
			res = append(res, str+str)
		}
	} else {
		smaller := genPalindromesOfLength(i - 1)
		if i%2 == 0 {
			for _, v := range smaller {
				half := v[:(len(v)/2)+1]
				res = append(res, half+reverse(half))
			}
		} else {
			for i := int64(0); i < 10; i++ {
				for _, v := range smaller {
					half := v[:(len(v) / 2)]
					res = append(res, half+strconv.FormatInt(i, 10)+reverse(half))
				}
			}
		}

	}
	sort.Strings(res)
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

func distinct(s []string) []string {
	res := []string{}
	uniq := map[string]bool{}
	for _, v := range s {
		_, present := uniq[v]
		if !present {
			uniq[v] = true
			res = append(res, v)
		}
	}
	return res
}
