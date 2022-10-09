func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    var (
        n = len(nums)
        best = math.MaxInt32
    )

    // 根据差值的绝对值来更新答案
    update := func(cur int) {
        if abs(cur - target) < abs(best - target) {
            best = cur
        }
    }

    // 枚举 a
    for i := 0; i < n; i++ {
        // 保证和上一次枚举的元素不相等
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        // 使用双指针枚举 b 和 c
        j, k := i + 1, n - 1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            // 如果和为 target 直接返回答案
            if sum == target {
                return target
            }
            update(sum)
            if sum > target {
                // 如果和大于 target，移动 c 对应的指针
                k0 := k - 1
                // 移动到下一个不相等的元素
                for j < k0 && nums[k0] == nums[k] {
                    k0--
                } 
                k = k0
            } else {
                // 如果和小于 target，移动 b 对应的指针
                j0 := j + 1
                // 移动到下一个不相等的元素
                for j0 < k && nums[j0] == nums[j] {
                    j0++
                }
                j = j0
            }
        }
    }
    return best
}

func abs(x int) int {
    if x < 0 {
        return -1 * x
    }
    return x
}