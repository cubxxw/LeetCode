function shuffle(nums: number[], n: number): number[] {
    for (let i = 0; i < n; i ++ ) {
        nums[i * 2] |= (nums[i] & 0x3ff) << 10;
        nums[i * 2 + 1] |= (nums[i + n] & 0x3ff) << 10;
    }
    return nums.map(item => item >> 10)
};