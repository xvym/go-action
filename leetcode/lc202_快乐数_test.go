/**
 * @author: xv
 * @date: 2021/12/3 17:32
 * @link: https://leetcode-cn.com/problems/happy-number/
 * @tag: 其他
 */

package leetcode

import "testing"

func TestLc202(t *testing.T) {
	isHappy(2)
}

func isHappy(n int) bool {
	if n < 10 {
		return false
	}
	myMap := make(map[int]int)
	temp := 0
	for n != 1 {
		for n > 0 {
			temp += (n % 10) * (n % 10)
			n -= n % 10
			n /= 10
		}
		if _, ok := myMap[temp]; ok {
			return false
		} else {
			myMap[temp] = temp
		}
		n = temp
		temp = 0
	}
	return true
}