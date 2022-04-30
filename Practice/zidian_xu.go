package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	ss := []string{"Hello", "Sdasd", "China", "Chili", "Japan", "Korean", "Hellosdasd", "Chi"}

// 	sort.Slice(ss, func(i, j int) bool {
// 		l1 := len(ss[i])
// 		l2 := len(ss[j])
// 		if l1 < l2 {
// 			for p := 0; p < l1; p++ {
// 				if ss[i][p] != ss[j][p] {
// 					return ss[i][p] < ss[j][p]
// 				}
// 			}
// 		} else {
// 			for p := 0; p < l2; p++ {
// 				if ss[i][p] != ss[j][p] {
// 					return ss[i][p] < ss[j][p]
// 				}
// 			}
// 		}
// 		return l1 < l2
// 	})

// 	fmt.Println(ss)
// }

func main() {
	ss := []string{"Hello", "Sdasd", "China", "Chili", "Japan", "Korean", "Hellosdasd", "Chi"}

	sort.Strings(ss)
	fmt.Println(ss)
}
