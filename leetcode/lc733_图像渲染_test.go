/**
 * @author: xv
 * @date: 2021/11/28 21:17
 * @link: https://leetcode-cn.com/problems/flood-fill/
 * @tag: dfs
 */

package leetcode

import "testing"

func TestLc733(t *testing.T) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	floodFill(image, 1, 1, 2)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	doFloodFill(image, sr, sc, newColor, image[sr][sc])
	return image
}

func doFloodFill(image [][]int, sr int, sc int, newColor int, oldColor int) {
	if sr < len(image) && sc < len(image[0]) && sr >= 0 && sc >= 0 && image[sr][sc] == oldColor && oldColor != newColor {
		image[sr][sc] = newColor
		doFloodFill(image, sr+1, sc, newColor, oldColor)
		doFloodFill(image, sr-1, sc, newColor, oldColor)
		doFloodFill(image, sr, sc+1, newColor, oldColor)
		doFloodFill(image, sr, sc-1, newColor, oldColor)
	} else {
		return
	}
}
