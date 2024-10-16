func addStrings(num1 string, num2 string) string {
    return recurseAdd(0, len(num1) - 1, len(num2) - 1, num1, num2, "")
}

func recurseAdd(carry, pointer1, pointer2 int, num1, num2, result string) string {
    if pointer1 < 0 && pointer2 < 0 {
        if carry != 0 {
            return fmt.Sprintf("%d", carry) + result
        }

        return result
    }

    temp := 0
    if pointer1 >= 0 && pointer1 < len(num1) {
        temp += int(num1[pointer1] - '0')
    }
    if pointer2 >= 0 && pointer2 < len(num2) {
        temp += int(num2[pointer2] - '0')
    }

    if carry != 0 {
        temp += carry
    }

    if temp >= 10 {
        carry = temp / 10
        temp = temp % 10
    } else {
        carry = 0
    }

    result = fmt.Sprintf("%d", temp) + result
    return recurseAdd(carry, pointer1 - 1, pointer2 - 1, num1, num2, result)
} 

func reverse(s string) string {
    result := ""
    for i := len(s) - 1; i >= 0; i-- {
        result += string(s[i])
    }

    return result
}

