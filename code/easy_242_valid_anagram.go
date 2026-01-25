/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s_map := make(map[rune]int)
	for _, char := range s {
		s_map[char]++
	}
	for _, char := range t {
		_, ok := s_map[char]
		if ok {
			s_map[char]--
		} else {
			return false
		}
		if s_map[char] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end
