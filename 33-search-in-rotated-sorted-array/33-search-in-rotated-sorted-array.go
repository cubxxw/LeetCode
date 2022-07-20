func search(nums []int, target int) int {
    start, end := 0, len(nums)-1
    for start+1 < end {
        mid := (start+end)>>1
        if nums[mid] > nums[len(nums)-1] {
            start = mid
        } else {
            end = mid
        }
    }
    var minIndex int
    if nums[start] < nums[end] {
        minIndex = start
    } else {
        minIndex = end
    }
    if nums[minIndex] == target {
        return minIndex
    }
    if target > nums[len(nums)-1] {
        start, end = 0, minIndex
    } else {
        start, end = minIndex, len(nums)-1
    }
    for start+1 < end {
        mid := (start+end)>>1
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            end = mid
        } else {
            start = mid
        }
    }
    if nums[start] == target {
        return start
    }
    if nums[end] == target {
        return end
    }
    return -1
}
