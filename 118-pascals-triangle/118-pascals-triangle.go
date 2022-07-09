func generate(numRows int) [][]int {
    ans := make([][]int,numRows)
    if len(ans) < 1{
        return ans
    }
    for i := range ans{
        ans[i] = make([]int,i+1)  //i从零开始de
        ans[i][0] = 1
        ans[i][i] = 1
        for j := 1 ;j < i;j++{
           ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
        }
    }
    return ans
}