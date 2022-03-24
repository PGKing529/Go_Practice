// 根据层序遍历 写出前序遍历 从构建到完成 其中空节点用-1 表示

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func CreateTreeNode(item int) *TreeNode {
	if item == -1 {
		return nil
	}

	return &TreeNode{item, nil, nil}
}

func CreateTree(intputList []int) *TreeNode {
	queue := []*TreeNode{}
	root := CreateTreeNode(intputList[0])
	queue = append(queue, root)

	intputList = intputList[1:]

	for len(intputList) > 0 {
		curNode := queue[0]
		queue = queue[1:]

		curNode.Left = CreateTreeNode(intputList[0])
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}
		intputList = intputList[1:]

		curNode.Right = CreateTreeNode(intputList[0])
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}
		intputList = intputList[1:]
	}

	return root
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	var core func(root *TreeNode)
	core = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		core(root.Left)
		core(root.Right)
	}
	core(root)
	return ans
}

func main() {
	// var n int
	// var p []int
	// fmt.Scanf("%d", &n)
	// p = make([]int, n)
	// fmt.Println(p)
	// for i := 0; i < n; i++ {
	// 	fmt.Scanf("依次输入数组 %d", &p[i])
	// }
	var n int
	fmt.Println("输入数组大小")
	fmt.Scanf("%d", &n)
	fmt.Scanln()
	list := make([]int, n)
	if len(list) > 0 {
		fmt.Println("创建完成")
	}
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scanf("%d", &temp)
		list[i] = temp
	}
	fmt.Println(list)
	root := CreateTree(list)
	arr := inorder(root)
	fmt.Println(arr)
}
