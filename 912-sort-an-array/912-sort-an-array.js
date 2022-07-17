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
