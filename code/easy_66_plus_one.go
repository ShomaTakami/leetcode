/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	num := 0
	for i := len(digits) - 1; i >= 0; i-- {
		num = digits[i] + 1
		if num == 10 {
			digits[i] = 0
		} else {
			digits[i] = num
			return digits
		}
	}
	return append([]int{1}, digits...)
}

// @lc code=end
