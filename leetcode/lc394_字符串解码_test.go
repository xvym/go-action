/**
 * @author: xv
 * @date: 2021/12/15 17:13
 * @link: https://leetcode-cn.com/problems/decode-string/
 * @tag: 字符串、递归、栈
 */

package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLc394(t *testing.T) {
	fmt.Println(decodeString("3[a]2[bc]"))
	//fmt.Println(decodeString("3[abc]"))
}

func decodeString(s string) string {
	res, _ := doDecode(s, 0)
	return res
}

func doDecode(s string, index int) (string, int) {
	temp := ""
	num := 0
	tempIndex := index
	for i := index; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			curNum, _ := strconv.Atoi(string(s[i]))
			num = num*10 + curNum
		} else if string(s[i]) == "[" {
			curStr, curIndex := doDecode(s, i+1)
			i = curIndex
			for j := 0; j < num; j++ {
				temp += curStr
			}
			num = 0
		} else if string(s[i]) == "]" {
			tempIndex = i
			break
		} else {
			temp += string(s[i])
		}
	}
	return temp, tempIndex
}