/**
 * @author: xv
 * @date: 2021/12/14 14:49
 * @link: https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
 * @tag: dfs、回溯
 */

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo012(t *testing.T) {
	fmt.Println(exist([][]byte{{97}, {97}}, "aaa"))
}

func exist(board [][]byte, word string) bool {
	used := make([][]int, len(board))
	for i := range used {
		used[i] = make([]int, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if backtrackInSo012(board, word, i, j, 0, used) {
					return true
				}
			}
		}
	}
	return false
}

func backtrackInSo012(board [][]byte, word string, i int, j int, index int, used [][]int) bool {
	if index == len(word) {
		return true
	} else if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) && used[i][j] != 1 {
		if board[i][j] == word[index] {
			used[i][j] = 1
			if backtrackInSo012(board, word, i+1, j, index+1, used) ||
				backtrackInSo012(board, word, i-1, j, index+1, used) ||
				backtrackInSo012(board, word, i, j+1, index+1, used) ||
				backtrackInSo012(board, word, i, j-1, index+1, used) {
				return true
			}
		}
		used[i][j] = 0
		return false
	} else {
		return false
	}
}
