/**
 * @author: xv
 * @date: 2021/12/13 22:57
 * @link: 最长不含重复字符的子字符串
 * @tag: 滑动窗口
 */

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo048(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}

func lengthOfLongestSubstring(s string) int {
	myMap := make(map[string]int)
	res := 0
	start := 0
	for i := range s {
		char := string(s[i])
		if value, ok := myMap[char]; ok {
			if value+1 > start {
				start = value + 1
			}
		}
		temp := i - start + 1
		if temp > res {
			res = temp
		}
		myMap[char] = i
	}
	return res
}
