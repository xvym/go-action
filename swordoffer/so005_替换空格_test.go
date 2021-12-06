/**
* @author: xv
* @date: 2021/12/6 16:40
* @link: https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
  @tag: 字符串
*/

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo005(t *testing.T) {
	fmt.Println(replaceSpace(" We are happy."))
}

func replaceSpace(s string) string {
	size := len(s)
	for i := 0; i < size; i++ {
		char := string(s[i])
		if char == " " {
			s = s[:i] + "%20" + s[i+1:]
			size += 2
		}
	}
	return s
}