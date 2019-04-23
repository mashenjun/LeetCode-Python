package bestTimeToBuyAndSellStockII

// 122. Best Time to Buy and Sell Stock II
// 思路：只要是上升的情况，都做一次买卖
func maxProfit(prices []int) int {
	rlt := 0
	var prev = 1<<31 - 1

	for _, v := range prices {
		if v > prev {
			rlt += v - prev
		}
		prev = v
	}
	return rlt
}
