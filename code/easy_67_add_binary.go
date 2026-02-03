/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0 // 繰り上がり
	res := ""
	//文字列の右端からスタート
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		// --- ここで strconv.Itoa を使わない魔法 ---
		// sum % 2 は「0 か 1」
		// それに '0'（ascii48）を足すと、文字の '0' か '1' に戻る
		char := string(rune((sum % 2) + '0'))
		res = char + res

		carry = sum / 2
	}
	return res
}

// @lc code=end
