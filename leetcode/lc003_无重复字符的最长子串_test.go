/**
 * @author: xv
 * @date: 2021/11/28 16:45
 * @link: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/submissions/
 * @tag: 滑动串口
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc003(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
}

func lengthOfLongestSubstring(s string) int {
	myMap := make(map[string]int)
	start := 0
	res := 0
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		value, ok := myMap[char]
		if ok {
			// 想清楚这一步
			start = maxInt(value+1, start)
		}
		res = maxInt(res, i-start+1)
		myMap[char] = i
	}
	return res
}

func maxInt(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	} else {
		return val2
	}
}
