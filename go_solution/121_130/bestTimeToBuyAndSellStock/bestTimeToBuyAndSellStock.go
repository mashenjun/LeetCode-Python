package bestTimeToBuyAndSellStock

// 121. Best Time to Buy and Sell Stock
// 思路：用两个变量记录找到的最小值 和 目前的最大收益
func maxProfit(prices []int) int {
	rlt := 0
	min := 1<<31 - 1
	for _, v := range prices {
		if v < min {
			min = v
		}
		if v-min > rlt {
			rlt = v - min
		}
	}
	return rlt
}
