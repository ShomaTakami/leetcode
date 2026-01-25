/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// 1. マップを作る
	// [合言葉] => [その合言葉に合う単語のリスト]
	m := make(map[string][]string)

	for _, s := range strs {
		// 2. 合言葉を作る（一旦、仮に "abc" としておきます）
		key := sortString(s)

		// 3. マップの「合言葉」の場所に、今の単語 s を追加する
		// ヒント：m[key] は []string 型なので、そこに append(m[key], s) すればOK！
		m[key] = append(m[key], s) // ここを書いてみてください
	}

	// 4. 最後に、ステップ1で練習した「詰め替え」をやる
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
