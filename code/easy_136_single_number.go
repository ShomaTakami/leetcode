/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		//XOR (排他的論理和)
		result ^= num // 全員をぶつけ合わせる
	}
	return result // 生き残った一人が答え
}

// @lc code=end
