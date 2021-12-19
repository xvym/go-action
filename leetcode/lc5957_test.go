/**
 * @author: xv
 * @date: 2021/12/19 10:49
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc5957(t *testing.T) {
	fmt.Println(addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15}))
}

func addSpaces(s string, spaces []int) string {
	res := make([]byte, len(s)+len(spaces))
	j := 0
	k := 0
	for i := range res {
		if j < len(spaces) {
			if i+j == spaces[j] {
				res[i+j] = ' '
				j++
			} else {
				res[i] = s[k]
				k++
			}
		}
	}
	return string(res)
}
