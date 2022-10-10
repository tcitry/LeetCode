def quicksort(nums):
    def partition(nums, start, end):
        pivot = nums[end]
        i = start
        for j in range(start, end):
            if pivot >= nums[j]:
                nums[i], nums[j] = nums[j], nums[i]
                i += 1
        nums[end] = nums[i]
        nums[i] = pivot
        return i

    def sort(nums, start, end):
        if end <= start:
            return nums
        i = partition(nums, start, end)
        sort(nums, start, i-1)
        sort(nums, i+1, end)
        return nums

    if len(nums) < 1:
        return nums
    return sort(nums, 0, len(nums)-1)

def mergesort(nums):
    def merge(nums1, nums2):
        result = []
        while(len(nums1) and len(nums2)):
            if nums1[0] < nums2[0]:
                result.append(nums1[0])
                nums1.pop(0)
            else:
                result.append(nums2[0])
                nums2.pop(0)
        if nums1:
            result += nums1
        else:
            result += nums2
        return result

    if len(nums) < 2:
        return nums
    middle = int(len(nums)/2)
    left = mergesort(nums[0:middle])
    right = mergesort(nums[middle:])
    return merge(left, right)
