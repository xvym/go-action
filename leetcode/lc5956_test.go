/**
 * @author: xv
 * @date: 2021/12/19 10:42
 */

package leetcode

import "testing"

func TestLc5956(t *testing.T) {

}

func firstPalindrome(words []string) string {
	for i := range words {
		if isPalindrome(words[i]) {
			return words[i]
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for start <= end {
		if s[start] == s[end] {
			start++
			end--
			continue
		} else {
			return false
		}
	}
	return true
}
