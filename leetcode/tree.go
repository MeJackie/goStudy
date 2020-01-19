package main

import (
	"fmt"
	"strconv"
)

// 二叉搜索树中第K小的元素    二叉搜索树：1.left < k < right  2.中序遍历为递增序列
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
var res int  // 记录结果
var count int    // 记录当前遍历到的几点

func kthSmallest(root *TreeNode, k int) int {
	count = k
	inorder(root)
	return res
}

func inorder(root *TreeNode) {
	if root == nil || count == 0 {
		return
	}
	inorder(root.Left)
	count -= 1
	if  count == 0 {
		res = root.Val
	}
	inorder(root.Right)
}


// 二叉树的最近公共祖先
var ans *TreeNode  // 公共祖先

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	recurseTree(root, p, q)
	return ans
}

func recurseTree(root, p, q *TreeNode) bool {
	if root == nil {
		return false
	}

	left := 0
	if recurseTree(root.Left, p, q) {
		left = 1
	}

	right := 0
	if recurseTree(root.Right, p, q) {
		right = 1
	}

	mid := 0
	if root == q || root == p {   // p、q所在的层
		mid = 1
	}

	if (left+right+mid) >= 2 {    // 公共祖先层
		ans = root
	}

	return (left+right+mid) > 0   // 从p、q层回溯每层都返回True
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}

	return nil
}

// 二叉树的序列化与反序列化

func serialize(root *TreeNode) string {
	str := ""
	if root == nil {
		str += "#!"
		return str
	}

	str += strconv.Itoa(root.Val)
	str += serialize(root.Left)
	str += serialize(root.Right)
	return str
}

func deserialize(str string) *TreeNode {
	strRune := []rune(str)
	root := &TreeNode{}
	deserializeHelper(strRune, root)
	return root
}

var index int

func deserializeHelper(s []rune, root *TreeNode) {
	if len(s) >= index || string(s) == "#!" {
		return
	}
	root.Val = int(s[index])
	index++
	deserializeHelper(s, root.Left)
	index++
	deserializeHelper(s, root.Right)
}


// 二叉树中序遍历
func inorderTraversal(root *TreeNode) []int {
	return helper(root)
}

func helper (root *TreeNode) []int {
	var leftSlice []int
	var rightSlice []int
	if root != nil {
		if root.Left != nil {
			leftSlice = helper(root.Left)
		}
		if root.Right != nil {
			rightSlice = helper(root.Right)
		}
	}

	return append(append(leftSlice, root.Val), rightSlice...)
}

func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	leftSlice := inorderTraversal2(root.Left)
	rightSlice := inorderTraversal2(root.Right)
	return append(append(leftSlice, root.Val), rightSlice...)
}


func main() {
	var result []int
	a := []int{2,3,5}
	b := []int{6,7,8}

	c := 0

	fmt.Println(append(append(a, c),b...))

}
