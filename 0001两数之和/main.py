class Solution:
    def twoSum(self, nums, target):
        num_index = dict()
        for i in range(len(nums)):
            if target - nums[i] in num_index:
                return [num_index[target-nums[i]], i]
            num_index[nums[i]] = i
