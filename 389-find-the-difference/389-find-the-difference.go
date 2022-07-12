func findTheDifference(s string, t string) byte {
    a := byte(0)
    for i,_ := range t {
        a^=t[i] 
    }
        for i,_ := range s {
        a^=s[i]
    }
    return a
}