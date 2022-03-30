package main

import "fmt"

func main() {
	var a, b, x, y int
	fmt.Scanf("%d %d %d %d", &a, &b, &x, &y)
	ans := core(a, b, x, y)
	fmt.Println(ans)
}

func core(a, b, x, y int) int {
	// 先单体
	ans1 := 0
	if a%x != 0 {
		ans1 = ans1 + 1 + a/x
	} else {
		ans1 = ans1 + a/x
	}

	if b%x != 0 {
		ans1 = ans1 + 1 + b/x
	} else {
		ans1 = ans1 + b/x
	}

	// 先aoe 再单体
	di := min(a, b)
	gao := max(a, b)
	ans2 := 0
	if di%y != 0 {
		ans2 = 1 + di/y
	} else {
		ans2 = di / y
	}
	// 烈阳风暴造成的伤害
	damage := ans2 * y

	if gao <= damage {
		return min(ans1, ans2)
	} else {
		// 首先要判断是不是 需要再选择烈阳风暴 即 烈阳风暴伤害比火球还高
		if y > x {
			if (gao-damage)%y != 0 {
				ans2 = ans2 + 1 + (gao-damage)/y
			} else {
				ans2 = ans2 + (gao-damage)/y
			}
		} else {
			if (gao-damage)%x != 0 {
				ans2 = ans2 + 1 + (gao-damage)/x
			} else {
				ans2 = ans2 + (gao-damage)/x
			}
		}
	}
	return min(ans1, ans2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
