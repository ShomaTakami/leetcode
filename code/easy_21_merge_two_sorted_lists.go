/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 1. 【土台作り】
	// 答えのリストの「出発地点」として、空の箱（dummy）を1つ用意する。
	// dummy自体は動かさず、最後のリターン用に住所をキープしておく。
	dummy := &ListNode{}

	// 2. 【作業員の配置】
	// いま工事している場所を指し示す「指」として current を用意。
	// 最初は dummy（土台）と同じ場所を指しておく。
	current := dummy

	// 3. 【合体ループ】
	// 両方のリストにまだ「材料（家）」がある間、繰り返す。
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			// list1の数字の方が小さい（または同じ）場合：

			// ① 今いる場所（current）の出口に、list1の住所を書き込む（道を繋ぐ）
			current.Next = list1
			// ② 作業員（current）自身を、いま繋いだ list1 の場所へ移動させる
			current = list1
			// ③ 材料（list1）のメモを、一個隣の家の住所に更新する
			list1 = list1.Next
		} else {
			// list2の方が小さい場合（やることは上と同じ）：

			current.Next = list2
			current = list2
			list2 = list2.Next
		}
	}

	// 4. 【最後の仕上げ】
	// どちらかの材料が尽きた時、まだ余っている方のリストがあるはず。
	// 連結リストは「先頭を繋げば後ろも全部ついてくる」ので、if文一発でOK！
	if list1 != nil {
		// list1が余っていたら、currentの後ろにガチャンと繋ぐ
		current.Next = list1
	} else {
		// list2が余っていたら、同様に繋ぐ
		current.Next = list2
	}

	// 5. 【結果を返す】
	// dummy自体は空の箱なので、その「一個後ろ（Next）」が本物の答えの先頭。
	return dummy.Next
}

// @lc code=end
