func findTheDifference(s string, t string) byte {
    ans := byte(0)
    for  i := range s {
        ans ^= s[i]
    }
    for  i := range t {
        ans ^= t[i]
    }
    return byte(ans)
}