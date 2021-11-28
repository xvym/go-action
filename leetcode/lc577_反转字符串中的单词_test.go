/**
 * @author: xv
 * @date: 2021/11/24 14:57
 * @link: https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/
 * @tag: 双指针
 */

package leetcode

import "testing"

func TestLc577(t *testing.T) {
	reverseWords("hello 你好")
}

func reverseWords(s string) string {
	length := len(s)
	var ret []byte
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		for p := start; p < i; p++ {
			ret = append(ret, s[start+i-1-p])
		}
		for i < length && s[i] == ' ' {
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}