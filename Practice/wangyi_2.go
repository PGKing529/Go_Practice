package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	// fmt.Println('a' - 'a')
	ans := core(s)
	fmt.Println(ans)
}

func core(s string) int {
	//
	if len(s) <= 1 {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 0
	dp[1] = sure(s[1], s[0])
	for i := 2; i < len(s); i++ {
		dp[i] = max(dp[i-2]+sure(s[i], s[i-1]), dp[i-1])
	}
	return dp[len(dp)-1]
}

func sure(a, b byte) int {
	ans := 0
	if a >= b && a-b <= 1 {
		ans = int(a+b-'a'-'a') + 2
	} else if a <= b && b-a <= 1 {
		ans = int(a+b-'a'-'a') + 2
	} else {
		ans = 0
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
