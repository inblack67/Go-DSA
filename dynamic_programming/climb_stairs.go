package dynamicprogramming

func Climb(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		if i == 1 {
			dp[i] = dp[i-1]
		} else if i == 2 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
		}
	}
	return dp[n]
}
