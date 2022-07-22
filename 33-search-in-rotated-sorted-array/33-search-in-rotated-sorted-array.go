func find(nums []int, start int,end int,target int) int {
    for start <= end {
        mid := ((end - start) >> 1 ) + start;
        
        if nums[mid] == target {
            return mid 
        }else if nums[mid] > target {
            end = mid -1
        }else{
            start = mid + 1
        }
    }
    return -1   //找不到返回-1
}

func search(nums []int, target int) int {
    len := len(nums)
    if len == 0 {
        return -1
    }
    if len == 1 {
        if target == nums[0] {
            return 0
        }else{
            return -1
        }
    }
    
    //找到旋转的位置
    cutoff := -1
    for i:=0; i < len - 1; i++ {
        if nums[i] > nums[i+1] {
            //smart
            cutoff = i    //7所取到的位置
            break
        }
    }
    
    //二分查找
    left := find(nums,0,cutoff,target)
    right := find(nums,cutoff+1,len-1,target)
    
    if left == -1 {
        return right    
    }else{
        return left
    }    
}



