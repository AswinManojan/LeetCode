func maxProfit(prices []int) int {
    mini:=prices[0]
    max_profit,cost:=0,0
    for i:=1;i<len(prices);i++{
        cost=prices[i]-mini
        mini=min(prices[i],mini)
        max_profit=max(max_profit,cost)
    }
    return max_profit
}
func min(a,b int) int{
    if a<b{
        return a
    }else{
        return b
    }
}
func max(a,b int) int{
    if a>b{
        return a
    }else{
        return b
    }
}