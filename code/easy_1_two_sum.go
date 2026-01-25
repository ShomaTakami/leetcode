/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	twos := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		j, ok := twos[diff]
		if ok {
			answer := []int{j, i}
			return answer
		}
		twos[num] = i
	}
	return nil
}

// @lc code=end
