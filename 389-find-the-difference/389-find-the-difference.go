func findTheDifference(s string, t string) byte {
    sum,sum1 := 0,0 //求和计数
    for _,v := range s {
        sum += int(v)
    }
    for _,v := range t {
        sum1 += int(v)
    }
    return  byte(sum1 - sum)
}