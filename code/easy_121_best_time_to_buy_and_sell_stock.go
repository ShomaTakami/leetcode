/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		//最安値を更新
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		//「今日の価格で売った場合の利益」を計算
		currentProfit := prices[i] - minPrice

		//それが最大利益を更新するかチェック
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
	}
	return maxProfit
}

//  [7,1,5,3,6,4]
// max_profit: 5
// @lc code=end
