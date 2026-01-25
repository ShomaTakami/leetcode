/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	k := 1
	for i, num := range nums {
		if i == 0 {
			nums[i] = num
		}
		if i >= 1 {
			if nums[k-1] != num {
				nums[k] = num
				k++
			}
		}
	}
	return k
}

// @lc code=end
