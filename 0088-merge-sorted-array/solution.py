class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        
        # nums1 = [1,2,2,3,5,6], m = 3, nums2 = [2,5,6], n = 3
        while n > 0:
            if m > 0 and nums2[n - 1] < nums1[m - 1]:
                nums1[m + n - 1] = nums1[m - 1]
                nums1[m - 1] = nums2[n - 1]
                m -= 1
            else:
                # because it's sorted, just assign it to its corresponding space in nums1
                nums1[m + n - 1] = nums2[n - 1]
                n -= 1



        
