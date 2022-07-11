func singleNumber(nums []int) int {
    //用异或解决
    r := 0
    for _,v := range nums{
        r = r ^ v 
    }
    return r
}