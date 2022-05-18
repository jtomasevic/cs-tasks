/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
 var twoSum = function(nums, target) {
    let map = {};
    
    for(let i = 0; i < nums.length; i++) {
        let p = target - nums[i];
        if(p in map) {
            return [map[p], i];
        }
        map[nums[i]] = i;
    }
};