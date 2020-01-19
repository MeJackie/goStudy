package main

import (
	"fmt"
)
/*
动态规划
1、状态定义；

2、状态转移方程；

3、初始化；

4、输出；

5、思考状态压缩。
*/


func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}


// 最长回文子串
func longestPalindromeByDP(s string) string {
	s2 := []rune(s)
	sLen := len(s2)
	dp := make([][]bool, sLen)

	// 初始化
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
	}

	var maxLen int = 1
	var start int = 0

	for j := 1; j < sLen; j++ {
		for i := 0; i < j; i++ {
			if s2[i] == s2[j] {
				if j - i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}

			// 只要dp[i][j] == true 成立， 就表示子串s[i, j]是回文， 此时记录回文长度和起始位置
			if dp[i][j] {
				var curLen int = j - i + 1
				if curLen > maxLen {
					maxLen = curLen
					start = i
				}
			}
		}
	}

	return s[start:start+maxLen]
}

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和
// 动态规划 1.最优子结构 2.无后效性
func maxSubArrayByDP(nums []int) int {
	n := len(nums)
	maxSum := nums[0]
	for i := 1; i < n; i++ {
		// 最优解
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		maxSum = max(nums[i], maxSum)
		fmt.Printf("%d:%d\n", i, maxSum)
	}

	return maxSum
}

// 分治算法
func maxSubArrayByFZ(nums []int) int {
	left := 0
	right := len(nums)-1
	return maxSubArray(nums, left, right)
}

func maxSubArray(nums []int, left int, right int) int {
	// 只有一个元素时跳出
	if right - left == 1 {
		return nums[left]
	}
	center := (left+right)/2
	//fmt.Println(left,"-",center,"-",right)
	maxLeftSum := maxSubArray(nums, left, center)
	maxRightSum := maxSubArray(nums, center, right)

	// 计算左边界最大子序列和
	currSum := 0
	leftSubSum := 0
	for i := center-1; i >= left; i-- {
		currSum += nums[i]
		leftSubSum = max(leftSubSum, currSum)
	}

	// 计算右边界最大子序列和
	currSum = 0
	rightSubSum := 0
	for i := center; i <= right; i++ {
		currSum += nums[i]
		rightSubSum = max(rightSubSum, currSum)
	}

	return max(leftSubSum + rightSubSum, max(maxLeftSum, maxRightSum))
}

// 贪心算法 1.每次最优
func maxSubArrayByTX(nums []int) int {
	n := len(nums)
	currSum := nums[0]
	maxSum := nums[0]
	for i := 1; i < n; i++ {
		currSum = max(nums[i], nums[i] + currSum)
		maxSum = max(maxSum, currSum)
	}

	return maxSum
}


// 爬楼梯
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int,n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		// 最优子结构，状态转移方程
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}

// 买卖股票最佳时机  tp.121
func maxProfit(prices []int) int {
	var minPriceIndex int = 0
	var maxProfit int = 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[minPriceIndex] {    // 寻找最小谷底
			prices[minPriceIndex] = prices[i]
		} else if (prices[i] - prices[minPriceIndex]) > maxProfit {
			maxProfit = prices[i] - prices[minPriceIndex]
		}
	}

	return maxProfit
}

// 非动态规划问题:买卖股票最佳时机 II tp.122
func maxProfit2(prices []int) int {
	return calculate(prices, 0)
}

func calculate(prices []int, s int) int {
	var max int
	n := len(prices)
	if s > n {
		return 0
	}
	for i := s; i < n; i++ {
		var maxProfit int
		for j := i+1; j < n; j++ {
			if prices[i] < prices[j] {
				profit := calculate(prices, j+1) + prices[j] - prices[i]
				if profit > maxProfit {
					maxProfit = profit
				}
			}
		}

		if maxProfit > max {
			max = maxProfit
		}
	}

	return max
}

// 分割等和子集
func canPartition(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}

	var sum int
	for _, num := range nums {
		sum += num
	}

	// 奇数不符合要求
	if (sum & 1) == 1 {
		return false
	}

	var target int = sum/2

	// 创建二维状态数组  行：物品索引  列：容量（包括0）
	dp := make([][]bool, n)
	// 初始化
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, target+1)
	}

	// 先填表格0行，第1个数只能让容积为自己的背包恰好装满
	if nums[0] <= target {
		dp[0][nums[0]] = true
	}

	// 再填后几行
	for i := 1; i < n; i++ {
		for j := 0; j <= target; j++ {
			// 直接从上一行先把结果抄下来，然后再修正
			dp[i][j] = dp[i-1][j]

			if nums[i] == j {
				dp[i][j] = true
				continue
			}

			if nums[i] < j {
				// 状态转换方程，选nums[i]与不选nums[i]
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}

			//// 由于状态转换方程的特殊性，提前结束，可以认为是剪枝操作
			//if dp[i][target] {
			//	return true
			//}
		}

	}

	for _, dp1 := range dp {
		fmt.Print("[")
		for _, dp2  := range dp1 {
			if dp2 == false {
				fmt.Print("F\t")
			} else {
				fmt.Print("true\t")
			}
		}
		fmt.Print("]\n")
	}

	return dp[n-1][target]
}

// 零钱兑换
// 填表得出DP状态转移方程 dp[x] = dp[x] + dp[x-coin]
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		fmt.Println(coin)
		for x := coin; x < amount+1; x++ {
			dp[x] += dp[x-coin]
		}
	}

	return dp[amount]
}

// 大家劫舍 tp.198
func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		// 状态转移方程
		dp[i] = max(dp[i-1], dp[i-2] + nums[i-1])
	}

	return dp[n]
}

// tp.213
// 1. 在不偷窃第一个房子的情况下（即 nums[1:]），最大金额是 p1
// 2. 在不偷窃最后一个房子的情况下（即 nums[:n−1]），最大金额是 p2
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	return max(rob(nums[0:n-1]),rob(nums[1:n]))
}

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

// tp.337
// 1.取本节点+两孙
// 2.不去本节点，取两子
func rob3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	table := make(map[TreeNode]int)
	if table[*root] != 0 {
		return table[*root]
	}

	var money int = root.Val

	// 左孙
	if root.Left != nil {
		money += rob3(root.Left.Right) + rob3(root.Left.Left)
	}

	// 右孙
	if root.Right != nil {
		money += rob3(root.Right.Right) + rob3(root.Right.Left)
	}

	table[*root] = max(money, rob3(root.Right) + rob3(root.Left))
	return table[*root]
}

// 动态规划：
// 0.不偷当前节点=左孩子偷最大值+右孩子偷是最大值
// 1.偷当前节点=左孩子不偷最大值+右孩子不偷是最大值+当前节点值
func rob31(root *TreeNode) int {
	result := rob31Helper(root)
	return max(result[0], result[1])
}

func rob31Helper(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 2)
	left := rob31Helper(root.Left)
	right := rob31Helper(root.Right)

	result[0] = left[1] + right[1]
	result[1] = left[0] + right[0] + root.Val

	return result
}

func createBinaryTreeByFirst(nums []int) *TreeNode {
	n := len(nums)

	if n < 1 {
		return nil
	}

	tree := &TreeNode{
		Val: nums[0],
	}

	for i := 1; i < n; i++ {
		node := &TreeNode{Val:nums[i]}
		insertNode(tree, node)
	}

	return tree
}

func insertNode(tree *TreeNode, node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if tree.Left == nil {
		tree.Left = node
	} else {
	   insertNode(tree.Left, node)
	}

	if tree.Right == nil {
		tree.Right = node
	} else {
	    insertNode(tree.Right, node)
	}

	return tree
}

// 最长上升子序列的长度
func lengthOfLIS(nums []int) int {
	//n := len(nums)
	//var maxLength int
	//for i := 0; i < n-1; i++ {
	//	for j := i+1; j < n; j++ {
	//		nums[]
	//	}
	//}
}


func main() {
	//var p int
	//n := []int{2,7,9,3,1}
	//nums := []int{3,2,3,0,3,0,1}

	nums := []int{10,9,2,5,3,7,101,18}

	fmt.Println(lengthOfLIS(nums))
}
