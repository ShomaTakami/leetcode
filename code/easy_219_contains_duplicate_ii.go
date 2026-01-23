/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	numToIndex := make(map[int]int)

	for i, num := range nums {
		if j, ok := numToIndex[num]; ok {
			if (i - j) <= k {
				return true
			}
		}
		numToIndex[num] = i
	}
	return false
}

// @lc code=end
