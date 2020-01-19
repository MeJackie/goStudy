package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func Min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func Max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}


type SortData []int

func (s SortData) Len() int {
	return len(s)
}

func (s SortData) Less(i int, j int) bool {
	return s[i] < s[j]
}

func (s SortData) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j, y := range nums[i+1:] {
			if x+y == target {
				index := []int{i, j + i + 1}
				return index
			}
		}
	}

	return nil
}

func twoSumByMap(nums []int, target int) []int {
	maps := make(map[int]int)
	for index, value := range nums {
		maps[value] = index
	}

	for index_x, x := range nums {
			index_y, isTrue := maps[target - x]
			if isTrue {
				if index_y != index_x {
					index := []int{index_x, index_y}
					return index
				}
			}
	}

	return nil
}

type ListNode struct{
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	headn := head
	l1n := l1
	l2n := l2
	var carry int = 0

	for l1n != nil || l2n != nil {
		var x, y int
		if l1n != nil {
			x = l1n.Val
			headn.Next = &ListNode{Val:x+carry}
		}

		if l2n != nil {
			y = l2n.Val
			headn.Next = &ListNode{Val:y+carry}
		}

		sum := x + y
		if sum < 10 {
			if sum + carry >= 10 {
				headn.Next = &ListNode{Val:sum+carry-10}
				carry = 1
			} else {
				headn.Next = &ListNode{Val:sum+carry}
				carry = 0
			}
		} else {
			headn.Next = &ListNode{Val:sum+carry-10}
			carry = 1
		}

		if l1n != nil {
			l1n = l1n.Next
		}

		if l2n != nil {
			l2n = l2n.Next
		}

		headn = headn.Next
	}

	if carry != 0 {
		headn.Next = &ListNode{Val:1}
	}

	return head.Next
}

func lengthOfLongestSubstring(s string) int {
	var sub string
	s2 := []rune(s)
	maps := make(map[int32]string)
	for i := 0; i < len(s2); i++ {
		if maps[s2[i]] == "" {
			sub = sub + fmt.Sprintf("%c",s2[i])
			maps[s2[i]] = "true"
		} else {
			break
		}
	}


	return len([]rune(sub))
}
// todo !!!
func lengthOfLongestSubstring2(s string) int {
	// 定义游标尺寸大小,游标的左边位置
	window,start := 0,0

	// 循环字符串
	for key := 0; key < len(s); key++ {

		fmt.Println(s[start:key])
		fmt.Println(s[key])
		// 查看当前字符串是否在游标内
		isExist := strings.Index(string(s[start:key]), string(s[key]));

		// 如果不存在游标内部,游标长度重新计算并赋值
		if (isExist == -1) {
			if (key - start + 1 > window) {
				window = key - start + 1
			}
		} else { //存在，游标开始位置更换为重复字符串位置的下一个位置
			start = start + 1 + isExist
		}
	}

	return window
}

func maxArea(height []int) int {
	var maxArea int = 0
	for i := 0; i < len(height)-1; i++ {
		for j := i+1; j < len(height); j++ {
			maxArea = Max(maxArea, (j-i) * Min(height[j],height[i]))
		}
	}

	return maxArea
}
// num := []int{1,8,6,2,5,4,8,3,7}
func maxAreaByTwoPointer(height []int) int  {
	var (
		maxArea int = 0
		l int = 0
		r int = len(height)-1
	)

	for l < r {
		maxArea = Max(maxArea, Min(height[l], height[r]) * (r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}


// nums = [-1, 0, 1, 2, -1, -4]
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var (
		k int
		i int
		j int
		sum int
		result [][]int
	)

	for k = 0; k <= (len(nums) -2); k++ {
		for nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i = k + 1
		j = len(nums) - 1

		for i < j {
			sum = nums[k] + nums[i] + nums[j]
			if sum == 0 {
				result = append(result,[]int{nums[k], nums[i], nums[j]})
				i++
				j--
				// 去除所有重复的i,j
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else if sum < 0 {
				i++
			} else {
				j--
			}
		}
	}

	return result
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		k int
		i int
		j int
		sum int
		result = nums[0] + nums[1] + nums[2]
	)

	for k = 0; k < len(nums); k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i = k + 1
		j = len(nums) - 1

		for i < j {
			sum = nums[k] + nums[i] + nums[j]
			closeAbs := int(math.Abs(float64(target - sum)))
			resultCloseAbs := int(math.Abs(float64(target - result)))
			if closeAbs < resultCloseAbs {
				result = sum
			}

			if sum > target {
				j--
			} else if sum < target {
				i++
			} else {
				break
			}
		}
	}

	return result
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[i] != val {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}

func longestPalindrome(s string) string {
	var max int
	var maxSub string
	s2 := []rune(s)
	for i := 0; i < len(s2); i++ {
		for j := i+1; j <= len(s2); j++ {
			sub := s2[i:j]
			if isPalindrome(sub) && len(sub) > max {
				max = len(sub)
				maxSub = string(sub)
			}
		}
	}

	return maxSub
}

func isPalindrome(s []rune) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

 // 求最长升序子序列长度
func LIS(s []int) int {
	var f [200]int
	for i := 0; i < 200; i++ {
		f[i] = 1
	}
	var ans int
	sLen := len(s)
	for x := 0; x < sLen; x++ {
		for p := 0; p < x; p++ {
			if s[p] < s[x] {
				f[x] = Max(f[x], f[p] + 1)
			}
		}
		fmt.Println(x, "-",f[x])
	}

	for x := 0; x < sLen; x++ {
		ans = Max(ans, f[x])
	}

	return ans
}





func main() {
	s := []int{-2,1,-3,4,-1,2,1,-5,4}
	//s2 := s[0]
	//fmt.Println(s2)
	//os.Exit(1)
	//s := "aabccvbac"
	fmt.Println(maxSubArrayByTX(s))

}
