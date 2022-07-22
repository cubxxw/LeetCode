func countNegatives(grid [][]int) int {
    num := 0
    for _,v := range grid {
        for _, v2:= range v {
            if v2 < 0 {
                num++
            }
        }
    }
    return num
}