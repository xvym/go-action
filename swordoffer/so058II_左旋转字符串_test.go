/**
 * @author: xv
 * @date: 2021/12/6 17:00
 * @link: https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
 * @tag: 字符串
 */

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo058II(t *testing.T) {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}

/**
 * 不是很了解这题想考察什么，如果不用切片可以参考lc189轮转数组
 */
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}