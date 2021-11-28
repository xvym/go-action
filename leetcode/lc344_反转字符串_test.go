/**
 * @author: xv
 * @date: 2021/11/24 14:50
 * @link: https://leetcode-cn.com/problems/reverse-string/
 * @tag: 双指针
 */

package leetcode

import "testing"

func TestLc344(t *testing.T) {
	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
}

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left <= right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp
		left++
		right--
	}
}