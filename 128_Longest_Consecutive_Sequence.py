from typing import List
class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        
        longest_seq = 0
        hashset = set(nums)
        for num in hashset:

            #determine whether the num is the entry point of a sequence
            if (num - 1) not in hashset:
                seq = 1
                while (num + seq) in hashset:
                    seq += 1
                longest_seq = max(seq, longest_seq)
        return longest_seq


input = [0,3,7,2,5,8,4,6,0,1]
solution = Solution()
res = solution.longestConsecutive(input)
print(res)