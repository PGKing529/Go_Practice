package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1010100101"
	b := "10101010101"
	res := addBinary(a, b)
	fmt.Println(res)
}

func addBinary(a string, b string) string {
	var reverse func(x string) []byte
	reverse = func(x string) []byte {
		p := []byte(x)
		for i := 0; i < len(x)/2; i++ {
			p[i], p[len(x)-i-1] = p[len(x)-i-1], p[i]
		}
		return p
	}

	b_a := reverse(a)
	b_b := reverse(b)

	add := 0
	var res []int
	if len(b_a) < len(b_b) {
		for i := 0; i < len(b_a); i++ {
			tempa := int(b_a[i] - '0')
			tempb := int(b_b[i] - '0')
			sum := tempa + tempb + add
			tail := sum % 2
			add = sum / 2
			res = append(res, tail)
		}
		for j := len(b_a); j < len(b_b); j++ {
			// tempa := strconv.Atoi(a[i])
			tempb := int(b_b[j] - '0')
			sum := tempb + add
			tail := sum % 2
			add = sum / 2
			res = append(res, tail)
		}
		if add != 0 {
			res = append(res, add)
		}
	} else {
		for i := 0; i < len(b_b); i++ {
			tempa := int(b_a[i] - '0')
			tempb := int(b_b[i] - '0')
			sum := tempa + tempb + add
			tail := sum % 2
			add = sum / 2
			res = append(res, tail)
		}
		for j := len(b_b); j < len(b_a); j++ {
			tempa := int(b_a[j] - '0')
			// tempb := strings.Atoi(b[i])
			sum := tempa + add
			tail := sum % 2
			add = sum / 2
			res = append(res, tail)
		}
		if add != 0 {
			res = append(res, add)
		}
	}

	var ans string
	for i := 0; i < len(res); i++ {
		ans += strconv.Itoa(res[i])
	}
	reverse_ans := reverse(ans)
	return string(reverse_ans)
}
