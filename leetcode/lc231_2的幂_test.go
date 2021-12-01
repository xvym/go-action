/**
 * @author: xv
 * @date: 2021/12/1 15:19
 * @link: https://leetcode-cn.com/problems/power-of-two/
 * @tag: 位运算
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc231(t *testing.T) {
	fmt.Println(isPowerOfTwo(10))
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}