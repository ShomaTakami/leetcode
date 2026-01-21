/*
 * @lc app=leetcode id=1768 lang=golang
 *
 * [1768] Merge Strings Alternately
 */

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {
	var res string = ""
	var res_b []byte
	count := max(len(word1), len(word2))

	for i := 0; i < count; i++ {
		if i < len(word1) {
			res_b = append(res_b, word1[i])
		}
		if i < len(word2) {
			res_b = append(res_b, word2[i])
		}
	}
	res = string(res_b)
	return res
}

// @lc code=end
