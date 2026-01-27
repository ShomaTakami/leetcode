/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	// Input: s = "([)]"
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []rune{}

	for _, char := range s {
		// charが閉じ括弧の場合
		if openBracket, ok := pairs[char]; ok {
			// スタックが空 or 最後の一文字がペアじゃない
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				// 1つでもダメなら、その瞬間に「この文字列は不合格！」と確定して終了
				return false
			}

			// ペアが合っていたら、積んでいた開き括弧を消して次へ
			stack = stack[:len(stack)-1]
			// charが括弧の場合
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

// @lc code=end
