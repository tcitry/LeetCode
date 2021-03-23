class Solution:
    def reverse(self, x):
        max_int = 2**31
        res = 0
        ltzero = True if x < 0 else False
        if ltzero:
            x *= -1
        while x != 0:
            res = res * 10 + x % 10
            if ltzero is True:
                if res >= max_int - 1:
                    return 0
            else:
                if res >= max_int:
                    return 0
            x = int(x / 10)
        if ltzero:
            res *= -1
        return res


s = Solution()
print(s.reverse(1534236469))
