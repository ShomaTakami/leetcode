/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	words := strings.Fields(s)
	last := words[len(words)-1]
	return len(last)
}

// @lc code=end
