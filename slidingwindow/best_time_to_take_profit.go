// medium problem
package slidingwindow

func MaxProfit(prices []int) int {
	mp := 0
	i, j := 0, 1

	for j < len(prices) {
		if prices[i] > prices[j] {
			i = j
			j = i + 1
		} else {
			mp = max(mp, prices[j]-prices[i])
			j++
		}
	}

	return mp
}
