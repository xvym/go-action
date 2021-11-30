/**
 * @author: xv
 * @date: 2021/12/1 0:18
 * @link: https://leetcode-cn.com/problems/letter-case-permutation/
 * @tag: 回溯
 */

package leetcode

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func TestLc784(t *testing.T) {
	s := "你好abcd"
	for i := range s {
		fmt.Println(rune(s[i]))
	}
}

func letterCasePermutation(s string) []string {
	res := make([]string, 0)
	path := ""
	return backtrack784(0, res, path, s)
}

func backtrack784(index int, res []string, path string, s string) []string {
	if index == len(s) {
		res = append(res, path)
	} else {
		if unicode.IsLetter(rune(s[index])) {
			upperCh := strings.ToUpper(string(s[index]))
			lowerCh := strings.ToLower(string(s[index]))
			upperPath := path + upperCh
			lowerPath := path + lowerCh
			res = backtrack784(index+1, res, upperPath, s)
			res = backtrack784(index+1, res, lowerPath, s)
		} else {
			char := string(s[index])
			res = backtrack784(index+1, res, path+char, s)
		}
	}
	return res
}
