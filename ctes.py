class Solution:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          
    def longestValidParenthese(self, s):
        """
        :type s: str
        :rtype: int
        """
        stack = []
        s1 = list(s)
        count = 0
        maxlen = 0
        stack.append(-1)
        for i in range(len(s1)):
            if s1[i] == "(":
                stack.append(i)
            else:
                stack.pop()
                if not stack:
                    stack.append(i)
                else:
                    maxlen = max(maxlen, i - stack[-1]
    return maxlen                

              