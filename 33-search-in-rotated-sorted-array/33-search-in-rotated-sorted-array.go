func search(nums []int, target int) int {
// 将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
// 此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环. 
    for i,v := range nums {
        if v ^ target == 0 {
            //1. show ==
            return i 
        }
    }
    return -1
}