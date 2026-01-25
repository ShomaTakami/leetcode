/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, s := range strs {
		key := sortString(s)

		m[key] = append(m[key], s)
	}

	ans := [][]string{}
	for _, group := range m {
		ans = append(ans, group)
	}
	return ans
}
func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

// @lc code=end
