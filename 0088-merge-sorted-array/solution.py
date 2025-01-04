class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        ptr = len(nums1) - 1
        l = m - 1
        r = n - 1
        while l >= 0 and r >= 0:
            if nums1[l] >= nums2[r]:
                nums1[ptr] = nums1[l]
                l -= 1
            else:
                nums1[ptr] = nums2[r]
                r -= 1
            ptr -= 1

        while l >= 0:
            nums1[ptr] = nums1[l]
            ptr -= 1
            l -= 1

        while r >= 0:
            nums1[ptr] = nums2[r]
            ptr -= 1
            r -= 1

        
