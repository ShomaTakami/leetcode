/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {

	seen := make(map[int]bool)

	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = true
	}

	return false
}

// @lc code=end
