#### BM80 买卖股票的最好时机(一)

> https://mp.weixin.qq.com/s/-p_tyZBSGk0FyA08k93XNw

1. dp[i]表示当天结束时的最大利润，可以是当天卖出，或者之前已经就卖出了
1. 初始条件：dp[0] = 0
1. 状态转移方程：dp[i] = max(dp[i-1],prices[i]-min_price)