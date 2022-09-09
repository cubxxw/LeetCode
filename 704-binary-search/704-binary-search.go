func search(nums []int, target int) int {
    //给定的是一个升序的数组target
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (right - left)>>1  + left 
        num := nums[mid]
        if num == target {
            return mid
        } else if num > target{
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return -1
}


