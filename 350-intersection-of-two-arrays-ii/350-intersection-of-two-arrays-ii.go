func intersect(nums1 []int, nums2 []int) []int {
    //双指针算法、
    sort.Ints(nums1)
    sort.Ints(nums2)
    length1 := len(nums1)
    length2 := len(nums2)
    a, b := 0, 0
    text := []int{}   //切片
    for a < length1 && b < length2 {
        if nums1[a] < nums2[b] {
            a++
        }else if nums1[a] > nums2[b] {
            b++
        }else {
            //注意切片使用的是
            text = append(text,nums2[b])
            a, b = a+1, b+1
        }
    }
    return text
}

