<h2><a href="https://leetcode.com/problems/sort-an-array/">912. Sort an Array</a></h2><h3>Medium</h3><hr><div><p>Given an array of integers <code>nums</code>, sort the array in ascending order.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<pre><strong>Input:</strong> nums = [5,2,3,1]
<strong>Output:</strong> [1,2,3,5]
</pre><p><strong>Example 2:</strong></p>
<pre><strong>Input:</strong> nums = [5,1,1,2,0,0]
<strong>Output:</strong> [0,0,1,1,2,5]
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>-5 * 10<sup>4</sup> &lt;= nums[i] &lt;= 5 * 10<sup>4</sup></code></li>
</ul>

</div>

### go语言sort排序

```
func sortArray(nums []int) []int {
    sort.Ints(nums)
    return nums
}
```

### 冒泡排序

> Go语言冒泡排序超时

```
func sortArray(nums []int) []int {
    //sort.Ints(nums)
    for i:=0; i<len(nums); i++ {
        b := true
        for j:=0; j<len(nums)-i-1; j++ {
            if nums[j]>nums[j+1] {
                //升序排序
                nums[j], nums[j+1] = nums[j+1], nums[j] 
                b = false
            }
        }
       if b==true {break}
    }
    return nums
}
```

### JavaScript冒泡和sort

```
/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortArray = function(nums) {

    for(var i = 0; i<nums.length; i++) {
       var bz = false;
        for(var j =0; j<nums.length-i-1; j++) {
            if(nums[j]>nums[j+1]){
                var test = nums[j];
                nums[j] = nums[j+1];
                nums[j+1] = test;
                bz = true;
            }
        }
         if(bz != true){
            break;
         }
    }
    return nums;
};

```

**JavaScript**

```
   var sortArray = function(nums) {
   a = nums.sort();
   return a;
```

C++快排代码

```
class Solution {
public:
    void swap(vector<int>& nums,int i,int j){
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
    void sort(vector<int>& nums,int start,int end){
        if(start >= end){
            //开始指针和最后指针走到一起
            return;
        }
        int pivot = nums[start];   //选择基准点为开始的位置
        int l = start + 1;     
        int r = end;
        while(l<=r){   //保证退出时不重复
            if(nums[l] < pivot) {
                //如果左指针比基准值要小
                l++;
                continue;
            }
            if(nums[r] >= pivot) {
                //如果左指针比基准值要小
                r--;
                continue;
            }
            //如果都不满足 -- 左边大于或者等于，或者右边小，交换
            swap(nums, l ,r);
        }
        swap(nums,start,r);
        sort(nums,start,r - 1);
        sort(nums,l,end);
    }
    vector<int> sortArray(vector<int>& nums) {
        //快排
        sort(nums,0,nums.size() - 1);
        return nums;
    }
};
```

