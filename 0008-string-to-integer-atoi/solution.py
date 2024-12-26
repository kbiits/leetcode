class Solution:
    def myAtoi(self, s: str) -> int:
        
        def is_digit(c: str) -> bool:
            return ord(c) >= ord('0') and ord(c) <= ord('9')
        
        def is_sign(c: str) -> bool:
            return c == "+" or c == "-"

        init_found_non_whitespace = False
        sign = 1
        found_digit = False
        data = 0
        found_sign = False
        for i, c in enumerate(s):

            # ignoring leading whitespace
            if c == " " and not init_found_non_whitespace:
                continue
            else:
                init_found_non_whitespace = True

            if not is_digit(c) and (not is_sign(c) or (is_sign(c) and found_sign)):
                break
            
            if (not found_digit) and is_sign(c) and not found_sign:
                found_sign = True
                sign = 1 if c == "+" else -1
            
            if (not found_digit) and is_digit(c):
                found_digit = True
            
            if found_digit and is_digit(c):
                data = (data * 10) + (ord(c) - ord('0'))
            elif found_digit and not is_digit(c):
                break

            if sign == -1 and data > 2**31:
                data = 2**31
            elif sign == 1 and data > 2 ** 31 - 1:
                data = 2**31 - 1

        return sign * data

