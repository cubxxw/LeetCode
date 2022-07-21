func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    //对饼干和小孩的胃口 数组arry进行排序

    //chile := 0
     child := 0   //满足的数量初始值为0
     sIdx := 0
     for child < len(g) && sIdx < len(s) {
        if s[sIdx] >= g[child] {
//如果饼干的大小大于或等于孩子的为空则给与，否则不给予，继续寻找选一个饼干是否符合
        child++
       }
        sIdx++
    }
    return child
}