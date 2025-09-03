from typing import List
class Solution:
    def isValid(self, s: str) -> bool:
        
        hashmap = {"[":"]","{":"}", "(":")"}
        stack = []
        for bracket in s:
            if bracket in hashmap:
                stack.append(bracket)
            elif not stack or hashmap[stack.pop()] != bracket:
                return False
        return not stack

# core lesson:
# I've learned not to use return True in this problem,
# return False and checking whether stack is empty is enough


input = ["}"]
solution = Solution()
res = solution.isValid(input)
print(res)